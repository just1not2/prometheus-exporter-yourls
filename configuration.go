// Copyright: (c) 2022, Justin Béra (@just1not2) <me@just1not2.org>
// GNU General Public License v3.0+ (see LICENSE or https://www.gnu.org/licenses/gpl-3.0.txt)

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	DEFAULT_HTTP_TIMEOUT = 10
	DEFAULT_PORT         = 9923
)

type ExporterConfiguration struct {
	HTTPTimeout float64 `json:"exporter_timeout"`
	Port        float64 `json:"exporter_port"`
	Signature   string  `json:"signature"`
	YourlsURL   string  `json:"url"`
}

type YourlsConfiguration struct {
	HTTPTimeout time.Duration
	Port        float64
	Signature   string
	YourlsURL   string
}

func NewConfiguration() *YourlsConfiguration {
	// Initializes exporter configuration
	configuration := &ExporterConfiguration{
		HTTPTimeout: DEFAULT_HTTP_TIMEOUT,
		Port:        DEFAULT_PORT,
	}

	// Uses configuration file if it is passed as a parameter
	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Printf("warn: %s, cannot read %s (skipping)\n", err, os.Args[1])
		} else {
			if err := json.NewDecoder(file).Decode(configuration); err != nil {
				fmt.Printf("warn: %s, cannot read %s (skipping)\n", err, os.Args[1])
			}
		}
		file.Close()
	}

	// Overwrites configuration with existing environment variables
	if keyString, defined := os.LookupEnv("YOURLS_URL"); defined {
		configuration.YourlsURL = keyString
	}
	if keyString, defined := os.LookupEnv("YOURLS_SIGNATURE"); defined {
		configuration.Signature = keyString
	}
	if keyString, defined := os.LookupEnv("YOURLS_EXPORTER_PORT"); defined {
		if keyFloat, err := strconv.ParseFloat(keyString, 64); err != nil {
			fmt.Printf("warn: %s, cannot use YOURLS_EXPORTER_PORT environment variable (skipping)\n", err)
		} else {
			configuration.Port = keyFloat
		}
	}
	if keyString, defined := os.LookupEnv("YOURLS_EXPORTER_TIMEOUT"); defined {
		if keyFloat, err := strconv.ParseFloat(keyString, 64); err != nil {
			fmt.Printf("warn: %s, cannot use YOURLS_EXPORTER_TIMEOUT environment variable (skipping)\n", err)
		} else {
			configuration.HTTPTimeout = keyFloat
		}
	}

	// Throws error for missing fields
	if configuration.YourlsURL == "" {
		log.Fatal("error: YOURLS_URL environment variable and 'url' configuration parameter were both undefined")
	} else if configuration.Signature == "" {
		log.Fatal("error: YOURLS_SIGNATURE environment variable and 'signature' configuration parameter were both undefined")
	}

	return &YourlsConfiguration{
		HTTPTimeout: time.Duration(configuration.HTTPTimeout * float64(time.Second)),
		YourlsURL:   fmt.Sprintf("%s/yourls-api.php?format=json&signature=%s", configuration.YourlsURL, configuration.Signature),
		Port:        configuration.Port,
		Signature:   configuration.Signature,
	}
}
