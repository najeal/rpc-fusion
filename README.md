# Welcome to rpc-fusion

rpc-fusion [buf tool](https://github.com/bufbuild/buf) plugin.
it generates JSONRPC-Server and  GRPC-Server using a common interface.

The user is free to implement the common interface and use it in both grpc and jsonrpc servers.

# Install rpc-fusion:
```go install github.com/najeal/rpc-fusion/cmd/rpc-fusion@latest```

# Use rpc-fusion:

Put the plugin in your `buf.gen.yaml`:
```
version: v2
...
plugins:
  - local: rpc-fusion
    opt: paths=source_relative
    out: gen
.....
```

Take a look on the example in `tests` directory:

- `tests/gen/coreapi/v1/coreapifusion/`: contains the grpc & jsonrpc servers using the common interface + the jsonrpc client
- `tests/cmd/common.go`: contains the struct implementing the common interface.
- `tests/cmd/rpc-server`: use the common interface to run a jsonrpc-server.
- `tests/cmd/rpc-client`: use the generated client to request the jsonrpc-server.
- `tests/cmd/grpc-server`: use the common interface to run a grpc-server.
- `tests/cmd/grpc-client`: use the generated client to request the grpc-server.
