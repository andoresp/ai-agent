package main

import (
	product "ai/agents/cmd/api/product/infra/http"
)

func (a *App) routes() {
	product.ProductRoutes(a.api)
}
