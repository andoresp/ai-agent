package categoryHandler

import "github.com/gin-gonic/gin"

type ListCategoriesHandler struct {
}

func NewListCategoriesHandler() *ListCategoriesHandler {
	return &ListCategoriesHandler{}
}

func (h *ListCategoriesHandler) Handle(c *gin.Context) {}
