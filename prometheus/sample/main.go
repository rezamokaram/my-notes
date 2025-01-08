package main

import (
    "fmt"
    "net/http"
    "time"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    httpRequestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "handler", "code"},
    )

    httpRequestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "http_request_duration_seconds",
            Help:    "Duration of HTTP requests",
            Buckets: prometheus.DefBuckets,
        },
        []string{"method", "handler"},
    )

    httpRequestSize = prometheus.NewSummaryVec(
        prometheus.SummaryOpts{
            Name:       "http_request_size_bytes",
            Help:       "Size of HTTP requests",
            Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
        },
        []string{"method", "handler"},
    )

    httpResponseSize = prometheus.NewSummaryVec(
        prometheus.SummaryOpts{
            Name:       "http_response_size_bytes",
            Help:       "Size of HTTP responses",
            Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
        },
        []string{"method", "handler"},
    )
)

func init() {
    prometheus.MustRegister(httpRequestsTotal, httpRequestDuration, httpRequestSize, httpResponseSize)
}

func instrumentHandler(handlerName string, handler http.Handler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        // Wrap the ResponseWriter to capture the status code and size
        wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
        handler.ServeHTTP(wrapped, r)

        duration := time.Since(start).Seconds()
        httpRequestsTotal.WithLabelValues(r.Method, handlerName, fmt.Sprint(wrapped.statusCode)).Inc()
        httpRequestDuration.WithLabelValues(r.Method, handlerName).Observe(duration)
        httpRequestSize.WithLabelValues(r.Method, handlerName).Observe(float64(r.ContentLength))
        httpResponseSize.WithLabelValues(r.Method, handlerName).Observe(float64(wrapped.size))
    }
}

type responseWriter struct {
    http.ResponseWriter
    statusCode int
    size       int
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.statusCode = code
    rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
    size, err := rw.ResponseWriter.Write(b)
    rw.size += size
    return size, err
}

func main() {
    http.Handle("/", instrumentHandler("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })))

    http.Handle("/hello", instrumentHandler("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, Go!")
    })))

    http.Handle("/metrics", promhttp.Handler())

    http.ListenAndServe(":8080", nil)
}
