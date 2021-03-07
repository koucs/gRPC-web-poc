package main

import (
    "fmt"
    "log"
    "net/http"
    "html"
)

func main() {
    http.HandleFunc("/", hello)
    fmt.Println("Server started")
    log.Fatal(http.ListenAndServe(":9090", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"message":"hello world!"}`))
	fmt.Println("Hello, %q", html.EscapeString(r.URL.Path))
}
