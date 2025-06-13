package api

import (
	"context"

	"github.com/elzamin/air_tickets/gateway/internal/entity"
	"github.com/gin-gonic/gin"
)

type Server struct{
	cl Client
}

type Client interface {
	Create (ctx context.Context, user entity.User) error
	Get (ctx context.Context, id string) (entity.User, error)
	GetAll (ctx context.Context) ([]entity.User, error)
	Delete (ctx context.Context, id string) error
	Update (ctx context.Context, user entity.User) error
}

func New (
	cl Client,
) *Server {
	return &Server{
		cl: cl,
	} 
}

func (s *Server) RunHTTPServer() {
	router := gin.Default()

	router.GET("/", home)
	router.GET("/ping", ping)
	router.GET("/id/:id", s.get)
	router.GET("/all", s.getAll)
	router.POST("/create", s.create)
	router.DELETE("/delete/:id", s.delete)
	router.PATCH("/update", s.update)

  	router.Run() // listen and serve on 0.0.0.0:8080
}

func ping (c *gin.Context) {
    	c.JSON(200, gin.H{
		"message": "pong",
		})
}

func home (c *gin.Context) {
	c.String(200, "Home page")
}

func (s *Server) get (c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")
	user, err := s.cl.Get(ctx, id)
	if err != nil{
		c.String(400, err.Error())
	} else {
		c.JSON(200, user)
	}
}

func (s *Server) getAll (c *gin.Context) {
	ctx := context.Background()
	users, err := s.cl.GetAll(ctx)
	if err != nil{
		c.String(400, err.Error())
	} else {
		c.JSON(200, users)
	}
}

func (s *Server) create (c *gin.Context) {
	ctx := context.Background()
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.String(400, err.Error())
	}
	err := s.cl.Create(ctx, user)
	if err != nil {
		c.String(400, err.Error())
	} else {
		c.JSON(200, "Created user with ID: '" + user.Id + "'")
	}
}

func (s *Server) delete (c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")
	err := s.cl.Delete(ctx, id)
	if err != nil{
		c.String(400, err.Error())
	} else {
		c.JSON(200, "Deleted user with ID: '" + id + "'")
	}
}

func (s *Server) update (c *gin.Context) {
	ctx := context.Background()
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.String(400, err.Error())
	}
	err := s.cl.Update(ctx, user)
	if err != nil {
		c.String(400, err.Error())
	} else {
		c.JSON(200, "Updated user with ID: '" + user.Id + "'")
	}
}