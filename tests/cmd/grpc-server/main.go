package main

import (
	"log"
	"net"

	"github.com/najeal/rpc-fusion/tests/cmd"
	coreapiv1fusion "github.com/najeal/rpc-fusion/tests/gen/coreapi/v1/coreapifusion"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	coreapiv1fusion.RegisterGrpcCoreApiServiceServer(s, cmd.NewCommonServer())
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
