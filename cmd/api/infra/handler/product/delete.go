package productHandler

import "github.com/gin-gonic/gin"

type DeleteProductHandler struct {
}

func NewDeleteProductHandler() *DeleteProductHandler {
	return &DeleteProductHandler{}
}

func (h *DeleteProductHandler) Handle(c *gin.Context) {

}
