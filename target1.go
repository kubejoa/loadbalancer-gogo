package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/target1/home", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Println("[origin server] received request at: %s\n", time.Now())
		_, _ = fmt.Fprint(rw, "LoadBalancer-GOGO target1 server response")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}