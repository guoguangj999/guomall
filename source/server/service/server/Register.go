package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	c.JSON(http.StatusOK, "hello")

}
