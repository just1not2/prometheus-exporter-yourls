# Copyright: (c) 2022, Justin Béra (@just1not2) <me@just1not2.org>
# GNU General Public License v3.0+ (see LICENSE or https://www.gnu.org/licenses/gpl-3.0.txt)

VERSION=1.0.0

default: install

build:
	docker build -t just1not2/prometheus-exporter-yourls:test .

install: build
	docker run -p 9923:9923 -v ${shell pwd}/config.json:/config.json just1not2/prometheus-exporter-yourls:test config.json

local:
	go build -o ./bin/prometheus-exporter-yourls
	./bin/prometheus-exporter-yourls config.json

release: build
	docker tag just1not2/prometheus-exporter-yourls:test just1not2/prometheus-exporter-yourls:latest
	docker tag just1not2/prometheus-exporter-yourls:test just1not2/prometheus-exporter-yourls:${shell echo ${VERSION} | cut -d '.' -f -2}
	docker tag just1not2/prometheus-exporter-yourls:test just1not2/prometheus-exporter-yourls:${VERSION}
	docker push just1not2/prometheus-exporter-yourls:latest
	docker push just1not2/prometheus-exporter-yourls:${shell echo ${VERSION} | cut -d '.' -f -2}
	docker push just1not2/prometheus-exporter-yourls:${VERSION}
