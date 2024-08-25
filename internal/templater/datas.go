package templater

type Service struct {
	ServiceName    string
	CommonMethods  []string
	JsonrpcMethods []string
	GrpcMethods    []string
	ResponseTypes  []string
	MethodNames    []string
}

type File struct {
	PackageName    string
	PackageImports map[string]struct{}
	Services       []Service
}
