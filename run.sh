export LOG_LEVEL="INFO"
export HTTP_PORT="8080"
export HEALTHCHECK_TIMEOUT="1s"
export HEALTHCHECK_STOP_ON_FAILURE="false"
export HEALTHCHECK_MAX_PROCESSING_GOROUTINE="-1"

# Run the Go application
go mod download
go run main.go