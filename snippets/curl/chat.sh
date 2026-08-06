#!/usr/bin/env bash
# Run: bash chat.sh

if [ -z "$AGENTROUTER_API_KEY" ]; then
  echo "Error: AGENTROUTER_API_KEY environment variable is not set."
  exit 1
fi

curl -s -X POST https://agentrouter.org/v1/chat/completions \
  -H "Authorization: Bearer $AGENTROUTER_API_KEY" \
  -H "Content-Type: application/json" \
  -H "Originator: codex_cli_rs" \
  -H "Version: 0.101.0" \
  -H "User-Agent: codex_cli_rs/0.101.0 (Mac OS 26.0.1; arm64) Apple_Terminal/464" \
  -d '{
    "model": "gpt-5.6-sol",
    "messages": [
      {
        "role": "user",
        "content": "Hi, who are you?"
      }
    ],
    "max_tokens": 256
  }' | grep -o '"content":"[^"]*"' | head -n 1
