package web

func SetupRoutes(server *Server) {
	routes := server.Router.Group("/api/users")
	routes.POST("", server.Handlers.CreateUser)
}
