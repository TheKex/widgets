package api

import (
	"github.com/gin-gonic/gin"
	db "widgets/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{
		store: store,
	}
	router := gin.Default()

	// Routes begin
	router.POST("/users", server.createUser)

	// Routes end

	server.router = router
	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
