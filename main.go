package main

import (
	"fmt"
	"github.com/TheQueenIsDead/protoc-gen-go-tqid/pkg"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
	"io"
	"os"
)

func main() {
	// protoc passes pluginpb.CodeGeneratorRequest in via stdin marshalled with Protobuf
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	var req pluginpb.CodeGeneratorRequest
	err = proto.Unmarshal(input, &req)
	if err != nil {
		panic(err)
	}

	// Initialise our plugin with default options
	opts := protogen.Options{}
	plugin, err := opts.New(&req)
	if err != nil {
		panic(err)
	}

	// Generate boilerplate, and proto specific code
	err = pkg.Generate(plugin)
	if err != nil {
		panic(err)
	}

	// Marshal a plugin response as protobuf
	response := plugin.Response()
	output, err := proto.Marshal(response)
	if err != nil {
		panic(err)
	}

	// Write the response to os.Stdout, to be picked up by protoc
	n, err := fmt.Fprintf(os.Stdout, string(output))
	if err != nil {
		panic(err)
	} else if n != len(output) {
		panic(pkg.ErrBadWrite)
	}
}
