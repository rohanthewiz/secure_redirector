// Standalone Http to Https redirector
// Credit: https://gist.github.com/d-schmidt/587ceec34ce1334a5e60
package main

import (
	"log"
	"net/http"
)

func main() {
	// redirect every http request to https
	log.Fatal(http.ListenAndServe(":80",
		http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			target := "https://ccswm.org"
			if len(req.URL.RawQuery) > 0 {
				target += "?" + req.URL.RawQuery
			}
			//fmt.Printf("redirecting to: %s", target)
			http.Redirect(w, req, target, http.StatusPermanentRedirect)
		})))

	// serve index (and anything else) as https
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", index)
	//http.ListenAndServeTLS(":443", "cert.pem", "key.pem", mux)
}

//func index(w http.ResponseWriter, req *http.Request) {
//	// all calls to unknown url paths should return 404
//	if req.URL.Path != "/" {
//		log.Printf("404: %s", req.URL.String())
//		http.NotFound(w, req)
//		return
//	}
//	http.ServeFile(w, req, "index.html")
//}
