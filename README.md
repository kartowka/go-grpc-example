## Prerequisites

- Go 1.22 or later
- `protoc` (Protocol Buffers compiler)
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins

## Setup

1. Clone the repository:

```sh
git clone https://github.com/antfley/go-grpc-example.git
cd go-grpc-example
```

2. Install dependencies:

```sh
go mod tidy
```

3. Generate gRPC code from the .proto file:

```sh
protoc --go_out=. --go-grpc_out=. chat/chat.proto
```

## Running the Server

To run the gRPC server:

```sh
go run cmd/api/main.go
```

The server will start and listen on the port specified in the .env file (default is 3000).

## Running the Client

To run the gRPC client:

```sh
go run cmd/client/main.go
```

The client will send a message to the server and print the response.

### License

This project is licensed under the Apache License 2.0 - see the LICENSE file for details.
