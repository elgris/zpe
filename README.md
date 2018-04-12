# Zipkin-Prometheus extractor
Small and experimental service that queries Zipkin for span data and exposes following metrics for Prometheus:
- Trace duration
- Number of traces collected.

## Configuration
Configuration should be provided as a YAML file with following structure:
```yaml
zipkin_url: 127.0.0.1:9411 # base URL pointing to Zipkin
period: 5s # how often to query data. In this case Zipkin is queried every 5 seconds
queries:   # queries to run against Zipkin. Each query is represented by:
  total_errors: # uniques metric name to be used by Prometheus
    query: "error"  # Zipkin query (e.g. "http.path=/foo/bar")
  loggable_activity:
    service_name: "log-message-processor" # and/or service name
```