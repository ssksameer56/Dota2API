
FROM golang:latest as builder

WORKDIR /app

COPY ../go.mod .
COPY ../go.sum .

RUN go mod download

COPY ../. .
RUN CGO_ENABLED=0 GOOS=linux go build -o main/grpcapi main/main.go 
ENTRYPOINT ["/bin/bash"]