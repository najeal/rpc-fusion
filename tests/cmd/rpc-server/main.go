package main

import (
	"log"
	"net/http"

	"github.com/gorilla/rpc/v2"
	json "github.com/gorilla/rpc/v2/json2"

	"github.com/najeal/rpc-fusion/tests/cmd"
	coreapiv1fusion "github.com/najeal/rpc-fusion/tests/gen/coreapi/v1/coreapifusion"
)

func main() {
	jrpcServer := rpc.NewServer()
	jrpcServer.RegisterCodec(json.NewCodec(), "application/json")
	coreapiv1fusion.RegisterJsonrpcCoreApiService(jrpcServer, cmd.NewCommonServer())
	http.Handle("/rpc", jrpcServer)
	log.Println("Starting server on :8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
