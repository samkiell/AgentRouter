#!/usr/bin/env bash
# Run:   bash claude.sh
# Model: claude-opus-5 | claude-opus-4-8

if [ -z "$AGENTROUTER_API_KEY" ]; then
  echo "Error: AGENTROUTER_API_KEY not set."
  exit 1
fi

curl -s -X POST https://agentrouter.org/v1/messages \
  -H "Authorization: Bearer $AGENTROUTER_API_KEY" \
  -H "x-api-key: $AGENTROUTER_API_KEY" \
  -H "Content-Type: application/json" \
  -H "anthropic-version: 2023-06-01" \
  -H "Originator: codex_cli_rs" \
  -H "Version: 0.101.0" \
  -H "User-Agent: codex_cli_rs/0.101.0 (Mac OS 26.0.1; arm64) Apple_Terminal/464" \
  -d '{
    "model": "claude-opus-5",
    "max_tokens": 256,
    "messages": [{"role": "user", "content": "Hi, who are you?"}]
  }'
