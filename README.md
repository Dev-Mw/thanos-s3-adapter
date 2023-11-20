# Thanos S3 Adapter

![technology Go](https://img.shields.io/badge/technology-go-blue.svg)

The metric extraction adapter is a service configured to transport large volumnes of data, apply transforms to the metrics consumed from the [Thanos Query](https://thanos.io/tip/components/query.md/) endpoint and then take them to the [AWS S3](https://aws.amazon.com/es/s3/) storage service using the [line-delimited JSON format](https://jsonlines.org/) compressed as GZ.

## How to use

To execute the adapter run

```bash
# Export ENV
go run cmd/api/main.go
```

For build the adapter executable you can use

```bash
swag init --dir cmd/api --output cmd/api/docs && go build -o adapter ./cmd/api
```

### Environment variables

* `QUERY_STRING`: is the variable that contains the query that is needed to ask the service of thanos for the data, this uses the same syntax of [Prometheus](https://prometheus.io/docs/prometheus/latest/querying/basics/).
* `QUERY_STEP`: It is the value that is used to define what granularity the metric packages consumed from the thanos query service will have.
* `ENDPOINT_ADDRESS`: Endpoint or url that the adapter would consume (Thanos querier).
* `INTERVAL`: User-defined time (in seconds) that the adapter will use to determine the start and end of the interval when consuming the thanos service.
* `SCHEDULE`: User-defined time (in valid [Cron](https://crontab.guru/) format) that the adapter will use to consume the Thanos service.
  > Make sure that the relationship between `METRIC_SCRAPER_TIME` and `SCHEDULE` does not generate duplicates when stepping time intervals.
* `METRIC_BUCKET`: Name of the bucket in which the files in JSON format will be uploaded.
* `METRIC_NAME`: Name of the metric that will be used as the base value to define the name of folders in S3 service.
* `BACK_TIME`: Time back in hours in which the metric will be consumed (check Thanos's configuration to understand how often he reconciles the information of all the connected clusters).

An example of config

```bash
#!/bin/sh

export QUERY_STRING='avg(cluster:cpu_usage_nosum:rate5m{cluster_name="CLUSTER_NAME"})by(node,pod,container,resource)'
export QUERY_STEP="60s"
export ENDPOINT_ADDRESS="http://localhost:9090"
export INTERVAL="3600" # Query 1 hour of the metric
export SCHEDULE="*/60 * * * *" # Every 60 minutes
export BACK_TIME="3" # 3 Hour Before
export METRIC_BUCKET="raw-data-metrics-test"
export METRIC_NAME="test"
```

This configuration will consult the information given by the query in an interval of 1 hour(`3600` seconds) with data every minute(`"60s"`) and repeat the process every hour(`*/60 * * * *`).

Queries to the Thanos Query API will look like this:

* `http://localhost:9090/api/v1/query_range?query=avg(cluster:cpu_usage_nosum:rate5m{cluster_name="cluster-one"})by(node,pod,container,resource)&step=60s&dedup=true&start=2023-01-01T15:00:00.000Z&end=2023-01-01T15:59:59.000Z`
* `http://localhost:9090/api/v1/query_range?query=avg(cluster:cpu_usage_nosum:rate5m{cluster_name="cluster-two"})by(node,pod,container,resource)&step=60s&dedup=true&start=2023-01-01T15:00:00.000Z&end=2023-01-01T15:59:59.000Z`
* ... So on changing the `cluster_name`

> **Important:** `CLUSTER_NAME` is a reserved word that can be used in the query and will be replaced internally by each of the clusters that are connected to the Thanos service.

### Storage

Inside the S3 bucket you will find a structure with `[Metric_Name]/[Year]/[Month]/[Day]/[Hour]/[Cluster_Name].json.gz`.

```console
Metric_Name
└── 2023
    └── 04
        └── 01
            ├── 01
            │   └── cluster_name.json.gz
            ├── 02
            │   └── cluster_name.json.gz
            ├── ...
            │   └── cluster_name.json.gz
            └── 23
                └── cluster_name.json.gz
```

> You can find all available image tags for your Dockerfile
> [here](https://github.com/mercadolibre/fury_go-mini#supported-tags).

## How to run using Docker?

Set the `Dockerfile`

```docker
FROM golang:alpine as build

ENV GOPROXY=https://proxy.golang.org

WORKDIR /go/src/adapter
COPY . .

RUN GOOS=linux go build -o /go/bin/adapter ./cmd/api/main.go

FROM alpine

COPY --from=build /go/bin/adapter /go/bin/adapter

RUN mkdir /root/.aws

COPY ~/.aws/credentials /root/.aws/credentials
COPY ~/.aws/config /root/.aws/config

ENTRYPOINT [ "/go/bin/adapter" ]
```

To build `docker build --tag adapter-service:latest .` and then run using

```bash
$ docker run --network=host \
    -e ENDPOINT_ADDRESS=$ENDPOINT_ADDRESS \
    -e QUERY_STRING=$QUERY_STRING \
    -e QUERY_STEP=$QUERY_STEP \
    -e INTERVAL=$INTERVAL \
    -e METRIC_BUCKET=$METRIC_BUCKET \
    -e METRIC_NAME=$METRIC_NAME \
    -e BACK_TIME=$BACK_TIME \
    --name adapter-service-container adapter-service:latest;
```

> **NOTE:** If you are running in local environment, you should to have the AWS credentials file in folder with the path name `~/.aws/credentials` and config file `~/.aws/config`. See more info [AWS config](https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/guide_credentials_profiles.html).

## Test

Run `go mod tidy` to ensure that the `go.mod` file matches the source code in the module.

To generate coverage report you can use:

```bash
go test -v ./... -covermode=atomic -coverprofile=coverage.out -coverpkg=./... -count=1  -race -timeout=30m
```
