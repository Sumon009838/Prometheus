package api

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_counter",
		Help: "number of ping request",
	})

func ping(w http.ResponseWriter, r *http.Request) {
	pingCounter.Inc()
	fmt.Fprintf(w, "Sumon")
}

func RunServer(Port string) {
	fmt.Println("Inside RunServer")
	prometheus.MustRegister(pingCounter)
	http.HandleFunc("/ping", ping)
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(Port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
