package client

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	//"github.com/elzamin/air_tickets/gateway/internal/entity"
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

//func (c *Client) Create (ctx context.Context, uc pb.UserClient, user entity.User) error {
func (c *Client) Create (ctx context.Context) error {
	r, err := c.uc.CreateUser(ctx, &pb.CreateUserRequest{
		User: &pb.UserDTO{
			Id:      "1",
			Name:    "Dmitriy",
			Age:     22,
			Address: "Kitay-gorod",
			Work:    "Doctor",
		},
	})
	if err != nil {
		log.Printf("Receiver: %s", err.Error())
		return err
	}
	log.Printf("Receiver: %s", r.GetError().Message)
	
	return err
}