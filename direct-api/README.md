# ❌ Option 1: Direct API (Raw HTTP / cURL)

> **Status**: Blocked by Design / Requires Client Spoofing ❌

Directly calling AgentRouter's base URL from custom code or standard HTTP clients via raw HTTP/cURL is **blocked by Cloudflare/Aliyun WAF design**.

---

## 🛑 Why Direct API Calls Fail

AgentRouter's WAF inspects the client fingerprint. Standard HTTP libraries (e.g. `fetch`, unconfigured `curl`, default Anthropic SDKs) return:

```json
{
  "error": "unauthorized client detected"
}
```

---

## 🛠️ WAF Bypass Headers

To send raw HTTP/cURL requests directly to AgentRouter, you **must explicitly inject** these 3 headers on every request:

```http
Originator: codex_cli_rs
Version: 0.101.0
User-Agent: codex_cli_rs/0.101.0 (Mac OS 26.0.1; arm64) Apple_Terminal/464
```

> Read more in [`waf-bypass.md`](waf-bypass.md).

---

## 🌐 Base URL Endpoints

| Format | Base URL | Endpoint Path |
|---|---|---|
| **OpenAI Style** | `https://agentrouter.org/v1` | `/chat/completions` |
| **Claude Native Style** | `https://agentrouter.org` | `/v1/messages` |

---

## 💻 Full Standalone Code Examples

### 1. cURL (OpenAI Chat Completions)
```bash
curl -X POST https://agentrouter.org/v1/chat/completions \
  -H "Authorization: Bearer sk-your-api-key" \
  -H "Content-Type: application/json" \
  -H "Originator: codex_cli_rs" \
  -H "Version: 0.101.0" \
  -H "User-Agent: codex_cli_rs/0.101.0 (Mac OS 26.0.1; arm64) Apple_Terminal/464" \
  -d '{
    "model": "gpt-5.6-sol",
    "messages": [{"role": "user", "content": "Hello world"}]
  }'
```

### 2. cURL (Claude Messages API)
```bash
curl -X POST https://agentrouter.org/v1/messages \
  -H "Authorization: Bearer sk-your-api-key" \
  -H "x-api-key: sk-your-api-key" \
  -H "Content-Type: application/json" \
  -H "anthropic-version: 2023-06-01" \
  -H "Originator: codex_cli_rs" \
  -H "Version: 0.101.0" \
  -H "User-Agent: codex_cli_rs/0.101.0 (Mac OS 26.0.1; arm64) Apple_Terminal/464" \
  -d '{
    "model": "claude-opus-4-8",
    "max_tokens": 256,
    "messages": [{"role": "user", "content": "Hello world"}]
  }'
```

### 3. Node.js (Native HTTP - Bypasses SDK Envelope Issues)
```javascript
const https = require("https");

const apiKey = process.env.AGENTROUTER_API_KEY || "sk-your-api-key";

const body = JSON.stringify({
  model: "claude-opus-4-8",
  max_tokens: 256,
  messages: [{ role: "user", content: "Hello from Node.js" }],
});

const req = https.request({
  hostname: "agentrouter.org",
  path: "/v1/messages",
  method: "POST",
  headers: {
    "Authorization": `Bearer ${apiKey}`,
    "x-api-key": apiKey,
    "Content-Type": "application/json",
    "anthropic-version": "2023-06-01",
    "Originator": "codex_cli_rs",
    "Version": "0.101.0",
    "User-Agent": "codex_cli_rs/0.101.0 (Mac OS 26.0.1; arm64) Apple_Terminal/464",
    "Content-Length": Buffer.byteLength(body),
  },
}, (res) => {
  let data = "";
  res.on("data", chunk => data += chunk);
  res.on("end", () => {
    console.log(JSON.parse(data));
  });
});

req.on("error", console.error);
req.write(body);
req.end();
```

### 4. Python (`requests`)
```python
import os
import requests

api_key = os.getenv("AGENTROUTER_API_KEY", "sk-your-api-key")

response = requests.post(
    "https://agentrouter.org/v1/chat/completions",
    headers={
        "Authorization": f"Bearer {api_key}",
        "Content-Type": "application/json",
        "Originator": "codex_cli_rs",
        "Version": "0.101.0",
        "User-Agent": "codex_cli_rs/0.101.0 (Mac OS 26.0.1; arm64) Apple_Terminal/464",
    },
    json={
        "model": "gpt-5.6-sol",
        "messages": [{"role": "user", "content": "Hello from Python"}],
    },
)

print(response.json())
```

---

## 🚨 Common Error Troubleshooting

- **`unauthorized client detected`**: You forgot one of the 3 required WAF headers. Include `Originator`, `Version`, and `User-Agent`.
- **`content-blocked`**: The prompt triggered safety filtering. Rephrase your prompt or switch between GPT / Claude models.
- **`no available channel`**: Target model is temporarily over capacity. Wait 5 seconds or switch models.
