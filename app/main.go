package main

import (
    "fmt"
    "net/http"
    "sync/atomic"
)

var requestCount atomic.Int64

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Request received:", r.URL.Path)
    requestCount.Add(1)
    fmt.Fprintf(w, "Hello! My KEDA demo application is running!\nTotal requests: %d\n", requestCount.Load())
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain; version=0.0.4")
    fmt.Fprintf(w, "app_requests_total %d\n", requestCount.Load())
}

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/metrics", metricsHandler)

    fmt.Println("Server listening on port 8080")

    http.ListenAndServe(":8080", nil)
}