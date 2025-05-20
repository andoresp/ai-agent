package productHandler

import "github.com/gin-gonic/gin"

type CreateProductParams struct {
	Name string `json:"name"`
}

type CreateProductHandler struct {
	listProductUseCase any
}

func NewCreateProductHandler() *CreateProductHandler {
	return &CreateProductHandler{
		listProductUseCase: "",
	}
}

func (h *CreateProductHandler) Handle(c *gin.Context) {

}
