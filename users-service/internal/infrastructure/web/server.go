package web

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router        *gin.Engine
	Handlers      UserHandlers
	WebServerPort string
}

func NewWebServer(serverPort string, handlers UserHandlers) *Server {
	return &Server{
		Router:        gin.New(),
		Handlers:      handlers,
		WebServerPort: serverPort,
	}
}

func (s *Server) Start() error {
	return s.Router.Run(s.WebServerPort)
}
