package handler

import (
	"context"

	dao "github.com/53jk1/todo-task/dao"
	pb "github.com/53jk1/todo-task/proto"
	"github.com/rs/zerolog/log"
)

type Service struct {
	*pb.UnimplementedTodoServiceServer
	Todo dao.Todo
}

func New() *Service {
	return &Service{}
}

func (s *Service) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	return &pb.CreateTodoResponse{
		Todo: &pb.Todo{
			Id: req.Todo.Id,
		},
	}, nil
}

func (s *Service) GetAllTodos(ctx context.Context, req *pb.GetAllTodosRequest) (*pb.GetAllTodosResponse, error) {
	log.Debug().Msg("GetTodo")
	return &pb.GetAllTodosResponse{
		Todos: []*pb.Todo{},
	}, nil
}