package plugin

import (
	"fmt"
	"path"
	"strings"

	"github.com/najeal/rpc-fusion/internal/templater"
	"google.golang.org/protobuf/compiler/protogen"
)

const packageNameSuffix = "fusion"

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
	g := gen.NewGeneratedFile(file.GeneratedFilenamePrefix+"/fusion"+file.GeneratedFilenamePrefix+".fusion.go", file.GoImportPath)
	g.P(string(content))
	return nil
}

func generateTemplateDatas(files []*protogen.File) []templater.File {
	fileDatas := []templater.File{}
	for _, f := range files {
		if !f.Generate {
			continue
		}
		fileDatas = append(fileDatas, generateFileData(nil, f))
	}
	return fileDatas
}

func generateFileData(gen *protogen.Plugin, file *protogen.File) templater.File {
	fileData := templater.File{
		PackageName:    string(file.GoPackageName) + packageNameSuffix,
		PackageImports: getRequiredGoImports(file),
		Services:       getServicesData(file),
	}
	return fileData
}

func getServicesData(file *protogen.File) []templater.Service {
	svcs := []templater.Service{}
	for _, isvc := range file.Services {
		svc := templater.Service{}
		for _, method := range isvc.Methods {
			inputPackage := path.Base(strings.TrimSuffix(method.Input.GoIdent.GoImportPath.String(), "\""))
			outputPackage := path.Base(strings.TrimSuffix(method.Output.GoIdent.GoImportPath.String(), "\""))
			inputType := fmt.Sprintf("%s.%s", inputPackage, method.Input.GoIdent.GoName)
			outputType := fmt.Sprintf("%s.%s", outputPackage, method.Output.GoIdent.GoName)
			svc.CommonMethods = append(svc.CommonMethods,
				fmt.Sprintf("%s(ctx context.Context, arg *connect.Request[%s], res *connect.Response[%s]) error",
					method.GoName,
					inputType,
					outputType,
				))
			svc.GrpcMethods = append(svc.GrpcMethods,
				fmt.Sprintf("%s(ctx context.Context, arg *connect.Request[%s]) (res *connect.Response[%s], err error)",
					method.GoName,
					inputType,
					outputType,
				))
			svc.JsonrpcMethods = append(svc.JsonrpcMethods,
				fmt.Sprintf("%s(req *http.Request, arg *connect.Request[%s], res *connect.Response[%s]) error",
					method.GoName,
					inputType,
					outputType,
				))
			svc.MethodNames = append(svc.MethodNames, method.GoName)
			svc.ResponseTypes = append(svc.ResponseTypes, outputType)
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
