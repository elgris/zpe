package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/elgris/zpe/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	yaml "gopkg.in/yaml.v2"
)

// AppConfig specifies all the necessary properties for Zipkin-Prometheus extractor
// to run
type AppConfig struct {
	Listen    string        `yaml:"listen"`
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

	config, err := loadConfig(*configFilePath)
	if err != nil {
		log.Fatal(err.Error())
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
		RunScraper(scraperConfig)
	}

	http.Handle("/metrics", promhttp.Handler())
	log.Printf("Starting server with %d zipkin scrapers on %s", len(config.Queries), config.Listen)

	log.Fatal(http.ListenAndServe(config.Listen, nil))
}

func loadConfig(configFilePath string) (*AppConfig, error) {
	configFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not read configuration file '%s': %s", configFilePath, err.Error())
	}

	config := &AppConfig{}
	err = yaml.Unmarshal(configFile, config)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal configuration from file '%s': %s", configFilePath, err.Error())
	}

	return config, nil
}
