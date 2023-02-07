package main

import (
	"fmt"
	"log"
	"net/http"
	"vmc-nsx-t-goproxy/utils"
)

var Port = "8889"

func main() {

	log.Println("VMC NSX-T Go Proxy listening on localhost:" + Port)

	proxyHandler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		_, b := utils.DoHttp(req)
		log.Println("NSX-T Go Proxy forwarding response to localhost:" + Port + "\n")
		_, _ = fmt.Fprint(rw, b)
	})

	log.Fatal(http.ListenAndServeTLS(":"+Port, "./certs/nsx_certificate.crt", "./certs/nsx_certificate.key", proxyHandler))
}
