GOLINT := golangci-lint

all: dep lint build

dep:
	go mod download

check-lint:
	@which $(GOLINT) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GO_PATH)/bin v1.54.2

lint: dep check-lint
	 $(GOLINT) run --timeout 10m -c .golangci.yml

build: dep
	go build -v -o /bin/urls-health-check ./
	@echo "Done building."

check-mockgen:
	@which mockgen || go install github.com/golang/mock/mockgen@v1.6.0

mockgen-generate: dep check-mockgen
	mockgen -package health -self_package github.com/Master0fMagic/rest-health-check-service/health -source health/health.go > health/health_mock.go

test: dep
	 go test -tags unit -cover -count=1 ./health/... ./server/...

build-docker:
	docker build -t wotb-auction-bot .