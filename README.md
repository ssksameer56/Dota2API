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
-  