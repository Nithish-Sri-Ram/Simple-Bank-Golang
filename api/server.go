package api

import (
	db "github.com/Nithish-Sri-Ram/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service/
type Server struct {
	store  *db.Store
	router *gin.Engine // This router will help us send each API request to the correct handler for processing
}

// New server creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount) // We can pass one or multiple handler functions - if we pass multiple functions, then the last one should be the real handler and all the other functions should be middlewares
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

// Start runs the http server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
