package main

import (
	"flag"
	"io/ioutil"
	"log"
	"time"

	"github.com/elgris/zpe/client"
	"github.com/elgris/zpe/client/operations"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	yaml "gopkg.in/yaml.v2"
)

type AppConfig struct {
	ZipkinURL string        `yaml:"zipkin_url"`
	Period    time.Duration `yaml:"period"`
	Queries   map[string]struct {
		Query       string `yaml:"query"`
		ServiceName string `yaml:"service_name"`
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
			Period:      config.Period,
			Client:      zipkin,
			QueryLimit:  1000,
		}
		runScraper(scraperConfig)
	}

	// 1. Get data from Zipkin
	// 2. Expose data for prometheus

	// Dependencies:
	// - prometheus

	select {}
}

type ScraperConfig struct {
	MetricName  string
	ServiceName string // TODO: do I need ServiceName?
	Query       string
	Period      time.Duration
	Client      *client.Zipkin
	QueryLimit  int64
}

func runScraper(config ScraperConfig) {
	go func() {
		previousEndTimestampMs := (time.Now().UnixNano() - config.Period.Nanoseconds()) / int64(time.Millisecond)

		for {
			endTimestampMs := time.Now().UnixNano() / int64(time.Millisecond)
			lookbackMs := endTimestampMs - previousEndTimestampMs

			params := operations.NewGetTracesParams()
			params.ServiceName = &config.ServiceName
			params.AnnotationQuery = &config.Query
			params.Lookback = &lookbackMs
			params.EndTs = &endTimestampMs
			params.Limit = &config.QueryLimit

			traces, err := config.Client.Operations.GetTraces(params)
			if err != nil {
				log.Printf("[ERROR] could not fetch the traces: %s", err.Error())
			}

			<-time.After(config.Period)
			// TODO: add stop by command

			previousEndTimestampMs = endTimestampMs
		}
	}()
}
