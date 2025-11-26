package request

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Post_Api_Category_Request struct {
	Name        string `json:"name" validate:"required, min=1, max=128"`
	Description string `json:"description" validate:"min=0, max=1024"`
}

type Get_Api_Category_Detail_Request struct {
	CategoryID uint `json:"category_id" validate:"required"`
}

func (req Get_Api_Category_Detail_Request) BindQueryWithGinContext(c *gin.Context) error {
	query_category_id := c.Query("category_id")

	category_id, err := strconv.ParseUint(query_category_id, 10, 64)
	if err != nil {
		return err
	}

	req = Get_Api_Category_Detail_Request{
		CategoryID: uint(category_id),
	}

	return nil
}
