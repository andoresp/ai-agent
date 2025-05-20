package productHandler

import "github.com/gin-gonic/gin"

type FindProductByIdHandler struct {
}

func NewFindProductByIdHandler() *FindProductByIdHandler {
	return &FindProductByIdHandler{}
}

func (h *FindProductByIdHandler) Handle(c *gin.Context) {

}
