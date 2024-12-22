// Package prom provides function to create prometheus metrics
package prom

import(
  "time"
  "github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promauto"
)

// MetricsEndpoint function creates a counter metric and increments it every 2 seconds
func MetricsEndpoint() {
	go func() {
	  for {
		  opsProcessed.Inc()
		  time.Sleep(2 * time.Second)
	  }
	}()
}

// opsProcessed is a counter metric that counts the number of processed events
var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
	  Name: "person_api_server_processed_ops_total",
	  Help: "The total number of processed events",
	})
)