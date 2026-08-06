# Why the extra headers?

AgentRouter sits behind an Aliyun WAF that checks more than just your API key.

It looks at the client fingerprint. Plain `curl`, most Node fetch calls, and generic SDKs get rejected with:

```
unauthorized client detected
```

The three headers below make the request look like it comes from the official Codex CLI (one of the allowed clients):

```
Originator: codex_cli_rs
Version: 0.101.0
User-Agent: codex_cli_rs/0.101.0 (Mac OS 26.0.1; arm64) Apple_Terminal/464
```

Send them on **every** request and the WAF lets you through.

This is not official support. It can stop working if AgentRouter updates their checks.
