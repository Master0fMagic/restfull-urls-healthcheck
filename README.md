## URLs Health Check Service

### This service prove a http api to ping specified URLs and return their statuses

### API:
 
- **POST** `/urls-health` 
- Request Body: JSON `["url1", "url2"]`
- Response: JSON `{"url1":"active", "url2":"inactive"}`
- Possible errors:
  - Code: `400 Bad Request`, Response Body: JSON `{"error": "error message"}`. For cases:
    - invalid URL format
    - empty URLs list
    - request body is not a list
  - Code: `500 Internal Server Error`, Response Body: JSON `{"error": "error message"}`
- If response from API takes more than `$HEALTHCHECK_TIMEOUT` - API marked as inactive
- If `$HEALTHCHECK_STOP_ON_FAILURE` is true, then the first inactive API stops health check process for all APIs left

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

