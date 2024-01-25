package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := os.Getenv("RESPONSE")
	if len(response) == 0 {
		response = "Hello OpenShift for Developers!"
	}

	fmt.Fprintln(w, response)
	fmt.Println("Servicing an impatient beginner's request.")
}

func listenAndServe(port string) {
	fmt.Printf("serving on %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	go listenAndServe(port)

	select {}
}

var httpRequestsTotal = prometheus.NewCounter(
    prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of http requests.",
    },
)

func handler(w http.ResponseWriter, r *http.Request) {
    httpRequestsTotal.Inc()
    msg := "Received a request"
    fmt.Fprint(w, msg)
    fmt.Println(msg)
}

func main() {
    port := "8080"
    prometheus.MustRegister(httpRequestsTotal)
    http.HandleFunc("/", handler)
    http.Handle("/metrics", promhttp.Handler())
    log.Printf("Server started on port %v", port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}