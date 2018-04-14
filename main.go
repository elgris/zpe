package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/elgris/zpe/client"
	"github.com/elgris/zpe/client/operations"
	"github.com/elgris/zpe/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/kr/pretty"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	yaml "gopkg.in/yaml.v2"
)

type AppConfig struct {
	Listen    string        `yaml:"listen"`
	ZipkinURL string        `yaml:"zipkin_url"`
	Period    time.Duration `yaml:"period"`
	Queries   map[string]struct {
		Query            string    `yaml:"query"`
		ServiceName      string    `yaml:"service_name"`
		HistogramBuckets []float64 `yaml:"histogram_buckets"`
	} `yaml:"queries"`
}

func main() {
	configFilePath := flag.String("config", "config.yml", "path to config file")
	flag.Parse()

	configFile, err := ioutil.ReadFile(*configFilePath)
	if err != nil {
		log.Fatalf("could not read configuration file '%s': %s", *configFilePath, err.Error())
	}

	config := AppConfig{}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("could not unmarshal configuration from file '%s': %s", *configFilePath, err.Error())
	}

	transport := httptransport.New(config.ZipkinURL, "/api/v2", []string{"http"})
	zipkin := client.New(transport, strfmt.Default)

	for metric, queryConfig := range config.Queries {
		scraperConfig := ScraperConfig{
			MetricName:  metric,
			ServiceName: queryConfig.ServiceName,
			Query:       queryConfig.Query,
			DurationsHistogramBuckets: queryConfig.HistogramBuckets,
			Period:     config.Period,
			Client:     zipkin,
			QueryLimit: 1000,
		}
		runScraper(scraperConfig)
	}

	http.Handle("/metrics", promhttp.Handler())
	log.Printf("Starting server with %d zipkin scrapers on %s", len(config.Queries), config.Listen)

	log.Fatal(http.ListenAndServe(config.Listen, nil))
}

type ScraperConfig struct {
	MetricName                string
	ServiceName               string
	Query                     string
	Period                    time.Duration
	Client                    *client.Zipkin
	QueryLimit                int64
	DurationsHistogramBuckets []float64
}

func runScraper(config ScraperConfig) {
	traceDurationsHistogram := prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    config.MetricName + "_durations_histogram_ms",
		Help:    "Latency distributions for " + config.MetricName,
		Buckets: config.DurationsHistogramBuckets,
	})
	traceCounter := prometheus.NewCounter(prometheus.CounterOpts{
		Name: config.MetricName + "_counter",
		Help: "Counter for traces for" + config.MetricName,
	})

	prometheus.MustRegister(traceDurationsHistogram)
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
					pretty.Println("Duration", traceDuration)
					traceDurationsHistogram.Observe(float64(traceDuration))
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
