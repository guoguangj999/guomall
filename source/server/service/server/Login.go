package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
}

type LoginResponse struct {
}

func Login(c *gin.Context) {

	c.JSON(http.StatusOK, "hello")

}
