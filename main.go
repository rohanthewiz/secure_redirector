// Standalone Http to Https redirector
// Credit: https://gist.github.com/d-schmidt/587ceec34ce1334a5e60
package main

import (
	"fmt"
	"log"
	"net/http"
)

func redirect(w http.ResponseWriter, req *http.Request) {
	// drop non-default ports
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	fmt.Printf("redirect to: %s", target)
	http.Redirect(w, req, target, http.StatusPermanentRedirect)
}


func main() {
	// redirect every http request to https
	log.Fatal(http.ListenAndServe(":80", http.HandlerFunc(redirect)))

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
