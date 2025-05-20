package categoryHandler

import "github.com/gin-gonic/gin"

type DeleteCategoryHandler struct {
}

func NewDeleteCategoryHandler() *DeleteCategoryHandler {
	return &DeleteCategoryHandler{}
}

func (h *DeleteCategoryHandler) Handle(c *gin.Context) {}
