// Copyright: (c) 2022, Justin Béra (@just1not2) <me@just1not2.org>
// GNU General Public License v3.0+ (see LICENSE or https://www.gnu.org/licenses/gpl-3.0.txt)

package main

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type YourlsCollector struct {
	client                       *YourlsClient
	YourlsClicksTotalDescription *prometheus.Desc
	YourlsLinksTotalDescription  *prometheus.Desc
}

func NewYourlsCollector(configuration *YourlsConfiguration) *YourlsCollector {
	return &YourlsCollector{
		client: NewYourlsClient(configuration.YourlsURL, configuration.Signature, configuration.HTTPTimeout),
		YourlsClicksTotalDescription: prometheus.NewDesc("yourls_clicks_total",
			"Total number of clicks.",
			nil, nil,
		),
		YourlsLinksTotalDescription: prometheus.NewDesc("yourls_links_total",
			"Total number of deployed links.",
			nil, nil,
		),
	}
}

func (collector *YourlsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.YourlsClicksTotalDescription
	ch <- collector.YourlsLinksTotalDescription
}

func (collector *YourlsCollector) Collect(ch chan<- prometheus.Metric) {
	// Gathers the specific query parameters
	parameters := map[string]string{
		"action": "stats",
	}

	body, err := collector.client.Request(parameters)
	if err != nil {
		panic(err)
	}

	// Converts returned statistics into floats
	linksTotalValue, _ := strconv.ParseFloat(body.Stats["total_links"], 64)
	clicksTotalValue, _ := strconv.ParseFloat(body.Stats["total_clicks"], 64)

	// Sends metrics to the channel
	m1 := prometheus.MustNewConstMetric(collector.YourlsClicksTotalDescription, prometheus.GaugeValue, clicksTotalValue)
	m2 := prometheus.MustNewConstMetric(collector.YourlsLinksTotalDescription, prometheus.GaugeValue, linksTotalValue)
	m1 = prometheus.NewMetricWithTimestamp(time.Now(), m1)
	m2 = prometheus.NewMetricWithTimestamp(time.Now(), m2)
	ch <- m1
	ch <- m2
}
