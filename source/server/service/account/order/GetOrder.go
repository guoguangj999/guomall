package serviceOrder

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrder(c *gin.Context) {

	c.JSON(http.StatusOK, "hello")

}
