package handler

import (
	pb "github.com/53jk1/todo-task/proto"
)

type Service struct {
	EmbeddedService pb.TodoServiceServer
}

func New() *Service {
	return &Service{}
}