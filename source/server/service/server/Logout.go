package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogoutRequest struct {
}

type LogoutResponse struct {
}

func Logout(c *gin.Context) {

	c.JSON(http.StatusOK, "hello")

}
