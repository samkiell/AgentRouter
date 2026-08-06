# Option 1: Direct API (Raw HTTP / cURL)

Using AgentRouter's base URL directly in custom applications via raw HTTP or cURL is **blocked by Cloudflare WAF design ❌**.

## WAF Restrictions
Standard HTTP requests return `unauthorized client detected`. AgentRouter requires client fingerprint spoofing headers.

## Required Headers
If making raw HTTP calls from custom code, include these headers on every request:

```http
Originator: codex_cli_rs
Version: 0.101.0
User-Agent: codex_cli_rs/0.101.0 (Mac OS 26.0.1; arm64) Apple_Terminal/464
```

## Base URLs
- **OpenAI Endpoint (`/v1/chat/completions`)**: `https://agentrouter.org/v1`
- **Claude Native Endpoint (`/v1/messages`)**: `https://agentrouter.org` (NO `/v1`)

## Recommended Path
Use **Option 2 (Coding Agents)** or **Option 3 (Desktop Apps)** for reliable, working integrations.
