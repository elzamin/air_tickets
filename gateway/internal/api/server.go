package api

import (
	"context"
	//"net/http"

	"github.com/elzamin/air_tickets/gateway/internal/entity"
	"github.com/gin-gonic/gin"
)

type Server struct{
	cl Client
}

type Client interface {
	Create (ctx context.Context, user entity.User) error
	Get (ctx context.Context, id string) (entity.User, error)
	Delete (ctx context.Context, id string) error
}

func New (
	cl Client,
) *Server {
	return &Server{
		cl: cl,
	} 
}

func (s *Server) RunHTTPServer(ctx context.Context) {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "This is my home page")
	})

	router.GET("/ping", func(c *gin.Context) {
    	c.JSON(200, gin.H{
    		"message": "pong",
    	})
  	})

	router.GET("/id/:id", func(c *gin.Context) {
		id := c.Param("id")
		user, err := s.Get(ctx, id)
		if err != nil{
			c.String(400, err.Error())
		}
    	c.JSON(200, user)
  	})

	router.POST("/create", func(c *gin.Context) {
		var user entity.User
		if err := c.ShouldBindJSON(&user); err != nil {
    		c.String(400, err.Error())
  		}
		err := s.Create(ctx, user)
    	c.JSON(200, "Deleted user with ID: " + err.Error())
  	})

	router.DELETE("/delete/:id", func(c *gin.Context) {
		id := c.Param("id")
		err := s.Delete(ctx, id)
		if err != nil{
			c.String(400, err.Error())
		}
    	c.JSON(200, "Deleted user with ID: " + id)
  	})

  	router.Run() // listen and serve on 0.0.0.0:8080
}

func (s *Server) Create (ctx context.Context, user entity.User) error {
    return s.cl.Create(ctx, user)
} 

func (s *Server) Get (ctx context.Context, id string) (entity.User, error) {
	return s.cl.Get(ctx, id)
}

func (s *Server) Delete (ctx context.Context, id string) error {
	return s.cl.Delete(ctx, id)
}