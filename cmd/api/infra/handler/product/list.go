package productHandler

import "github.com/gin-gonic/gin"

type ListProductsHandler struct {
}

func NewListProductsHandler() *ListProductsHandler {
	return &ListProductsHandler{}
}

func (h *ListProductsHandler) Handle(c *gin.Context) {
}
