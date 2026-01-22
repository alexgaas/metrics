## metrics
Production-proven framework to manage *prometheus* metrics for your *Go* service. 

Framework supports:
- _counters_ (monotonically increasing values)
- _gauges_ (tracks single float64 value) and _timers_ (measures gauge duration)
- _histograms_



#### Example of using _metrics_ framework

Add metric value:
```go
r := prometheus.NewRegistry(NewRegistryOpts())

cnt := r.Counter("mycounter")
cnt.Add(42)

gg := r.Gauge("mygauge")
gg.Set(2.4)

hs := r.Histogram("myhistogram", metrics.NewBuckets(1, 2, 3))
hs.RecordValue(0.5)
hs.RecordValue(1.5)
hs.RecordValue(1.7)
hs.RecordValue(2.2)
hs.RecordValue(42)
```

( you will need to expose yours _NewHandler_ to use this example):
```go
// chi router
func Example_chi() {
	// Create registry.
	registry := prometheus.NewRegistry(prometheus.NewRegistryOpts())

	// Create HTTP router.
	r := chi.NewMux()

	// Collect http metrics.
	r.Use(httpmetrics.New(registry.WithPrefix("http"), httpmetrics.WithPathEndpoint()))

	// Expose metrics to fetcher.
	r.Handle("/metrics", NewHandler(registry))
}

// stdlib http
func Example_stdlib() {
	// Create registry.
	registry := prometheus.NewRegistry(prometheus.NewRegistryOpts())

	middleware := httpmetrics.New(registry.WithPrefix("http"), httpmetrics.WithPathEndpoint())

	myHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte("Hello"))
	})

	http.Handle("/endpoint", middleware(myHandler))

	// Expose metrics to fetcher.
	http.Handle("/metrics", NewHandler(registry))
}
```
