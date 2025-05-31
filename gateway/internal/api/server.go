package api

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Server struct{
	cl Client
}

type Client interface {
	Create (ctx context.Context) error
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

	router.GET("/ping", func(c *gin.Context) {
    	c.JSON(200, gin.H{
    		"message": "pong",
    	})
  	})

	router.GET("/create", func(c *gin.Context) {
		err := s.Create(ctx)
    	c.JSON(200, gin.H{
    		"err": err,
    	})
  	})

  	router.Run() // listen and serve on 0.0.0.0:8080
}

func (s *Server) Create(ctx context.Context) error {
    return s.cl.Create(ctx)
} 