package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"runtime"
	"time"
)

var OS = runtime.GOOS

type PortNumHandler struct {
	PortNum string
}

func (pn *PortNumHandler) Handler(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("[reverse proxy server] received request at: %s\n", time.Now())

	originServerURL, err := url.Parse("http://localhost:" + pn.PortNum)

	if err != nil {
		log.Fatal("invalid origin serer URL")
	}

	// set req Host, URL and Request URI to foward a request to the origin server
	req.Host = originServerURL.Host
	req.URL.Host = originServerURL.Host
	req.URL.Scheme = originServerURL.Scheme
	req.RequestURI = ""

	// save the response from the origin server
	originServerResponse, err := http.DefaultClient.Do(req)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprint(rw, err)
		return
	}

	// return response to the client

	switch OS {
	case "windows":
		OS = "Windows"
	case "darwin":
		OS = "macOS"
	case "linux":
		OS = "Linux"
	}

	rw.Header().Set("Server", fmt.Sprintf("LoadBalancer-GOGO (%s)", OS))
	rw.WriteHeader(http.StatusOK)
	io.Copy(rw, originServerResponse.Body)
}