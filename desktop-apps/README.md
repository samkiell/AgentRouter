# Option 3: Desktop Applications Integration

Integrate AgentRouter into desktop AI clients (Claude Desktop, Cowork, Cursor, etc.). **Confirmed Working ✅**

## Endpoint Rules
- **Claude / Anthropic Style Apps**: Set Base URL to `https://agentrouter.org` (NO `/v1`).
- **OpenAI Style Apps**: Set Base URL to `https://agentrouter.org/v1` (WITH `/v1`).

## Claude Desktop Setup
1. Open Claude Desktop settings.
2. Under Custom Provider / API Settings:
   - Set Base URL to `https://agentrouter.org`
   - Set API Key to your AgentRouter key (`sk-...`)
   - Select Model: `claude-opus-4-8` or `claude-opus-5`
