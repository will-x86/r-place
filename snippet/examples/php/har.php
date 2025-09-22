<?php

$curl = curl_init();

curl_setopt_array($curl, [
  CURLOPT_URL => "https://place.willx86.com/api/pixels",
  CURLOPT_RETURNTRANSFER => true,
  CURLOPT_ENCODING => "",
  CURLOPT_MAXREDIRS => 10,
  CURLOPT_TIMEOUT => 30,
  CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
  CURLOPT_CUSTOMREQUEST => "POST",
  CURLOPT_POSTFIELDS => "{\"x\":1, \"y\":2, \"hex\":\"#102932\"}",
  CURLOPT_HTTPHEADER => [
    "Accept: */*",
    "Content-Type: application/json",
    "User-Agent: curl/7.68.0"
  ],
]);

$response = curl_exec($curl);
$err = curl_error($curl);

curl_close($curl);

if ($err) {
  echo "cURL Error #:" . $err;
} else {
  echo $response;
}
