# Dota2API

Dota 2 API built using GraphQL, gRPC and HTTP2

## Before You Start Running

- Create a `logs` folder where you run the API to store the logs
- Generate a cert and key file(mkcert recommended) to use TLS
- Obtain an API Key from OpenDota if you want higher rate limit for using Dota2 API. Add to `config.json`
- Run the `scripts/favourites.sql` file to generate the required database and tables
- Add your database credentials to the `config.json`

## Running API
- The `main` package runs both gRPC and graphQL API simaltaneously when run. You can add required flags to start just one of the API.
  - `-grpc` and `-graph` are the command line flags
- Provide the `certificate` and `key` location in the config file
- Build using `go build main/main.go -o {{filename}}`

## Using the API
- GraphQL exposes a playground that you can directly use
- For testing gRPC, you can use [BloomRPC](https://github.com/bloomrpc/bloomrpc)

## Clients
- You can use the gRPC Command Line Client by going to `client/grpc-client/main`. Build normally.
- To use specific service, you can pass the number as command line argument followed by any parameter as required.
- Logs are stored at `log/clientlog.log`
- Need to pass the location of `cert.pem` file of server to client to use TLS.

## Built Using and with Help of 
- Go
- [GQLGen](https://gqlgen.com/)
- [MkCert](https://github.com/FiloSottile/mkcert)
- [BloomRPC](https://github.com/bloomrpc/bloomrpc)
- [Logrus](https://github.com/sirupsen/logrus)

## TO-DO
- Add extensive Logging
- Add GUI Client using Svelte
- Dockerize the services using Docker Compose (MySQL,gRPC, Graph, Svelte)