
FROM golang:latest as builder

WORKDIR /app

COPY ../go.mod .
COPY ../go.sum .

RUN go mod download

COPY ../. .
RUN CGO_ENABLED=0 GOOS=linux go build -o main/grpcapi main/main.go 

#Final Build
FROM alpine:latest
WORKDIR /app
RUN apk add --update --no-cache bash ca-certificates git
RUN mkdir main
COPY --from=builder /app/main/grpcapi ./main
COPY --from=builder /app/config.json .
COPY --from=builder /app/cert.pem .
COPY --from=builder /app/key.pem .

ENTRYPOINT ["main/grpcapi"]

CMD ["-graph=false","-grpc=true"]