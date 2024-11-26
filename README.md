

# Requirements

You can find installation instructions Go, protoc, and prococ-gen-go-grpc [here](https://grpc.io/docs/languages/go/quickstart/)

# Usage

```shell
go install .
protoc --proto_path . -I=. sample.proto --go-tqid_out=./out --go_out=./out
```