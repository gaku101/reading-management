package api

import (
	db "github.com/gaku101/my-portfolio/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// New Server creates a new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}

	server.setupRouter()
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(adress string) error {
	return server.router.Run(adress)
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)

	router.POST("/accounts", server.createAccount)
	router.GET	("/accounts/:id", server.getAccount)
	router.GET	("/accounts", server.listAccount)

	router.POST("/transfers", server.createTransfer)

	server.router = router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
