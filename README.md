## URLs Health Check Service

### This service prove a http api to ping specified URLs and return their statuses

### Config params (via ENVs):
```
LOG_LEVEL=INFO
HTTP_PORT=8080
HEALTHCHECK_TIMEOUT=1s
HEALTHCHECK_STOP_ON_FAILURE=false
HEALTHCHECK_MAX_PROCESSING_GOROUTINE=-1
```

### Run
- run sh script:
```
./run.sh
```
- run with docker:
```
docker build -t urls-healch-check && docker run -it \
-e LOG_LEVEL=INFO \
-e HTTP_PORT=8080 \
-e HEALTHCHECK_TIMEOUT=1s \
-e HEALTHCHECK_STOP_ON_FAILURE=false \
-e HEALTHCHECK_MAX_PROCESSING_GOROUTINE=-1 \
urls-health-check
```

