package serviceOrder

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListRequest struct {
	Page int64 `json:"page"`

	Limit int64 `json:"limit"`
}

type ListResponse struct {
	Items []*ListResponseItem
}

type ListResponseItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func List(c *gin.Context) {

	req := ListRequest{}

	c.ShouldBindJSON(&req)

	fmt.Println(req)

	//sql
	resp := ListResponse{}
	for i := 0; i < 5; i++ {
		resp.Items = append(resp.Items, &ListResponseItem{
			ID:   fmt.Sprintf("%d", i),
			Name: "name",
		})
	}

	c.JSON(http.StatusOK, &resp)

}
