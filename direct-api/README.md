# Option 1: Direct API (Raw HTTP / cURL)

Using AgentRouter's base URL straight in your own app/codebase via raw HTTP or cURL is **blocked by Cloudflare WAF design ❌**.

## Reason
AgentRouter enforces Cloudflare Web Application Firewall (WAF) checks that return `unauthorized client detected` for standard HTTP clients.

## WAF Bypass Headers (Experimental)
If making direct calls, the following client fingerprint headers must be included:

```http
Originator: codex_cli_rs
Version: 0.101.0
User-Agent: codex_cli_rs/0.101.0 (Mac OS 26.0.1; arm64) Apple_Terminal/464
```

## Base URL
- OpenAI Format: `https://agentrouter.org/v1`
- Claude Format: `https://agentrouter.org` (no `/v1`)

## Recommendation
Use **Option 2 (Coding Agents)** or **Option 3 (Desktop Apps)** for reliable integrations.
