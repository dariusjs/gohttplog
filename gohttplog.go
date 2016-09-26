package main

import (
    "fmt"
    "net/http"
    "log"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Time is: %s", time.Now())
}

func Log(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
        handler.ServeHTTP(w, r)
    })
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
    })
    http.ListenAndServe(":8080", Log(http.DefaultServeMux))
    log.Fatal(http.ListenAndServe(":12345", nil))
}
