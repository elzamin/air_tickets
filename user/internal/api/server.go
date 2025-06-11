package api

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	//"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/status"

	pb "github.com/elzamin/air_tickets/proto/gen/go"
	"github.com/elzamin/air_tickets/user/internal/entity"
)

type Server struct {
	pb.UnimplementedUserServer
	userSvc UserService
}

type UserService interface {
	Create(ctx context.Context, user entity.User) error
	Get(ctx context.Context, id string) (entity.User, error)
	GetAll(ctx context.Context) ([]entity.User, error)
	Update(ctx context.Context, user entity.User) error
	Delete(ctx context.Context, id string) error
}

func New(
	userSvc UserService,
) *Server {
	return &Server{
		userSvc: userSvc,
	}
}

func (s *Server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.Printf("Received: Create %v", in.GetUser())
	user := entity.User{
		Id:      in.GetUser().GetId(),
		Name:    in.GetUser().GetName(),
		Age:     int(in.GetUser().GetAge()),
		Address: in.GetUser().GetAddress(),
		Work:    in.GetUser().GetWork(),
	}

	err := s.userSvc.Create(ctx, user)
	if err != nil {
		return nil, err
		//return &pb.CreateUserResponse{Error: &pb.Error{Message: err.Error()}}, nil
	}

	return nil, nil
	//return &pb.CreateUserResponse{Error: &pb.Error{Message: "Created ID: '" + in.GetUser().GetId() + "'"}}, nil
}

func (s *Server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Printf("Received: Get ID%v", in.GetId())

	user, err := s.userSvc.Get(ctx, in.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		User: &pb.UserDTO{
			Id:      user.Id,
			Name:    user.Name,
			Age:     int32(user.Age),
			Address: user.Address,
			Work:    user.Work,
		},
	}, nil
}

func (s *Server) GetUsers(ctx context.Context, in *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	log.Print("Received: Get all")

	users, err := s.userSvc.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	usersResp := make([]*pb.UserDTO, len(users))
	for i, el := range users {
		usersResp[i] = &pb.UserDTO{
			Id:      el.Id,
			Name:    el.Name,
			Age:     int32(el.Age),
			Address: el.Address,
			Work:    el.Work,
		}
	}

	return &pb.GetUsersResponse{
		Users: usersResp,
	}, nil
}

func (s *Server) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	log.Printf("Received: Update %v", in.GetUser())
	user := entity.User{
		Id:      in.GetUser().GetId(),
		Name:    in.GetUser().GetName(),
		Age:     int(in.GetUser().GetAge()),
		Address: in.GetUser().GetAddress(),
		Work:    in.GetUser().GetWork(),
	}

	err := s.userSvc.Update(ctx, user)
	if err != nil {
		return nil, err
		//return &pb.UpdateUserResponse{Error: &pb.Error{Message: err.Error()}}, nil
	}

	return nil, nil
	//return &pb.UpdateUserResponse{Error: nil}, nil
}

func (s *Server) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	log.Printf("Received: Delete ID%v", in.GetId())

	err := s.userSvc.Delete(ctx, in.GetId())
	if err != nil {
		return nil, err
		//return &pb.DeleteUserResponse{Error: &pb.Error{Message: err.Error()}}, nil
	}
	return nil, nil
	//return &pb.DeleteUserResponse{Error: nil}, nil
}

func RunGRPCServer(port string, testServer *Server) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServer(s, testServer)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
