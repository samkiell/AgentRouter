# ✅ Option 3: Desktop Applications Integration

> **Status**: Confirmed Working ✅

Use AgentRouter inside desktop AI applications (Claude Desktop, Cowork, Cursor, etc.).

---

## 📌 Endpoint Base URL Rules

| Application Type | Required Base URL | Notes |
|---|---|---|
| **Claude / Anthropic Apps** | `https://agentrouter.org` | **NO `/v1` suffix** |
| **OpenAI Style Apps** | `https://agentrouter.org/v1` | **WITH `/v1` suffix** |

---

## 🖥️ Setup Guides

### 1. Claude Desktop
1. Open Claude Desktop settings (`Preferences` → `Developer / Custom Provider`).
2. Set Base URL: `https://agentrouter.org`
3. Enter API Key: `sk-your-api-key`
4. Set Model ID: `claude-opus-4-8` or `claude-opus-5`

---

### 2. Cursor Desktop / Cowork
1. Open Application Settings → API / Model Configuration.
2. Select OpenAI / Custom Endpoint mode.
3. Set Base URL: `https://agentrouter.org/v1`
4. Enter API Key: `sk-your-api-key`
5. Select Model ID: `gpt-5.6-sol` or `claude-opus-4-8`

---

### 3. Generic Desktop AI Client
For any third-party desktop app supporting custom OpenAI or Anthropic endpoints:
- **API Key**: Enter your AgentRouter API key.
- **Endpoint Base URL**: `https://agentrouter.org` (for Anthropic) or `https://agentrouter.org/v1` (for OpenAI).
- **Available Models**: `gpt-5.6-sol`, `claude-opus-4-8`, `claude-opus-5`.
