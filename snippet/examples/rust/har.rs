use serde_json::json;
use reqwest;

#[tokio::main]
pub async fn main() {
    let url = "http://localhost:8081/api/pixels";

    let payload = json!({
        "x": 1,
        "y": 2,
        "hex": "#102932"
    });

    let mut headers = reqwest::header::HeaderMap::new();
    headers.insert("Content-Type", "application/json".parse().unwrap());
    headers.insert("User-Agent", "curl/7.68.0".parse().unwrap());
    headers.insert("Accept", "*/*".parse().unwrap());

    let client = reqwest::Client::new();
    let response = client.post(url)
        .headers(headers)
        .json(&payload)
        .send()
        .await;

    let results = response.unwrap()
        .json::<serde_json::Value>()
        .await
        .unwrap();

    dbg!(results);
}
