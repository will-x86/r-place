(require '[clj-http.client :as client])

(client/get "https://will-x86.com/api/pixelsq" {:headers {:User-Agent "curl/7.68.0"
                                                          :Accept "*/*"}
                                                :query-params {:x "1"
                                                               :y "2"
                                                               :hex "#102932"}})
