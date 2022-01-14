FROM golang:latest as builder

WORKDIR /app
COPY ../go.mod .
COPY ../go.sum .

RUN go mod download

COPY ../ .

RUN CGO_ENABLED=0 GOOS=linux go build -o main/graphapi /main/main.go 

##Final Build
FROM alpine:latest

COPY --from=builder /app/main/graphapi ./main
COPY --from=builder /app/config.json ./app
COPY --from=builder /app/cert.pem ./app
COPY --from=builder /app/key.pem ./app


ENTRYPOINT ["/main/graphapi"]

CMD ["-graph=true","-grpc=false"]