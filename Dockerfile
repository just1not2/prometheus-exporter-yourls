# Copyright: (c) 2022, Justin Béra (@just1not2) <me@just1not2.org>
# GNU General Public License v3.0+ (see LICENSE or https://www.gnu.org/licenses/gpl-3.0.txt)

# Build
FROM golang:1.17 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 go build -o /prometheus-exporter-yourls


# Install
FROM scratch
LABEL maintainer="me@just1not2.org"

WORKDIR /

COPY --from=build /prometheus-exporter-yourls /prometheus-exporter-yourls
COPY config.json.template config.json

EXPOSE 9923

ENTRYPOINT ["/prometheus-exporter-yourls"]
