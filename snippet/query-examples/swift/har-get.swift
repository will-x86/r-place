import Foundation

let headers = [
  "User-Agent": "curl/7.68.0",
  "Accept": "*/*"
]

let request = NSMutableURLRequest(url: NSURL(string: "https://will-x86.com/api/pixelsq?x=1&y=2&hex=%23102932")! as URL,
                                        cachePolicy: .useProtocolCachePolicy,
                                    timeoutInterval: 10.0)
request.httpMethod = "GET"
request.allHTTPHeaderFields = headers

let session = URLSession.shared
let dataTask = session.dataTask(with: request as URLRequest, completionHandler: { (data, response, error) -> Void in
  if (error != nil) {
    print(error)
  } else {
    let httpResponse = response as? HTTPURLResponse
    print(httpResponse)
  }
})

dataTask.resume()
