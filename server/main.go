package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"grpc-test/api"
	"net"

	"google.golang.org/grpc"
)

var Port int

func main() {
	Port = 3000
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", Port))
	if err != nil {
		logrus.WithError(err).Fatal("Couldn't start http server")
	}

	apiserver := api.Server{}

	grpcServer := grpc.NewServer()

	api.RegisterGreetServer(grpcServer, &apiserver)

	logrus.Infof("Starting server on port %d", Port)

	if err := grpcServer.Serve(listener); err != nil {
		logrus.WithError(err).Fatalf("Failed to serve on port %d", Port)
	}
}