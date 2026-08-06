# ✅ Option 3: Desktop Applications Integration

> **Status**: Confirmed Working ✅

Use AgentRouter inside desktop AI applications (Claude Desktop, Cowork, Cursor, etc.).

👉 **Interactive Setup & Testing Suite**: **[agentrouter.samkiel.online](https://agentrouter.samkiel.online)**

---

## 📌 Endpoint Base URL Rules

| Application Type | Required Base URL | Notes |
|---|---|---|
| **Claude / Anthropic Apps** | `https://agentrouter.org` | **NO `/v1` suffix** |
| **OpenAI Style Apps** | `https://agentrouter.org/v1` | **WITH `/v1` suffix** |

---

## 🖥️ Setup Guides

### 1. Claude App (Desktop)
Claude App is an IDE / desktop Harness that connects to AgentRouter through a third-party inference gateway.

> ⚠️ **WARNING**: You need to enable developer mode before you can configure the gateway.

**Installation and configuration steps:**
1. **Enable developer mode**: `Help` → `Troubleshooting` → `Enable developer mode`.
2. **Configure AgentRouter Gateway**:
   | Configuration Item | Fill in the content |
   |---|---|
   | **Gateway base URL** | `https://agentrouter.org` |
   | **Gateway API key** | The API Key you applied for at AgentRouter |
   | **Gateway auth scheme** | `bearer` |
3. **Apply and restart**: Click `Apply locally` → `Relaunch now`, and after restarting, select the model in the lower left corner to use it.

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
