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
To run this app pull repo and execute `./run.sh`

