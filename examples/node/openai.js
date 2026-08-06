const OpenAI = require("openai");

const client = new OpenAI({
  apiKey: process.env.AGENTROUTER_API_KEY, // 
  baseURL: "https://agentrouter.org/v1",
  defaultHeaders: {
    Originator: "codex_cli_rs",
    Version: "0.101.0",
    "User-Agent": "codex_cli_rs/0.101.0 (Mac OS 26.0.1; arm64) Apple_Terminal/464",
  },
});

async function main() {
  const completion = await client.chat.completions.create({
    model: "gpt-5.6-sol",
    messages: [{ role: "user", content: "Hi, who are you?" }],
    max_tokens: 256,
  });

  console.log(completion.choices[0].message.content);
}

main().catch(console.error);
