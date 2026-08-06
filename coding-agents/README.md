# Option 2: Coding Agents & IDE Integration

Integrate AgentRouter directly with official developer tools, CLI agents, and IDE extensions. **Confirmed Working ✅**

## Supported Tools

1. **Claude Code CLI** (`npm install -g @anthropic-ai/claude-code`)
   - `export ANTHROPIC_BASE_URL=https://agentrouter.org`
   - `export ANTHROPIC_AUTH_TOKEN=your_key`
   - `export ANTHROPIC_API_KEY=your_key`
   - `export ANTHROPIC_MODEL=claude-opus-4-8`
2. **Claude Code VS Code Extension**
   - Configured via `claudeCode.environmentVariables` in user settings JSON.
3. **Cline VS Code Extension**
   - Anthropic provider (`https://agentrouter.org`) or OpenAI Compatible provider (`https://agentrouter.org/v1`).
4. **Cursor IDE**
   - Override OpenAI Base URL to `https://agentrouter.org/v1`.
5. **Roo Code / Kilo Code**
   - Provider profile configured for custom base URL.

## Crucial Rule
- **Claude / Anthropic style**: `https://agentrouter.org` (NO `/v1`)
- **OpenAI style**: `https://agentrouter.org/v1` (WITH `/v1`)
