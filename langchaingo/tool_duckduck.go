package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// DuckDuckGoTool implements the DuckDuckGo tool functionality
type DuckDuckGoTool struct {
	clientSession *mcp.ClientSession
	mcpTool       *mcp.Tool
	args          map[string]any
}

func (t *DuckDuckGoTool) Name() string {
	return t.mcpTool.Name
}

func (t *DuckDuckGoTool) Description() string {
	return t.mcpTool.Description
}

// Call implements the tool interface. It sets the arguments for the tool and calls the tool.
// For the fetch_content tool, the input is the URL to fetch.
// For the search tool, the input is the query to search for.
func (t *DuckDuckGoTool) Call(ctx context.Context, input string) (string, error) {
	switch t.mcpTool.Name {
	case "fetch_content":
		t.args["url"] = input
	case "search":
		t.args["query"] = input
	default:
		return "", fmt.Errorf("unsupported tool: %s", t.mcpTool.Name)
	}

	res, err := t.clientSession.CallTool(ctx, &mcp.CallToolParams{
		Name:      t.mcpTool.Name,
		Arguments: t.args,
	})
	if err != nil {
		return "", err
	}

	// Use strings.Builder for efficient string concatenation
	var builder strings.Builder
	for _, content := range res.Content {
		// Use type assertion for better performance - only TextContent is supported
		if textContent, ok := content.(*mcp.TextContent); ok {
			builder.WriteString(textContent.Text)
		} else {
			// Non-text content type encountered - determine its type for error reporting
			bs, err := content.MarshalJSON()
			if err != nil {
				return "", fmt.Errorf("marshal json content: %w", err)
			}

			var r response
			err = json.Unmarshal(bs, &r)
			if err != nil {
				return "", fmt.Errorf("unmarshal json content into: %w", err)
			}

			return "", fmt.Errorf("unsupported response type (yet): %s", r.Type)
		}
	}
	return builder.String(), nil
}

type response struct {
	Type string `json:"type"`
}
