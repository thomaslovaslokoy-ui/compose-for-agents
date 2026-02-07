"""Reviser agent for correcting inaccuracies based on verified findings."""

import os

from google.adk import Agent
from google.adk.models.lite_llm import LiteLlm

from . import prompt
from ...shared_utils.llm_callbacks import _remove_end_of_edit_mark, force_string_content


reviser_agent = Agent(
    # OPENAI_MODEL_NAME is set by entrypoint.sh with the model name
    model=LiteLlm(model=os.environ.get("OPENAI_MODEL_NAME", "")),
    name="reviser_agent",
    instruction=prompt.REVISER_PROMPT,
    before_model_callback=force_string_content,
    after_model_callback=_remove_end_of_edit_mark,
)
