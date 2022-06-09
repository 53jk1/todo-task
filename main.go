package main

import (
	"io/ioutil"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	db "github.com/53jk1/todo-task/dao/db"

	"github.com/53jk1/todo-task/gateway"
	"github.com/53jk1/todo-task/handler"
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

	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable"
	dsn := os.Getenv("DSN")
	gdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to database:", err)
	}

	handler := &handler.Service{
		UnimplementedTodoServiceServer: &pb.UnimplementedTodoServiceServer{},
		Todo:                           &db.Todo{Db: gdb},
	}

	s := grpc.NewServer()

	pb.RegisterTodoServiceServer(s, handler)

	// Serve gRPC Server
	log.Info("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	err = gateway.Run("dns:///" + addr)
	log.Fatalln(err)
}
