open Cohttp_lwt_unix
open Cohttp
open Lwt

let uri = Uri.of_string "http://localhost:8081/api/pixels" in
let headers = Header.add_list (Header.init ()) [
  ("Content-Type", "application/json");
  ("User-Agent", "curl/7.68.0");
  ("Accept", "*/*");
] in
let body = Cohttp_lwt_body.of_string "{\"x\":1, \"y\":2, \"hex\":\"#102932\"}" in

Client.call ~headers ~body `POST uri
>>= fun (res, body_stream) ->
  (* Do stuff with the result *)