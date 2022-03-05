# Proto

This directory contains protocol buffers file for each services.

## Generating Go client & server stub

This guide assumes you already have `protoc` compiler installed.
See [here](https://grpc.io/docs/protoc-installation/) if you haven't.

Install Go compiler plugins:

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

Then, generate the file.

```sh
protoc --go_out=../<service directory> --go-grpc_out=../<service directory> <proto file>

# Samples:
protoc --go_out=../worker --go-grpc_out=../worker worker.proto
```