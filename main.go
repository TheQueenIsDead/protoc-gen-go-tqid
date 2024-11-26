package main

import (
	"flag"
	"github.com/TheQueenIsDead/protoc-gen-go-tqid/pkg"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	var flags flag.FlagSet
	serviceName := flags.String("svc", "", "The name of the service to generate. Eg, 'posts'.")
	opts := &protogen.Options{
		ParamFunc: flags.Set,
	}
	opts.Run(func(p *protogen.Plugin) error {
		if *serviceName == "" {
			return pkg.ErrServiceNameFlagRequired
		}
		return pkg.Generate(p, pkg.Options{ServiceName: *serviceName})
	})
}
