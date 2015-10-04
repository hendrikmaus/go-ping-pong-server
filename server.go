package main

import (
	"fmt"
	"flag"
	"log"
	"net/http"
)

func main() {
	listen := flag.String("listen", ":8080", "Provide address:port to listen on")
	flag.Parse();

	log.Printf("Starting server on %s", *listen)
	http.HandleFunc("/", httpHandleRoot)
	http.HandleFunc("/ping", httpHandlePing)
	log.Fatal(http.ListenAndServe(*listen, nil))
}

func httpHandleRoot(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "/ping", http.StatusSeeOther)
}

func httpHandlePing(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("X-Powered-By", "golang")
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, "pong")
}
