FROM golang:1.21-alpine

RUN apk add --no-cache git make

WORKDIR /go/src/app

COPY . .

RUN make build

WORKDIR /bin
EXPOSE 8080

# Run the app when the container starts
CMD ["./urls-health-check"]