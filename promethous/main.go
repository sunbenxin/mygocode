package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var apnsTotalCounter prometheus.Counter

func main() {
	apnsTotalCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "putong",
		Subsystem: "push",
		Name:      "apns_total_count",
		Help:      "apns_total_counter",
	})
	prometheus.Register(apnsTotalCounter)

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
	}
	http.Handle("/debug/metrics", prometheus.Handler())
	go http.Serve(l, nil)

	/*
		tick := time.Tick(time.Second)
		for range tick {
			apnsTotalCounter.Inc()
		}
	*/
	apnsTotalCounter.Add(float64(time.Now().Unix()))
	time.Sleep(time.Hour)
}
