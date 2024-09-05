package templater

type Service struct {
	ServicePackage         string
	ServiceName            string
	CommonMethods          []string
	JsonrpcMethods         []string
	GrpcMethods            []string
	ResponseTypes          []string
	MethodNames            []string
	MustEmbedUnimplemented string
}

type File struct {
	PackageName    string
	PackageImports map[string]struct{}
	Services       []Service
}
