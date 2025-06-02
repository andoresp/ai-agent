package routes

import (
	"ai/agents/cmd/internal/appointment/infra/http/controllers"

	"github.com/gin-gonic/gin"
)

func NewAppointmentRouter(r *gin.Engine) {

	ac := controllers.NewAppointmentController()

	group := r.Group("/appointment")
	{
		group.GET("/", ac.Find)
		group.POST("/", ac.Make)
		group.GET("/available-hours", ac.ListAvailableHours)
		group.DELETE("/:id", ac.Cancel)
	}
}
