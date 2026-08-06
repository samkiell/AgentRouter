# ✅ Option 2: Coding Agents & IDE Integration

> **Status**: Confirmed Working ✅

Integrate AgentRouter directly into official AI coding tools, CLI agents, and IDE extensions.

👉 **Interactive Setup & Testing Suite**: **[agentrouter.samkiel.online](https://agentrouter.samkiel.online)**

---

## 📌 Golden Rule: Base URL Format

- **Claude / Anthropic Models**: Base URL **MUST NOT** include `/v1` (`https://agentrouter.org`).
- **GPT / OpenAI Models**: Base URL **MUST** include `/v1` (`https://agentrouter.org/v1`).

---

## 🛠️ Step-by-Step Tool Setup Guides

### 1. Claude Code CLI
Install the official Claude Code CLI:
```bash
npm install -g @anthropic-ai/claude-code
```

Set environment variables before running:
- **macOS / Linux**:
  ```bash
  export ANTHROPIC_BASE_URL="https://agentrouter.org"
  export ANTHROPIC_AUTH_TOKEN="sk-your-api-key"
  export ANTHROPIC_API_KEY="sk-your-api-key"
  export ANTHROPIC_MODEL="claude-opus-4-8"
  ```
- **Windows (PowerShell)**:
  ```powershell
  $env:ANTHROPIC_BASE_URL="https://agentrouter.org"
  $env:ANTHROPIC_AUTH_TOKEN="sk-your-api-key"
  $env:ANTHROPIC_API_KEY="sk-your-api-key"
  $env:ANTHROPIC_MODEL="claude-opus-4-8"
  ```
Run `claude` in your terminal.

---

### 2. Claude Code (VS Code Extension)
Search for **Claude Code** in VS Code Extensions and install it.

If required by your environment, install the local Claude CLI program once:
- **macOS / Linux / WSL**:
  ```bash
  curl -fsSL https://claude.ai/install.sh | bash
  ```
- **Windows (PowerShell)**:
  ```powershell
  irm https://claude.ai/install.ps1 | iex
  ```
Verify installation: `claude --version`.

Open VS Code User Settings JSON (`Cmd/Ctrl+Shift+P` → `Open User Settings (JSON)`) and append:
```json
{
  "claudeCode.environmentVariables": [
    { "name": "ANTHROPIC_AUTH_TOKEN", "value": "sk-your-api-key" },
    { "name": "ANTHROPIC_BASE_URL", "value": "https://agentrouter.org" },
    { "name": "CLAUDE_CODE_ENABLE_GATEWAY_MODEL_DISCOVERY", "value": "1" },
    { "name": "ANTHROPIC_MODEL", "value": "claude-opus-4-8" }
  ],
  "claudeCode.disableLoginPrompt": true,
  "claudeCode.initialPermissionMode": "acceptEdits"
}
```

> 💡 **Tip**: On the Claude Code extension details page, click the Gear icon → Settings and enable **Disable Login Prompt** to prevent official login box popups.

---

### 3. Claude App (Desktop)
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

### 4. Cline (VS Code Extension)
In Cline extension settings panel:
- **For Claude Models**:
  - API Provider: `Anthropic`
  - Use custom base URL: `Checked`
  - Base URL: `https://agentrouter.org` (NO `/v1`)
  - API Key: `sk-your-api-key`
  - Model: `claude-opus-4-8`
- **For GPT Models**:
  - API Provider: `OpenAI Compatible`
  - Base URL: `https://agentrouter.org/v1`
  - API Key: `sk-your-api-key`
  - Model: `gpt-5.6-sol`

---

### 4. Cursor IDE
1. Open Cursor Settings (`Cmd/Ctrl + ,`) → **Models**.
2. Paste your AgentRouter key under **OpenAI API Key**.
3. Toggle on **Override OpenAI Base URL**.
4. Set Base URL to `https://agentrouter.org/v1`.
5. Select `gpt-5.6-sol` or `claude-opus-4-8`.

---

### 5. Roo Code / Kilo Code
Add a new Provider Profile in extension settings:
- **Claude Models**: Provider = `Anthropic`, Base URL = `https://agentrouter.org`
- **GPT Models**: Provider = `OpenAI Compatible`, Base URL = `https://agentrouter.org/v1`
