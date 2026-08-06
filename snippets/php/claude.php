<?php
// Run:   php claude.php
// Model: claude-opus-4-8 | claude-opus-5

$apiKey = getenv('AGENTROUTER_API_KEY');

if (!$apiKey) {
    die("Error: AGENTROUTER_API_KEY not set.\n");
}

$ch = curl_init('https://agentrouter.org/v1/messages');

$data = [
    'model'      => 'claude-opus-4-8', // or 'claude-opus-5'
    'max_tokens' => 256,
    'messages'   => [
        ['role' => 'user', 'content' => 'Hi, who are you?']
    ],
];

curl_setopt_array($ch, [
    CURLOPT_POST            => true,
    CURLOPT_RETURNTRANSFER  => true,
    CURLOPT_HTTPHEADER      => [
        'Authorization: Bearer ' . $apiKey,
        'x-api-key: ' . $apiKey,
        'Content-Type: application/json',
        'anthropic-version: 2023-06-01',
        'Originator: codex_cli_rs',
        'Version: 0.101.0',
        'User-Agent: codex_cli_rs/0.101.0 (Mac OS 26.0.1; arm64) Apple_Terminal/464',
    ],
    CURLOPT_POSTFIELDS      => json_encode($data),
]);

$response = curl_exec($ch);

if (curl_errno($ch)) {
    echo 'cURL error: ' . curl_error($ch) . "\n";
    curl_close($ch);
    exit(1);
}

curl_close($ch);

$result = json_decode($response, true);
echo $result['content'][0]['text'] ?? $response;
echo "\n";
