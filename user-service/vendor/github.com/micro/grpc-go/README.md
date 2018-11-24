# gRPC-Go

gRPC-go is a gRPC transport library

It is a fork and stripped down version of the official [grpc-go](https://github.com/grpc/grpc-go) library.

## Why?

We continually encountered breaking changes, the library didn't feel like it was made to be used by developers. 
The transport was moved to an internal package which broke support for go-micro so we've forked to alleviate these problems.

This library will be solely focused on the gRPC transport. We have no desires to maintain a framework. That's what micro is for.

## Usage

See official [readme](https://github.com/grpc/grpc-go) for usage.

## Compatibility

- gRPC-go continues to be compatible with official grpc applications
- gRPC-go is not a replacement for [go-grpc](https://github.com/micro/go-grpc)

## License

- gRPC-go remains under the same license as the official library
