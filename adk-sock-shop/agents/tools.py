from typing import List, Sequence

from google.adk.tools.base_toolset import BaseToolset

from .shared_utils.mcp_tools import create_mcp_toolsets as _create_mcp_toolsets


def create_mcp_toolsets(
    tools_cfg: Sequence[str],
) -> List[BaseToolset]:
    """Return MCPToolset objects - let ADK handle async initialization naturally."""
    return _create_mcp_toolsets(tools_cfg)
