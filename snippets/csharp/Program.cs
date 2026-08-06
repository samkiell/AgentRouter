// Run: dotnet run

using System;
using System.Net.Http;
using System.Text;
using System.Text.Json;
using System.Threading.Tasks;

class Program
{
    static async Task Main(string[] args)
    {
        string apiKey = Environment.GetEnvironmentVariable("AGENTROUTER_API_KEY");
        if (string.IsNullOrEmpty(apiKey))
        {
            Console.WriteLine("Error: AGENTROUTER_API_KEY environment variable is not set.");
            return;
        }

        using var client = new HttpClient();

        var payload = new
        {
            model = "gpt-5.6-sol",
            messages = new[]
            {
                new { role = "user", content = "Hi, I'm a vibecoder." }
            },
            max_tokens = 256
        };

        var request = new HttpRequestMessage(HttpMethod.Post, "https://agentrouter.org/v1/chat/completions")
        {
            Content = new StringContent(JsonSerializer.Serialize(payload), Encoding.UTF8, "application/json")
        };

        request.Headers.Add("Authorization", $"Bearer {apiKey}");
        request.Headers.Add("Originator", "codex_cli_rs");
        request.Headers.Add("Version", "0.101.0");
        request.Headers.UserAgent.ParseAdd("codex_cli_rs/0.101.0 (Mac OS 26.0.1; arm64) Apple_Terminal/464");

        var response = await client.SendAsync(request);
        string responseBody = await response.Content.ReadAsStringNodeAsync?.ToString() ?? await response.Content.ReadAsStringAsync();

        if (!response.IsSuccessStatusCode)
        {
            Console.WriteLine($"Error {response.StatusCode}: {responseBody}");
            return;
        }

        using var doc = JsonDocument.Parse(responseBody);
        string content = doc.RootElement.GetProperty("choices")[0].GetProperty("message").GetProperty("content").GetString();
        Console.WriteLine(content);
    }
}
