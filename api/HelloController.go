package api

import (
	"aspgo/core"
	"aspgo/services"
	"net/http"
)

type HelloController struct {
	core.ControllerBase

	userService *services.UserService
}

func NewHelloController() *HelloController {
	return &HelloController{}
}

func (c *HelloController) HandleHello(w http.ResponseWriter, r *http.Request) {

	userInfo := c.userService.GetUser()

	w.Write([]byte(userInfo))
}
