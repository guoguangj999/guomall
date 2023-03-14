package serviceAccount

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {

	c.JSON(http.StatusOK, "hello")

}
