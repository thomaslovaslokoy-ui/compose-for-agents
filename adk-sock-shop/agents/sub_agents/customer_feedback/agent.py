"""Search Customer Feedback Agent."""

import os

from google.adk import Agent
from google.adk.models.lite_llm import LiteLlm

from . import prompt
from ...shared_utils.llm_callbacks import _remove_end_of_edit_mark, force_string_content
from ...tools import create_mcp_toolsets

tools = create_mcp_toolsets(tools_cfg=["mcp/mongodb:find", "mcp/mongodb:count"])

customer_feedback_agent = Agent(
    # Using local model runner with MODEL_RUNNER_URL
    model=LiteLlm(model=f"openai/{os.environ.get('MODEL_RUNNER_MODEL')}", api_base=f"{os.environ.get('MODEL_RUNNER_URL')}"),
    name="customer_feedback_agent",
    instruction=prompt.PROMPT,
    tools=tools, # type: ignore
)
