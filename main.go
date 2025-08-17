// main.go
package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	myCustomMetric = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "my_custom_metric",
		Help: "Exemplo de métrica customizada para Prometheus",
	})
)

func recordMetrics() {
	go func() {
		for {
			myCustomMetric.Set(rand.Float64() * 100)
			time.Sleep(5 * time.Second)
		}
	}()
}

func main() {
	prometheus.MustRegister(myCustomMetric)
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
