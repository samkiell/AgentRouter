// Install: npm install @anthropic-ai/sdk
// Run:     node claude.js
// Model:   claude-opus-4-8 | claude-opus-5

const Anthropic = require("@anthropic-ai/sdk");

const client = new Anthropic({
  apiKey: process.env.AGENTROUTER_API_KEY,
  baseURL: "https://agentrouter.org",
  defaultHeaders: {
    Originator: "codex_cli_rs",
    Version: "0.101.0",
    "User-Agent": "codex_cli_rs/0.101.0 (Mac OS 26.0.1; arm64) Apple_Terminal/464",
  },
});

async function main() {
  const message = await client.messages.create({
    model: "claude-opus-4-8", // or "claude-opus-5"
    max_tokens: 256,
    messages: [{ role: "user", content: "Hi, who are you?" }],
  });

  console.log(message.content[0].text);
}

main().catch(console.error);
