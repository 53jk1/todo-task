package handler

import (
	"context"

	pb "github.com/53jk1/todo-task/proto"
)

type Service struct {
	*pb.UnimplementedTodoServiceServer	
}

func New() *Service {
	return &Service{}
}

func (s *Service) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	
}
