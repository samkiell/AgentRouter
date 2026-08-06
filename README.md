# AgentRouter Direct API Guide

Use AgentRouter as an OpenAI-compatible API endpoint in your own apps, websites, or bots.

---

## Get an API Key

Register at [agentrouter.org](https://agentrouter.org/register?aff=uhNr) and copy your key.

---

## Set Your API Key

**Mac / Linux (bash/zsh)**
```bash
export AGENTROUTER_API_KEY="your_api_key_here"
```

**Windows — PowerShell**
```powershell
$env:AGENTROUTER_API_KEY = "your_api_key_here"
```

**Windows — CMD**
```cmd
set AGENTROUTER_API_KEY=your_api_key_here
```

> The key is read from the environment variable in every snippet. Don't hardcode it.

---

## Why Extra Headers?

AgentRouter's WAF rejects plain requests with `unauthorized client detected`. The three headers below spoof an allowed client fingerprint. **Include them on every request.**

```
Originator: codex_cli_rs
Version: 0.101.0
User-Agent: codex_cli_rs/0.101.0 (Mac OS 26.0.1; arm64) Apple_Terminal/464
```

See [`waf-bypass.md`](waf-bypass.md) for details.

---

## Models

| Model | Notes |
|---|---|
| `gpt-5.6-sol` | Fast, good for most tasks |
| `claude-opus-4-8` | Balanced quality / speed |
| `claude-opus-5` | Highest quality, slower |

---

## Quick Start — cURL

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

## Snippets

Each snippet installs cleanly and runs with a single command.

### OpenAI-Compatible (`/v1/chat/completions`)

| Language | File | Install | Run |
|---|---|---|---|
| cURL / Bash | [`snippets/curl/chat.sh`](snippets/curl/chat.sh) | *(none)* | `bash snippets/curl/chat.sh` |
| Node.js | [`snippets/node/openai.js`](snippets/node/openai.js) | `npm install openai` | `node snippets/node/openai.js` |
| TypeScript | [`snippets/typescript/openai.ts`](snippets/typescript/openai.ts) | `npm install openai tsx` | `npx tsx snippets/typescript/openai.ts` |
| Python | [`snippets/python/openai.py`](snippets/python/openai.py) | `pip install openai` | `python snippets/python/openai.py` |
| Go | [`snippets/go/main.go`](snippets/go/main.go) | `go mod init app && go mod tidy` | `go run snippets/go/main.go` |
| PHP | [`snippets/php/openai.php`](snippets/php/openai.php) | *(none — uses cURL extension)* | `php snippets/php/openai.php` |
| C# (.NET) | [`snippets/csharp/Program.cs`](snippets/csharp/Program.cs) | *(none — uses System.Net.Http)* | `dotnet run --project snippets/csharp` |

### Native Claude API (`/v1/messages`)

Uses the Anthropic Messages format. Models default shown — switch between `claude-opus-5` and `claude-opus-4-8` via the comment in each file.

| Language | File | Default Model | Install | Run |
|---|---|---|---|---|
| cURL / Bash | [`snippets/curl/claude.sh`](snippets/curl/claude.sh) | `claude-opus-5` | *(none)* | `bash snippets/curl/claude.sh` |
| Node.js | [`snippets/node/claude.js`](snippets/node/claude.js) | `claude-opus-4-8` | `npm install @anthropic-ai/sdk` | `node snippets/node/claude.js` |
| TypeScript | [`snippets/typescript/claude.ts`](snippets/typescript/claude.ts) | `claude-opus-5` | `npm install @anthropic-ai/sdk tsx` | `npx tsx snippets/typescript/claude.ts` |
| Python | [`snippets/python/claude.py`](snippets/python/claude.py) | `claude-opus-4-8` | `pip install anthropic` | `python snippets/python/claude.py` |
| Go | [`snippets/go/claude.go`](snippets/go/claude.go) | `claude-opus-5` | `go mod init app && go mod tidy` | `go run snippets/go/claude.go` |
| PHP | [`snippets/php/claude.php`](snippets/php/claude.php) | `claude-opus-4-8` | *(none — uses cURL extension)* | `php snippets/php/claude.php` |
| C# (.NET) | [`snippets/csharp/Claude.cs`](snippets/csharp/Claude.cs) | `claude-opus-5` | *(none — uses System.Net.Http)* | `dotnet run --project snippets/csharp` |

---

## Troubleshooting

| Error | Cause | Fix |
|---|---|---|
| `unauthorized client detected` | Missing WAF headers | Add all 3 headers from the section above |
| `null` / empty output | Key not set | Check `echo $AGENTROUTER_API_KEY` (or `echo %AGENTROUTER_API_KEY%` on Windows) |
| `Cannot find module 'openai'` | Dep not installed | Run `npm install openai` in the snippet folder |
| `ModuleNotFoundError: openai` | Dep not installed | Run `pip install openai` |

---

## Postman

Import these collections into Postman to test without writing code:

- [`postman/openai-collection.json`](postman/openai-collection.json)
- [`postman/claude-collection.json`](postman/claude-collection.json)

Set `AGENTROUTER_API_KEY` as a Postman environment variable.

---

## Disclaimer

This uses an unofficial client fingerprint workaround. AgentRouter may update WAF rules without notice. **Never commit API keys to version control.**

---

Built by **ѕαмкιєℓ.∂єν** · [Portfolio](https://samkiel.dev)
