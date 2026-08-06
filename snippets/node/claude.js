// Install: (none — uses built-in https)
// Run:     node claude.js
// Model:   claude-opus-4-8 | claude-opus-5

const https = require("https");

const apiKey = process.env.AGENTROUTER_API_KEY;
if (!apiKey) { console.error("Error: AGENTROUTER_API_KEY not set."); process.exit(1); }

const body = JSON.stringify({
  model: "claude-opus-4-8", // or "claude-opus-5"
  max_tokens: 256,
  messages: [{ role: "user", content: "Hi, who are you?" }],
});

const req = https.request(
  {
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
  },
  (res) => {
    let data = "";
    res.on("data", (chunk) => (data += chunk));
    res.on("end", () => {
      const json = JSON.parse(data);
      console.log(json?.content?.[0]?.text ?? data);
    });
  }
);

req.on("error", console.error);
req.write(body);
req.end();
