(require '[clj-http.client :as client])

(client/post "https://will-x86.com/api/pixels" {:headers {:User-Agent "curl/7.68.0"
                                                          :Accept "*/*"}
                                                :content-type :json
                                                :form-params {:x 1
                                                              :y 2
                                                              :hex "#102932"}})
