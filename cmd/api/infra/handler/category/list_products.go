package categoryHandler

import "github.com/gin-gonic/gin"

type ListProductsByCategoryParams struct {
	Name string `json:"name"`
}

type ListProductsByCategoryHandler struct {
	listProductsByCategoryUseCase any
}

func NewListProductsByCategoryHandler() *ListProductsByCategoryHandler {
	return &ListProductsByCategoryHandler{
		listProductsByCategoryUseCase: "",
	}
}

func (h *ListProductsByCategoryHandler) Handle(c *gin.Context) {}
