// Package prom provides function to create prometheus metrics
package prom

import(
  "time"
  "github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promauto"
  "personApi/models"
)


var (

	//Counter
	OpsProcessedCount = promauto.NewCounter(prometheus.CounterOpts{
	  Name: "person_api_server_processed_ops_total",
	  Help: "The total number of processed events",
	})

	//Counter
	RequestCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "person_api_requests_total",
			Help: "Total number of requests received by the API",
		},[]string{"endpoint", "method"},
	)

	//Guage
	PeopleCount = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "person_api_people_total",
			Help: "Current number of people in the database",
		},
	)

	//Histogram
	RequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "person_api_request_duration_seconds",
			Help: "Time taken to process the request",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"endpoint", "method"},
	)

	//ErrorCounter
	ErrorCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "person_api_errors_total",
			Help: "Total number of errors encountered by the API",
		},
		[]string{"endpoint", "method", "error_type"},
	)

)


// OpsProcessed function creates metrics and increments it every 2 seconds
func OpsProcessed() {
	go func() {
	  for {
		  OpsProcessedCount.Inc()
		  time.Sleep(2 * time.Second)
	  }
	}()
}

// UpdatePeopleCount function updates the people count metric
func UpdatePeopleCount(people []models.Person) {
	PeopleCount.Set(float64(len(people)))
}

// UpdateRequestDuration function updates the request duration metric
func UpdateRequestDuration(method, endpoint string, start time.Time) {
	duration := time.Since(start).Seconds()
	RequestDuration.WithLabelValues(endpoint, method).Observe(duration)
}

// UpdateErrorCount function updates the error count metric
func UpdateErrorCount(method, endpoint, errorType string) {
	ErrorCount.WithLabelValues(endpoint, method, errorType).Inc()
}

