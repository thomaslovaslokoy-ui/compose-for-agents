import sys
from pathlib import Path
from typing import List, Sequence

from google.adk.tools.base_toolset import BaseToolset

# Add the parent directory's parent to the Python path to import from adk
sys.path.insert(0, str(Path(__file__).parent.parent.parent.parent / "adk" / "agents"))

from shared_utils.mcp_tools import create_mcp_toolsets as _create_mcp_toolsets


def create_mcp_toolsets(
    tools_cfg: Sequence[str],
) -> List[BaseToolset]:
    """Return MCPToolset objects - let ADK handle async initialization naturally."""
    return _create_mcp_toolsets(tools_cfg)
