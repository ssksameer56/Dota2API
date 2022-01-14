FROM golang:latest as builder

WORKDIR /app

COPY ../. .

RUN CGO_ENABLED=0 GOOS=linux go build /clients/grpc-client/main/main.go -a -o grpcclient

##Final Build
FROM alpine:latest

#Copy the Required contents
COPY --from=builder /clients/grpc-client/main/grpcclient ./app/clients/main
COPY --from=builder /app/config.json .
COPY --from=builder /app/cert.pem .
COPY --from=builder /app/key.pem .

ENTRYPOINT ["./app/clients/main/grpcclient"]

CMD ["0"]