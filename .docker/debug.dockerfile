
FROM golang:latest as builder

WORKDIR /app

RUN go get github.com/go-delve/delve/cmd/dlv

COPY ../go.mod .
COPY ../go.sum .

RUN go mod download
COPY ../. .

RUN CGO_ENABLED=0 GOOS=linux go build -o main/grpcapi main/main.go 
EXPOSE 40000 1541
WORKDIR /app/main
## Code built using this logic current working dir is main
CMD [ "dlv", "debug", "/app/main", "--listen=:40000", "--headless=true", "--api-version=2", "--log" ]