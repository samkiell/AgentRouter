# Install: pip install openai
# Run:     python openai.py

from openai import OpenAI
import os

client = OpenAI(
    api_key=os.getenv("AGENTROUTER_API_KEY"),
    base_url="https://agentrouter.org/v1",
    default_headers={
        "Originator": "codex_cli_rs",
        "Version": "0.101.0",
        "User-Agent": "codex_cli_rs/0.101.0 (Mac OS 26.0.1; arm64) Apple_Terminal/464",
    },
)

completion = client.chat.completions.create(
    model="gpt-5.6-sol",
    messages=[{"role": "user", "content": "Hi, who are you?"}],
    max_tokens=256,
)

print(completion.choices[0].message.content)
