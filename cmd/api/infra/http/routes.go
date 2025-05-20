package routes

import (
	category "ai/agents/cmd/api/infra/handler/category"
	product "ai/agents/cmd/api/infra/handler/product"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	r.GET("/product", product.NewListProductsHandler().Handle)
	r.GET("/product/:id", product.NewFindProductByIdHandler().Handle)
	r.POST("/product", product.NewCreateProductHandler().Handle)
	r.DELETE("/product/:id", product.NewDeleteProductHandler().Handle)
}

func CategoryRoutes(r *gin.Engine) {
	r.GET("/category", category.NewListCategoriesHandler().Handle)
	r.GET("/category/:id", category.NewFindCategoryByIdHandler().Handle)
	r.GET("/category/:id/products", category.NewListProductsByCategoryHandler().Handle)
	r.POST("/category", category.NewCreateCategoryHandler().Handle)
	r.DELETE("/category", category.NewDeleteCategoryHandler().Handle)
}
