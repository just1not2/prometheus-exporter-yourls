// Copyright: (c) 2022, Justin Béra (@just1not2) <me@just1not2.org>
// GNU General Public License v3.0+ (see LICENSE or https://www.gnu.org/licenses/gpl-3.0.txt)

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Throws an error if a configuration file is not provided
	if len(os.Args) > 2 {
		log.Fatalf("error: command format is %s <configuration file>", os.Args[0])
	}
	configuration := NewConfiguration()

	collector := NewYourlsCollector(configuration)
	prometheus.MustRegister(collector)

	// Starts the Prometheus exporter webserver
	http.Handle("/metrics", promhttp.Handler())
	fmt.Printf("Starting the YOURLS Prometheus exporter, listening on port %v...\n", configuration.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", configuration.Port), nil))
}
