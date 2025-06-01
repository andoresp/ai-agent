package main

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	api *gin.Engine
}

func main() {
	app := App{
		api: gin.Default(),
	}

	app.api.Run("localhost:8080")
}
