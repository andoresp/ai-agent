package main

import (
	"ai/agents/cmd/internal/appointment/infra/http/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	api := gin.Default()
	routes.NewAppointmentRouter(api)
	api.Run("localhost:8080")
}
