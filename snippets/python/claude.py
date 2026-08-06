# Install: pip install anthropic
# Run:     python claude.py
# Model:   claude-opus-4-8 | claude-opus-5

import os
import anthropic

client = anthropic.Anthropic(
    api_key=os.getenv("AGENTROUTER_API_KEY"),
    base_url="https://agentrouter.org",
    default_headers={
        "Originator": "codex_cli_rs",
        "Version": "0.101.0",
        "User-Agent": "codex_cli_rs/0.101.0 (Mac OS 26.0.1; arm64) Apple_Terminal/464",
    },
)

message = client.messages.create(
    model="claude-opus-4-8",  # or "claude-opus-5"
    max_tokens=256,
    messages=[{"role": "user", "content": "Hi, who are you?"}],
)

print(message.content[0].text)
