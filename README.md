# AgentRouter Integration Guide

An unofficial, complete integration guide for **AgentRouter** — an OpenAI & Anthropic compatible API gateway.

---

## ⚡ Quick Start

Launch the interactive Web Dashboard directly to test live API calls, generate cURL commands, and view setup guides:

👉 **[agentrouter.samkiel.online](https://agentrouter.samkiel.online)**

---

## 🚀 The 3 Integration Methods

AgentRouter supports 3 distinct integration paths depending on your workflow:

<br/>

> ### 1️⃣ Direct API (Raw HTTP / cURL)
> **Status**: ❌ *Blocked by Design / Requires Client Spoofing*
> 
> Directly calling AgentRouter's base URL in custom apps or raw cURL scripts is blocked by Cloudflare WAF unless client spoofing headers are explicitly sent.
> 
> - **Why it's blocked**: WAF returns `unauthorized client detected` for standard HTTP clients.
> - **Fix**: Include `Originator`, `Version`, and `User-Agent` headers on every call.
> 
> 📖 **[Read Full Direct API Guide →](direct-api/README.md)**

<br/>

> ### 2️⃣ Coding Agents & IDEs
> **Status**: ✅ *Confirmed Working*
> 
> Seamlessly integrates into official developer tools, CLI agents, and IDE extensions without WAF issues.
> 
> - **Supported**: Claude Code CLI, Claude Code VS Code, Cline, Cursor IDE, Roo Code / Kilo Code.
> - **Rule**: Claude models use `https://agentrouter.org`, GPT models use `https://agentrouter.org/v1`.
> 
> 📖 **[Read Full Coding Agents Guide →](coding-agents/README.md)**

<br/>

> ### 3️⃣ Desktop Applications
> **Status**: ✅ *Confirmed Working*
> 
> Integrates into desktop AI clients (Claude Desktop, Cowork, Cursor Desktop, custom tools).
> 
> - **Supported**: Claude Desktop, Cowork, Cursor, and any app supporting custom Base URLs.
> 
> 📖 **[Read Full Desktop Apps Guide →](desktop-apps/README.md)**

---

## 📦 Additional Integration Resources

- **Postman Collections**: Pre-configured Postman collections in [`postman/`](postman/) (`openai-collection.json`, `claude-collection.json`).
- **7-Language Code Snippets**: Native code implementations in [`snippets/`](snippets/) (Node.js, TypeScript, Python, Go, PHP, C#, cURL).

---

## 🔑 API Key & Available Models

Register at [agentrouter.org](https://agentrouter.org/register?aff=uhNr) to get your API key.

| Model ID | Description | Primary Base Endpoint |
|---|---|---|
| `gpt-5.6-sol` | Fast GPT model | `https://agentrouter.org/v1/chat/completions` |
| `claude-opus-4-8` | Balanced Claude model | `https://agentrouter.org/v1/messages` |
| `claude-opus-5` | High quality Claude model | `https://agentrouter.org/v1/messages` |

---

Made with garri💔 by **[ѕαмкιєℓ.∂єν](https://samkiel.dev)**
