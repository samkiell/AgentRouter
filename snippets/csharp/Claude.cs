// Run:   dotnet run
// Model: claude-opus-5 | claude-opus-4-8

using System;
using System.Net.Http;
using System.Text;
using System.Text.Json;
using System.Threading.Tasks;

class Claude
{
    static async Task Main(string[] args)
    {
        string apiKey = Environment.GetEnvironmentVariable("AGENTROUTER_API_KEY");
        if (string.IsNullOrEmpty(apiKey))
        {
            Console.WriteLine("Error: AGENTROUTER_API_KEY not set.");
            return;
        }

        using var client = new HttpClient();

        var payload = new
        {
            model = "claude-opus-5", // or "claude-opus-4-8"
            max_tokens = 256,
            messages = new[]
            {
                new { role = "user", content = "Hi, who are you?" }
            }
        };

        var request = new HttpRequestMessage(HttpMethod.Post, "https://agentrouter.org/v1/messages")
        {
            Content = new StringContent(JsonSerializer.Serialize(payload), Encoding.UTF8, "application/json")
        };

        request.Headers.Add("Authorization", $"Bearer {apiKey}");
        request.Headers.Add("x-api-key", apiKey);
        request.Headers.Add("anthropic-version", "2023-06-01");
        request.Headers.Add("Originator", "codex_cli_rs");
        request.Headers.Add("Version", "0.101.0");
        request.Headers.UserAgent.ParseAdd("codex_cli_rs/0.101.0 (Mac OS 26.0.1; arm64) Apple_Terminal/464");

        var response = await client.SendAsync(request);
        string responseBody = await response.Content.ReadAsStringAsync();

        if (!response.IsSuccessStatusCode)
        {
            Console.WriteLine($"Error {response.StatusCode}: {responseBody}");
            return;
        }

        using var doc = JsonDocument.Parse(responseBody);
        string text = doc.RootElement
            .GetProperty("content")[0]
            .GetProperty("text")
            .GetString();

        Console.WriteLine(text);
    }
}
