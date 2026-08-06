# AgentRouter Direct API Guide

Use AgentRouter as a standard OpenAI-compatible API endpoint in your applications, services, or automation scripts.

## Active Models
- `gpt-5.6-sol`
- `claude-opus-4-8`
- `claude-opus-5`

---

## Quick Start

### 1. Requirements
* API Key from [agentrouter.org](https://agentrouter.org/register?aff=uhNr)
* Custom WAF Bypass Headers set on **every** request

```bash
export AGENTROUTER_API_KEY="your_api_key_here"
```

### 2. Required Headers
```http
Originator: codex_cli_rs
Version: 0.101.0
User-Agent: codex_cli_rs/0.101.0 (Mac OS 26.0.1; arm64) Apple_Terminal/464
```

### 3. cURL Example
```bash
curl -X POST https://agentrouter.org/v1/chat/completions \
  -H "Authorization: Bearer $AGENTROUTER_API_KEY" \
  -H "Content-Type: application/json" \
  -H "Originator: codex_cli_rs" \
  -H "Version: 0.101.0" \
  -H "User-Agent: codex_cli_rs/0.101.0 (Mac OS 26.0.1; arm64) Apple_Terminal/464" \
  -d '{
    "model": "gpt-5.6-sol",
    "messages": [{"role": "user", "content": "Hello!"}]
  }'
```

---

## Repository Structure

- [`examples/`](examples/)
  - [cURL / Bash](examples/curl/chat.sh)
  - [Go](examples/go/main.go)
  - [Node.js](examples/node/openai.js)
  - [Python](examples/python/openai.py)
  - [TypeScript](examples/typescript/openai.ts)
  - [PHP](examples/php/openai.php)
  - [C# (.NET)](examples/csharp/Program.cs)
- [`postman/`](postman/)
  - [OpenAI Collection](postman/openai-collection.json)
  - [Claude Collection](postman/claude-collection.json)
- [`waf-bypass.md`](waf-bypass.md) — Technical details regarding WAF header requirements

---

## Disclaimer

This project uses an unofficial workaround for client verification. AgentRouter may update fingerprint detection rules without notice. Never commit API keys to version control.

---

Built by **ѕαмкιєℓ.∂єν** · [Portfolio](https://samkiel.dev)
