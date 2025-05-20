package categoryHandler

import "github.com/gin-gonic/gin"

type FindCategoryByIdHandler struct {
}

func NewFindCategoryByIdHandler() *FindCategoryByIdHandler {
	return &FindCategoryByIdHandler{}
}

func (h *FindCategoryByIdHandler) Handle(c *gin.Context) {}
