# Option 2: Coding Agents & IDE Integration

Integrate AgentRouter directly into official coding agents, developer CLI tools, and IDE extensions. **Confirmed Working ✅**

## Supported Tools & Setup Guide

### 1. Claude Code CLI
Install the official package:
```bash
npm install -g @anthropic-ai/claude-code
```

Set environment variables:
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
Open VS Code User Settings JSON (`Cmd/Ctrl+Shift+P` → `Open User Settings (JSON)`) and add:
```json
{
  "claudeCode.environmentVariables": [
    { "name": "ANTHROPIC_AUTH_TOKEN", "value": "sk-your-api-key" },
    { "name": "ANTHROPIC_BASE_URL", "value": "https://agentrouter.org" },
    { "name": "ANTHROPIC_MODEL", "value": "claude-opus-4-8" }
  ],
  "claudeCode.disableLoginPrompt": true
}
```

---

### 3. Cline (VS Code Extension)
In Cline extension settings:
- **For Claude Models**:
  - API Provider = `Anthropic`
  - Base URL = `https://agentrouter.org` (NO `/v1`)
  - API Key = `sk-your-api-key`
  - Model = `claude-opus-4-8` or `claude-opus-5`
- **For GPT Models**:
  - API Provider = `OpenAI Compatible`
  - Base URL = `https://agentrouter.org/v1`
  - API Key = `sk-your-api-key`
  - Model = `gpt-5.6-sol`

---

### 4. Cursor IDE
1. Open Settings → Models.
2. Paste your AgentRouter key under OpenAI API Key.
3. Enable "Override OpenAI Base URL".
4. Set Base URL to `https://agentrouter.org/v1`.

---

### 5. Roo Code / Kilo Code
Create a custom provider profile:
- **Anthropic Base URL**: `https://agentrouter.org` (NO `/v1`)
- **OpenAI Base URL**: `https://agentrouter.org/v1`
