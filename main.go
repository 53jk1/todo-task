package main

import (
	"io/ioutil"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"google.golang.org/grpc/credentials/insecure"

	db "github.com/53jk1/todo-task/dao/db"

	"github.com/53jk1/todo-task/gateway"
	"github.com/53jk1/todo-task/handler"
	"github.com/53jk1/todo-task/insecure"
	pb "github.com/53jk1/todo-task/proto"
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

	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable"
	gdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	handler := &handler.Service{
		UnimplementedTodoServiceServer: &pb.UnimplementedTodoServiceServer{},
		Todo:                           &db.Todo{Db: gdb},
	}

	s := grpc.NewServer(
		// TODO: Replace with your own certificate!
		// grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	)
	pb.RegisterTodoServiceServer(s, handler)

	// Serve gRPC Server
	log.Info("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	err = gateway.Run("dns:///" + addr)
	log.Fatalln(err)
}
