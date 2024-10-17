package web

import (
	"github.com/gin-gonic/gin"
)

type UserHandlers interface {
	SayHello(c *gin.Context)
}

type userHandlers struct {
}

func NewUserHandlers() UserHandlers {
	return &userHandlers{}
}

func (h userHandlers) SayHello(c *gin.Context) {

}
