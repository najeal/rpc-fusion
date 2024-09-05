package plugin

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/najeal/rpc-fusion/internal/templater"
	"google.golang.org/protobuf/compiler/protogen"
)

const (
	packageNameSuffix = "fusion"
	fileExtension     = ".fusion.go"
)

func Run(gen *protogen.Plugin) error {
	for _, f := range gen.Files {
		if !f.Generate {
			continue
		}
		fileData := generateFileData(nil, f)
		content, err := templater.GenerateFile(fileData)
		if err != nil {
			return err
		}
		writeFile(gen, f, content)
	}
	return nil
}

func writeFile(gen *protogen.Plugin, file *protogen.File, content []byte) error {
	fileName := strings.Split(file.GeneratedFilenamePrefix, string(filepath.Separator))[0]
	g := gen.NewGeneratedFile(filepath.Join(file.GeneratedFilenamePrefix+packageNameSuffix, fileName+fileExtension), file.GoImportPath)
	g.P(string(content))
	return nil
}

func generateFileData(_ *protogen.Plugin, file *protogen.File) templater.File {
	fileData := templater.File{
		PackageName:    string(file.GoPackageName) + packageNameSuffix,
		PackageImports: getRequiredGoImports(file),
		Services:       getServicesData(file),
	}
	return fileData
}

var methodsFormat map[string]map[string]string = map[string]map[string]string{
	"basic": {
		"common":        "%s(ctx context.Context, arg *%s, res *%s) error",
		"grpc":          "%s(ctx context.Context, arg *%s) (res *%s, err error)",
		"jsonrpc":       "%s(req *http.Request, arg *%s, res *%s) error",
		"responsetypes": "%s",
	},
	"withconnect": {
		"common":        "%s(ctx context.Context, arg *connect.Request[%s], res *connect.Response[%s]) error",
		"grpc":          "%s(ctx context.Context, arg *connect.Request[%s]) (res *connect.Response[%s], err error)",
		"jsonrpc":       "%s(req *http.Request, arg *connect.Request[%s], res *connect.Response[%s]) error",
		"responsetypes": "connect.Response[%s]",
	},
}

func getServicesData(file *protogen.File) []templater.Service {
	svcs := []templater.Service{}
	servicePackage := path.Base(strings.TrimSuffix(file.GoImportPath.String(), "\""))
	for _, isvc := range file.Services {
		svc := templater.Service{
			ServiceName:    isvc.GoName,
			ServicePackage: servicePackage,
		}
		for _, method := range isvc.Methods {
			servicePackage := path.Base(strings.TrimSuffix(file.GoImportPath.String(), "\""))
			inputPackage := path.Base(strings.TrimSuffix(method.Input.GoIdent.GoImportPath.String(), "\""))
			outputPackage := path.Base(strings.TrimSuffix(method.Output.GoIdent.GoImportPath.String(), "\""))
			inputType := fmt.Sprintf("%s.%s", inputPackage, method.Input.GoIdent.GoName)
			outputType := fmt.Sprintf("%s.%s", outputPackage, method.Output.GoIdent.GoName)
			svc.CommonMethods = append(svc.CommonMethods,
				fmt.Sprintf(methodsFormat["basic"]["common"],
					method.GoName,
					inputType,
					outputType,
				))
			svc.GrpcMethods = append(svc.GrpcMethods,
				fmt.Sprintf(methodsFormat["basic"]["grpc"],
					method.GoName,
					inputType,
					outputType,
				))
			svc.JsonrpcMethods = append(svc.JsonrpcMethods,
				fmt.Sprintf(methodsFormat["basic"]["jsonrpc"],
					method.GoName,
					inputType,
					outputType,
				))
			svc.MethodNames = append(svc.MethodNames, method.GoName)
			svc.ResponseTypes = append(svc.ResponseTypes, fmt.Sprintf(methodsFormat["basic"]["responsetypes"], outputType))
			svc.MustEmbedUnimplemented = servicePackage + ".Unimplemented" + isvc.GoName + "Server"
		}
		svcs = append(svcs, svc)
	}
	return svcs
}

func getRequiredGoImports(file *protogen.File) map[string]struct{} {
	goimports := map[string]struct{}{}
	for _, svc := range file.Services {
		for _, method := range svc.Methods {
			inPackage := path.Base(strings.TrimSuffix(method.Input.GoIdent.GoImportPath.String(), "\""))
			outPackage := path.Base(strings.TrimSuffix(method.Output.GoIdent.GoImportPath.String(), "\""))
			goimports[fmt.Sprintf("%s %s", inPackage, method.Input.GoIdent.GoImportPath.String())] = struct{}{}
			goimports[fmt.Sprintf("%s %s", outPackage, method.Output.GoIdent.GoImportPath.String())] = struct{}{}
		}
	}
	return goimports
}
