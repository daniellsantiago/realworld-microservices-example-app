package web

import "github.com/gin-gonic/gin"

func SetupRoutes(server *Server) {
	// Health check endpoint for LB
	server.Router.GET("/health", func(c *gin.Context) {
		c.Status(200)
	})

	routes := server.Router.Group("/api/users")
	routes.POST("", server.Handlers.CreateUser)
}
