package utils

import (
	"io"
	"log"
	"net/http"
)

// DoHttp authenticates against the CSP and performs the HTTP request returning the full response and the parsed body
func DoHttp(r *http.Request) (*http.Response, string) {

	cli := Auth()
	req, err := http.NewRequest(r.Method, cli.ServiceUrl+r.URL.Path, r.Body)
	addHttpHeader(req, "csp-auth-token", cli.AccessToken)    //CSP token
	addHttpHeader(req, "x-da-access-token", cli.AccessToken) //VCDR token
	addHttpHeader(req, "Content-Type", "application/json")

	if err != nil {
		log.Printf("New Request Failed: %s", err)
		return nil, ""
	}

	resp, err := cli.Client.Do(req)

	log.Println("HTTP " + req.Method + " " + req.URL.String())

	if err != nil {
		log.Printf("Request Failed: %s", err)
		return nil, ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println("RESPONSE body :", string(body))

	return resp, string(body)
}
