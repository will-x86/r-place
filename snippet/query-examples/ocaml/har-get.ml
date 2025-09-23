open Cohttp_lwt_unix
open Cohttp
open Lwt

let uri = Uri.of_string "https://will-x86.com/api/pixelsq?x=1&y=2&hex=%23102932" in
let headers = Header.add_list (Header.init ()) [
  ("User-Agent", "curl/7.68.0");
  ("Accept", "*/*");
] in

Client.call ~headers `GET uri
>>= fun (res, body_stream) ->
  (* Do stuff with the result *)
