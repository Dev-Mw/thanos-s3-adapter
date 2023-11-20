export QUERY_STRING='avg(cluster:cpu_usage_nosum:rate5m{cluster_name="all-clusters"})by(node,pod,container,resource)'
export QUERY_STEP="60s"
export ENDPOINT_ADDRESS="http://localhost:9090"
export INTERVAL="3600" # Query 1 hour of the metric
export SCHEDULE="*/60 * * * *" # Every 60 minutes
export BACK_TIME="3" # 3 Hour Before
export METRIC_BUCKET="raw-data-metrics-test"
export METRIC_NAME="test"