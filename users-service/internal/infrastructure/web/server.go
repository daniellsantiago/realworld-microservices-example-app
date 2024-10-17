package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router        *gin.Engine
	Handlers      UserHandlers
	WebServerPort string
}

func NewWebServer(serverPort string, handlers UserHandlers) *Server {
	return &Server{
		Router:        gin.Default(),
		Handlers:      handlers,
		WebServerPort: serverPort,
	}
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.WebServerPort, s.Router)
}
