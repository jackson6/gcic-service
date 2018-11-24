# Micro gRPC [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GoDoc](https://godoc.org/github.com/micro/go-grpc?status.svg)](https://godoc.org/github.com/micro/go-grpc) [![Travis CI](https://api.travis-ci.org/micro/go-grpc.svg?branch=master)](https://travis-ci.org/micro/go-grpc) [![Go Report Card](https://goreportcard.com/badge/micro/go-grpc)](https://goreportcard.com/report/github.com/micro/go-grpc)

A micro gRPC framework. A simplified experience for building gRPC services.

## Overview

Go-grpc makes use of [go-micro](https://github.com/micro/go-micro) to create a simpler framework for gRPC development. 
We use [github.com/grpc/grpc-go](https://github.com/grpc/grpc-go) beneath the covers but provide a 
[micro.Service](https://godoc.org/github.com/micro/go-micro#Service) for the end user.

## Getting Started

See the [docs](https://micro.mu/docs/go-grpc.html) to get started.

## Examples

Find an example greeter service in [examples/greeter](https://github.com/micro/go-grpc/tree/master/examples/greeter).
