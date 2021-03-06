# Zipkin-Prometheus extractor
Small and experimental service that queries Zipkin for span data and exposes following metrics for Prometheus:
- Trace duration
- Number of traces collected.

## Configuration
Configuration should be provided as a YAML file with following structure:
```yaml
zipkin_url: 127.0.0.1:9411 # base URL pointing to Zipkin
listen: :8090 # what address and port should be used to expose Prometheus metrics
period: 5s # how often to query data. In this case Zipkin is queried every 5 seconds
queries:   # queries to run against Zipkin. Each query is represented by:
  total_errors: # uniques metric name to be used by Prometheus
    query: "error"  # Zipkin query (e.g. "http.path=/foo/bar")
  loggable_activity:
    service_name: "log-message-processor" # and/or service name
    histogram_buckets: [200.0, 500.0, 1000.0, 2000.0] # histogram buckets for metric duration
    # they should be manually identified and set in order to define latency SLO
```