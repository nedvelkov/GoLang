package main

import (
	"io"
	"net/http"
	"os"
)

func WebFunc() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3854", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./menu.txt")
	io.Copy(w, f)
}
