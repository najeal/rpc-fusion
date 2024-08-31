package main

import (
	"github.com/najeal/rpc-fusion/internal/plugin"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			if err := plugin.Run(gen); err != nil {
				return err
			}
		}
		return nil
	})
}
