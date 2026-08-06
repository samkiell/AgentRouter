# Why the extra headers?

AgentRouter sits behind an Aliyun / Cloudflare WAF that checks more than just your API key.

👉 **Test WAF Bypass Live**: **[agentrouter.samkiel.online](https://agentrouter.samkiel.online)**

---

It looks at the client fingerprint. Plain `curl`, standard browser fetches, and default Anthropic/OpenAI SDK calls get rejected with:

```json
{
  "error": "unauthorized client detected"
}
```

The three headers below emulate an official client (Codex CLI):

```http
Originator: codex_cli_rs
Version: 0.101.0
User-Agent: codex_cli_rs/0.101.0 (Mac OS 26.0.1; arm64) Apple_Terminal/464
```

Send them on **every** request and the WAF passes your traffic through.

> ⚠️ Note: This is an unofficial workaround. WAF checks may change at any time.
