package main

import (
	"io/ioutil"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	"github.com/53jk1/todo-task/gateway"
	"github.com/53jk1/todo-task/insecure"
	pbExample "github.com/53jk1/todo-task/proto"
	"github.com/53kj1/todo-task/server"
)

func main() {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	addr := "0.0.0.0:10000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer(
		// TODO: Replace with your own certificate!
		grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
	)
	pbExample.RegisterUserServiceServer(s, server.New())

	// Serve gRPC Server
	log.Info("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	err = gateway.Run("dns:///" + addr)
	log.Fatalln(err)
}