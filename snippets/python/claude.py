# Install: pip install requests
# Run:     python claude.py
# Model:   claude-opus-4-8 | claude-opus-5

import os
import requests

api_key = os.getenv("AGENTROUTER_API_KEY")
if not api_key:
    raise SystemExit("Error: AGENTROUTER_API_KEY not set.")

response = requests.post(
    "https://agentrouter.org/v1/messages",
    headers={
        "Authorization": f"Bearer {api_key}",
        "x-api-key": api_key,
        "Content-Type": "application/json",
        "anthropic-version": "2023-06-01",
        "Originator": "codex_cli_rs",
        "Version": "0.101.0",
        "User-Agent": "codex_cli_rs/0.101.0 (Mac OS 26.0.1; arm64) Apple_Terminal/464",
    },
    json={
        "model": "claude-opus-4-8",  # or "claude-opus-5"
        "max_tokens": 256,
        "messages": [{"role": "user", "content": "Hi, who are you?"}],
    },
)

data = response.json()
print(data.get("content", [{}])[0].get("text", response.text))
