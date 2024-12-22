package prom

import(
  "time"
  "github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promauto"
)

// Prometheus Metrics
func MetricsEndpoint() {
	go func() {
	  for {
		  opsProcessed.Inc()
		  time.Sleep(2 * time.Second)
	  }
	}()
}
  
var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
	  Name: "person_api_server_processed_ops_total",
	  Help: "The total number of processed events",
	})
)