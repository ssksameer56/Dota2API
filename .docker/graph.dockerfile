FROM golang:latest as builder

WORKDIR /app
COPY ../go.mod .
COPY ../go.sum .

RUN go mod download

COPY ../ .

RUN CGO_ENABLED=0 GOOS=linux go build -o main/graphapi /main/main.go 

##Final Build
FROM alpine:latest
WORKDIR /app
RUN apk add --update --no-cache bash ca-certificates git
RUN mkdir main
COPY --from=builder /app/main/graphapi ./main
COPY --from=builder /app/config.json ./app
COPY --from=builder /app/cert.pem ./app
COPY --from=builder /app/key.pem ./app
WORKDIR /app/main
ENTRYPOINT ["./graphapi"]
EXPOSE 8080
ENV DATABASE_HOST=docker-database
CMD ["-graph=true"]