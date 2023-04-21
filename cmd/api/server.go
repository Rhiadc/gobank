package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/rhiadc/gobank/db/sqlc"
)

type Server struct {
	store  db.StoreInterface
	router *gin.Engine
}

func NewServer(store db.StoreInterface) *Server {
	r := gin.Default()

	server := &Server{
		router: r,
		store:  store,
	}

	//register new custom validation
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	r.POST("/accounts", server.createAccount)
	r.GET("/accounts/:id", server.getAccount)
	r.GET("/accounts", server.listAccount)
	r.POST("/transfers", server.createTransfer)

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
