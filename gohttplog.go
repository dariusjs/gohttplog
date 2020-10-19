package main

import (
	"expvar"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/paulbellamy/ratecounter"
)

var (
	counter       *ratecounter.RateCounter
	hitsperminute = expvar.NewInt("hits_per_minute")
)

func increment(w http.ResponseWriter, r *http.Request) {
	counter.Incr(1)
	hitsperminute.Set(counter.Rate())
	io.WriteString(w, strconv.FormatInt(counter.Rate(), 10))
}

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
	counter = ratecounter.NewRateCounter(1 * time.Minute)
	http.HandleFunc("/increment", increment)
	http.ListenAndServe(":8080", Log(http.DefaultServeMux))
	log.Fatal(http.ListenAndServe(":12345", nil))
}
