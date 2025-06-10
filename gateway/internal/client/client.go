package client

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/elzamin/air_tickets/gateway/internal/entity"
	pb "github.com/elzamin/air_tickets/proto/gen/go"
)

type Client struct{
	uc pb.UserClient
}

func New(
	uc pb.UserClient,
) *Client{
		return &Client{
			uc: uc,
		}
	}

func ConnectToUserGRPC(port string) pb.UserClient{
	conn, err := grpc.NewClient("localhost:" + port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer conn.Close()
	log.Printf("Connected to User [port: %s]", port)
	return pb.NewUserClient(conn)
}

func (c *Client) Create (ctx context.Context, user entity.User) error {
	r, err := c.uc.CreateUser(ctx, &pb.CreateUserRequest{
		User: &pb.UserDTO{
			Id:      user.Id,
			Name:    user.Name,
			Age:     int32(user.Age),
			Address: user.Address,
			Work:    user.Work,
		},
	})
	if err != nil {
		log.Printf("Response with error: %s", err.Error())
		return err
	}
	
	return errors.New(r.GetError().Message)
}

func (c *Client) Get (ctx context.Context, id string) (entity.User, error) {
	r, err := c.uc.GetUser(ctx, &pb.GetUserRequest{Id: id})
	if err != nil {
		log.Printf("Response with error: %s", err.Error())
		return entity.User{}, err
	}

	user := entity.User{
		Id: r.User.Id, 
		Name: r.User.Name, 
		Age: int(r.User.Age),
		Address: r.User.Address,
		Work: r.User.Work,
	}

	return user, err
}

func (c *Client) Delete (ctx context.Context, id string) error {
	_, err := c.uc.DeleteUser(ctx, &pb.DeleteUserRequest{Id: id})
	if err != nil {
		log.Printf("Response with error: %s", err.Error())
		return err
	}

	return nil
}