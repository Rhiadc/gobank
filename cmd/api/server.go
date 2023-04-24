package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/rhiadc/gobank/config"
	db "github.com/rhiadc/gobank/db/sqlc"
	token "github.com/rhiadc/gobank/internal/token"
	rtoken "github.com/rhiadc/gobank/internal/token/paseto"
)

type Server struct {
	store       db.StoreInterface
	router      *gin.Engine
	token       token.Maker
	tokenConfig config.Token
}

func NewServer(config *config.Environments, store db.StoreInterface) (*Server, error) {
	token, err := rtoken.NewPasetoMaker(config.Token.TokenSynmmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		store:       store,
		token:       token,
		tokenConfig: config.Token,
	}

	server.setupRouter()
	//register new custom validation
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	return server, nil
}

func (server *Server) setupRouter() {
	r := gin.Default()
	r.POST("/accounts", server.createAccount)
	r.GET("/accounts/:id", server.getAccount)
	r.GET("/accounts", server.listAccount)
	r.POST("/transfers", server.createTransfer)
	r.POST("/users/login", server.loginUser)
	server.router = r
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
