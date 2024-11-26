

# Requirements

You can find installation instructions Go, protoc, and prococ-gen-go-grpc [here](https://grpc.io/docs/languages/go/quickstart/)

# Usage

## Buf

Buf is the preferred way to generate code with protoc.

```shell
cd example
buf generate
```

## Protoc

If you prefer to do things manually, the same effect can be achieved with the following

```shell
go install .
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@latest
cd example
protoc \
  --go_out=. \
  --go-tqid_out=. \
  --go-vtproto_out=. \
  --go-vtproto_opt=features=marshal+unmarshal+size \
  proto/sample.proto  
```

### Further Reading

 - https://astextract.lu4p.xyz/ is very helpful for converting code into native struct representation