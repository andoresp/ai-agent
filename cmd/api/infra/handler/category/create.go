package categoryHandler

import "github.com/gin-gonic/gin"

type CreateCategoryHandler struct {
}

func NewCreateCategoryHandler() *CreateCategoryHandler {
	return &CreateCategoryHandler{}
}

func (h *CreateCategoryHandler) Handle(c *gin.Context) {}
