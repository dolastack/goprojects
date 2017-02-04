package main

import (
	"goprojects/myweb"
	"net/http"
)

func main() {

	http.HandleFunc("/", myweb.ServeIndex)
	http.ListenAndServe(":3000", nil)
}
