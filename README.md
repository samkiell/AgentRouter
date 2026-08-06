# AgentRouter Direct API Guide

Use AgentRouter as a normal OpenAI-compatible endpoint in your own apps, websites, or bots.

**Working models right now:**
- `gpt-5.6-sol`
- `claude-opus-4-8`
- `claude-opus-5`

## Quick Start

1. Get your API key from [agentrouter.org](https://agentrouter.org/register?aff=uhNr)
2. Always send these 3 extra headers (this is the WAF bypass):

```
Originator: codex_cli_rs
Version: 0.101.0
User-Agent: codex_cli_rs/0.101.0 (Mac OS 26.0.1; arm64) Apple_Terminal/464
```

3. Base URL: `https://agentrouter.org/v1`

That’s it. Point any OpenAI-compatible client at it.

## Folders

- [`postman/`](postman/) → ready-to-import Postman collections
- [`examples/`](examples/) → Node, TypeScript, Python snippets
- [`waf-bypass.md`](waf-bypass.md) → short explanation why the headers are needed

## Important

This is an unofficial workaround. AgentRouter can change the fingerprint check anytime.  
Use at your own risk. Never commit your real API key.

---

Made with garri💔 by **ѕαмкιєℓ.∂єν** 
Portfolio · [ѕαмкιєℓ.∂єν](https://samkiel.dev)
