package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// DuckDuckGoTool implements the DuckDuckGo tool functionality
type DuckDuckGoTool struct {
	clientSession *mcp.ClientSession
	mcpTool       *mcp.Tool
	args          map[string]any
}

func (tool *DuckDuckGoTool) Name() string {
	return tool.mcpTool.Name
}

func (tool *DuckDuckGoTool) Description() string {
	return tool.mcpTool.Description
}

// Call implements the tool interface. It sets the arguments for the tool and calls the tool.
// For the fetch_content tool, the input is the URL to fetch.
// For the search tool, the input is the query to search for.
func (tool *DuckDuckGoTool) Call(ctx context.Context, input string) (string, error) {
	switch tool.mcpTool.Name {
	case "fetch_content":
		tool.args["url"] = input
	case "search":
		tool.args["query"] = input
	default:
		return "", fmt.Errorf("unsupported tool: %s", tool.mcpTool.Name)
	}

	toolResponse, err := tool.clientSession.CallTool(ctx, &mcp.CallToolParams{
		Name:      tool.mcpTool.Name,
		Arguments: tool.args,
	})
	if err != nil {
		return "", err
	}

	resultText := ""
	for _, content := range toolResponse.Content {
		jsonBytes, err := content.MarshalJSON()
		if err != nil {
			return "", fmt.Errorf("marshal json content: %w", err)
		}

		var contentResponse ContentResponse
		err = json.Unmarshal(jsonBytes, &contentResponse)
		if err != nil {
			return "", fmt.Errorf("unmarshal json content into: %w", err)
		}

		switch contentResponse.Type {
		case "text":
			resultText += content.(*mcp.TextContent).Text
		default:
			return "", fmt.Errorf("unsupported response type (yet): %s", contentResponse.Type)
		}
	}
	return resultText, nil
}

type ContentResponse struct {
	Type string `json:"type"`
}
