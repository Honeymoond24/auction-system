# Auction-system
Prototype of the auction system

## How to run
### Requirements:
- Docker engine, docker-compose

### Set environment variables
Copy .env.example to .env and adjust the values
```shell
cp .env.example .env
# edit .env
```

### Run with docker-compose
```shell
docker-compose up -d --build
```

### Dev commands
Install golang-migrate:
```shell
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```
Create migration:
```shell
migrate create -ext sql -dir internal/infrastructure/database/migrations -seq create_tablename_table
```
Run migrations:
```shell
migrate -database ${POSTGRESQL_URL} -path internal/infrastructure/database/migrations up
```
Generate gRPC code:
```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/interfaces/grpc_server/pb/auction.proto
```
Generate gRPC Gateway code:
```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative --grpc-gateway_opt=grpc_api_configuration=./internal/interfaces/grpc_server/pb/auction_service.yaml ./internal/interfaces/grpc_server/pb/auction.proto
```
Run server:
```shell
go run cmd/server/main.go
```