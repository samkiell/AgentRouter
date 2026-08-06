# AgentRouter Integration

Complete integration guide for AgentRouter — an unofficial OpenAI & Anthropic compatible API gateway.

> 💡 **Interactive Dashboard**: Access the live interactive sandbox and guide at [`index.html`](index.html) (or [`dash/index.html`](dash/index.html) for deployment at `agentrouter.samkiel.dev`).

---

## 3 Integration Methods

AgentRouter supports 3 distinct integration paths:

### 1. Direct API (Raw HTTP / cURL) — ❌ Blocked by Design
Using AgentRouter's base URL straight in your own custom app or codebase via raw HTTP / cURL calls is **blocked by Cloudflare WAF design**.

- **Why it's blocked**: Cloudflare WAF returns `unauthorized client detected` for standard HTTP clients.
- **WAF Bypass Headers (Experimental)**: Requires spoofing 3 specific client headers on every request:
  ```http
  Originator: codex_cli_rs
  Version: 0.101.0
  User-Agent: codex_cli_rs/0.101.0 (Mac OS 26.0.1; arm64) Apple_Terminal/464
  ```
- **Base URLs**:
  - OpenAI Format (`/v1/chat/completions`): `https://agentrouter.org/v1`
  - Claude Format (`/v1/messages`): `https://agentrouter.org` (NO `/v1`)
- **Folder**: See [`direct-api/README.md`](direct-api/README.md) for details & raw snippets.

---

### 2. Coding Agents & IDEs — ✅ Confirmed Working
Integrating AgentRouter directly into official coding tools, CLI agents, and IDE extensions.

- **Supported Tools**:
  - **Claude Code CLI** (`npm install -g @anthropic-ai/claude-code`)
  - **Claude Code VS Code Extension** (Configured via `claudeCode.environmentVariables` in settings.json)
  - **Cline Extension** (Anthropic provider `https://agentrouter.org` or OpenAI Compatible provider `https://agentrouter.org/v1`)
  - **Cursor IDE** (Override OpenAI Base URL to `https://agentrouter.org/v1`)
  - **Roo Code / Kilo Code** (Custom Provider profile)
- **Folder**: See [`coding-agents/README.md`](coding-agents/README.md) for step-by-step setup guides and environment configurations.

---

### 3. Desktop Applications — ✅ Confirmed Working
Using AgentRouter inside desktop AI applications (Claude Desktop, Cowork, Cursor, etc.).

- **General Endpoint Rules**:
  - **Claude / Anthropic Style Apps**: Set Base URL to `https://agentrouter.org` (NO `/v1`).
  - **OpenAI Style Apps**: Set Base URL to `https://agentrouter.org/v1` (WITH `/v1`).
- **Claude Desktop Setup**:
  1. Open Settings → Custom Provider / API.
  2. Set Base URL to `https://agentrouter.org`.
  3. Paste your AgentRouter API key.
  4. Select model (`claude-opus-4-8` or `claude-opus-5`).
- **Folder**: See [`desktop-apps/README.md`](desktop-apps/README.md) for full instructions.

---

## Additional Integration Resources

- **Postman Collections**: Pre-built Postman collections in [`postman/`](postman/) (`openai-collection.json`, `claude-collection.json`).
- **7-Language Snippets**: Native code implementations in [`snippets/`](snippets/) (Node.js, TypeScript, Python, Go, PHP, C#, cURL).

---

## API Key & Models

Register at [agentrouter.org](https://agentrouter.org/register?aff=uhNr) to get your API key.

| Model Name | Description | Base Endpoint |
|---|---|---|
| `gpt-5.6-sol` | Fast GPT model | `/v1/chat/completions` |
| `claude-opus-4-8` | Balanced Claude model | `/v1/messages` or `/v1/chat/completions` |
| `claude-opus-5` | High quality Claude model | `/v1/messages` or `/v1/chat/completions` |

---

Built by **ѕαмкιєℓ.∂єν** · [Portfolio](https://samkiel.dev)
