<?php
// Run: php openai.php

$apiKey = getenv('AGENTROUTER_API_KEY');

if (!$apiKey) {
    die("Error: AGENTROUTER_API_KEY environment variable is not set.\n");
}

$ch = curl_init('https://agentrouter.org/v1/chat/completions');

$data = [
    'model' => 'gpt-5.6-sol',
    'messages' => [
        ['role' => 'user', 'content' => 'Hi, who are you?']
    ],
    'max_tokens' => 256
];

curl_setopt_array($ch, [
    CURLOPT_POST => true,
    CURLOPT_RETURNTRANSFER => true,
    CURLOPT_HTTPHEADER => [
        'Authorization: Bearer ' . $apiKey,
        'Content-Type: application/json',
        'Originator: codex_cli_rs',
        'Version: 0.101.0',
        'User-Agent: codex_cli_rs/0.101.0 (Mac OS 26.0.1; arm64) Apple_Terminal/464'
    ],
    CURLOPT_POSTFIELDS => json_encode($data)
]);

$response = curl_exec($ch);

if (curl_errno($ch)) {
    echo 'cURL error: ' . curl_error($ch) . "\n";
    curl_close($ch);
    exit(1);
}

curl_close($ch);

$result = json_decode($response, true);
echo $result['choices'][0]['message']['content'] ?? $response;
echo "\n";
