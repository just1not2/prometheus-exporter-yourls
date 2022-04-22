# YOURLS Prometheus Exporter [![Listed in Awesome YOURLS!](https://img.shields.io/badge/Awesome-YOURLS-C5A3BE)](https://github.com/YOURLS/awesome-yourls/)

A Prometheus exporter to fetch some metrics from a YOURLS URL-shortener instance.

## Exported metrics

### Gauges
Name | Description
--- | ---
`yourls_clicks_total`|Total number of clicks
`yourls_links_total`|Total number of deployed links

### Go metrics
A variety of Go metrics that describe the exporter are exported as well.


## Installation

### Docker image
The preferred way to use the exporter is by running the provided Docker image. It is currently available on Docker Hub:

```bash
docker pull just1not2/prometheus-exporter-yourls:latest
```

In addition to the `latest` tag which points to the version currently in the `main` branch, tagged versions are also available.

### From source
You can clone the repository and build the exporter locally.

```bash
git clone https://github.com/just1not2/prometheus-exporter-yourls.git
cd prometheus-exporter-yourls
make local
```


## Configuration

### Exporter configuration
There are two ways to configure the exporter:
* Environment variables
* Configuration file

The first method takes precedence over the second one.

To use the first one, you just have to declare environment variables while launching the Docker image, a Docker-compose or directly the executable:
```bash
docker run -p 9923:9923 \
           -e YOURLS_URL=http://yourls.example.com \
           -e YOURLS_SIGNATURE=SECRET_API_KEY \
           just1not2/prometheus-exporter-yourls:latest
```

If you prefer to configure your provider with the second method, you should copy the [config.json.template](./config.json.template) file and replace the sample values. You can then launch the exporter:
```bash
docker run -p 9923:9923 -v $(pwd)/config.json:/config.json just1not2/prometheus-exporter-yourls:latest config.json
```

The configuration parameters that can be declared are summarized in this table:
Parameter | Environment variable | Description | Note
--- | --- | --- | ---
`url`|`YOURLS_URL`|URL of the YOURLS instance|_Required_
`signature`|`YOURLS_SIGNATURE`|Signature of the YOURLS instance|_Required_
`exporter_port`|`YOURLS_EXPORTER_PORT`|Port on which the exporter listens|_Default is 9923_
`exporter_timeout`|`YOURLS_EXPORTER_TIMEOUT`|Timeout of requests to the YOURLS instance|_Default is 10_

The `signature` parameter can be found at `http://<your YOURLS instance URL>/admin/tools.php`.

### Scrape configuration
The exporter will query the YOURLS server every time it is scraped by Prometheus. You can modify the configuration of the exporter to change the scrape interval:

```yml
scrape_configs:
  - job_name: yourls
    scrape_interval: 1m
    static_configs:
      - targets: ['<your YOURLS exporter host>:9923']
```


## See also

* [YOURLS official documentation](https://yourls.org)
* [Use Prometheus exporter](https://prometheus.io/docs/instrumenting/exporters/)


## Contributing to this exporter

This exporter started as personal project, but I welcome community contributions to this exporter. If you find problems, please open an issue or create a PR against the [YOURLS exporter repository](https://github.com/just1not2/prometheus-exporter-yourls).

You can also reach me by email at `me@just1not2.org`.


## Licensing

GNU General Public License v3.0 or later.

See [LICENSE](./LICENSE) to see the full text.


## Author information

This exporter was created in 2022 by Justin Béra.
