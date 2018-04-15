package main

import (
	"log"
	"sort"
	"time"

	"github.com/elgris/zpe/client"
	"github.com/elgris/zpe/client/operations"
	"github.com/elgris/zpe/models"
	"github.com/prometheus/client_golang/prometheus"
)

// ScraperConfig is a piece of configuration for Zipkin scraper
// Scraper will query Zipkin periodically (specified by Period) and will get data
// for given ServiceName, constrained by Query.
// MetricName is used for exposing Prometheus metrics
type ScraperConfig struct {
	MetricName  string
	ServiceName string
	Query       string
	Period      time.Duration
	Client      *client.Zipkin
	QueryLimit  int64
}

// Start scraper in background using given configuration
func RunScraper(config ScraperConfig) {
	traceDurationsSummary := prometheus.NewSummary(prometheus.SummaryOpts{
		Name: config.MetricName + "_durations_summary",
		Help: "Latency distributions for " + config.MetricName,
	})
	traceCounter := prometheus.NewCounter(prometheus.CounterOpts{
		Name: config.MetricName + "_counter",
		Help: "Counter for traces for" + config.MetricName,
	})

	prometheus.MustRegister(traceDurationsSummary)
	prometheus.MustRegister(traceCounter)
	go func() {
		previousEndTimestampMs := (time.Now().UnixNano() - config.Period.Nanoseconds()) / 1e6

		for {
			endTimestampMs := time.Now().UnixNano() / 1e6
			lookbackMs := endTimestampMs - previousEndTimestampMs

			params := operations.NewGetTracesParams()
			params.ServiceName = &config.ServiceName
			params.AnnotationQuery = &config.Query
			params.Lookback = &lookbackMs
			params.EndTs = &endTimestampMs
			params.Limit = &config.QueryLimit

			response, err := config.Client.Operations.GetTraces(params)
			if err == nil {
				for _, trace := range response.Payload {
					traceDuration := Trace(trace).Duration()
					traceDurationsSummary.Observe(float64(traceDuration))
					traceCounter.Inc()
				}
			} else {
				log.Printf("[ERROR] could not fetch the traces: %s", err.Error())
			}

			<-time.After(config.Period)
			// TODO: add stop by command

			previousEndTimestampMs = endTimestampMs
		}
	}()
}

type Trace models.Trace

func (t Trace) Len() int           { return len(t) }
func (t Trace) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t Trace) Less(i, j int) bool { return t[i].Timestamp < t[j].Timestamp }

func (t Trace) Duration() int64 {
	if len(t) == 0 {
		return 0
	}

	if len(t) == 1 {
		return t[0].Duration / 1000
	}

	sort.Sort(t) // TODO: don't call sort.Sort every time
	last := len(t) - 1

	return (t[last].Timestamp - t[0].Timestamp + t[last].Duration) / 1000
}
