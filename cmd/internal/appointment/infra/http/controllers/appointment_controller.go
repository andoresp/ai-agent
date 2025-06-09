package controllers

import (
	usecases "ai/agents/cmd/internal/appointment/core/use_cases"
	"ai/agents/cmd/internal/appointment/infra/repositories"
	"ai/agents/cmd/internal/appointment/infra/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppointmentController struct {
	findAppointmentUseCase               usecases.FindAppointmentUseCase
	makeAppointmentUseCase               usecases.MakeAppointmentUseCase
	cancelAppointmentUseCase             usecases.CancelAppointmentUseCase
	listAvailableAppointmentHoursUseCase usecases.ListAvailableAppointmentHoursUseCase
}

func (ac *AppointmentController) Find(c *gin.Context) {
	cn := c.DefaultQuery("client_name", "")
	cgi := c.DefaultQuery("client_government_id", "")

	dto := usecases.FindAppointmentDto{
		ClientName:           cn,
		ClientGovernmentName: cgi,
	}

	app, err := ac.findAppointmentUseCase.Execute(dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":                   app.Id(),
		"date":                 app.Date().UTC(),
		"client_name":          app.ClientName(),
		"client_government_id": app.ClientGovernmentId(),
	})
}

func (ac *AppointmentController) Make(c *gin.Context) {
	var dto usecases.MakeAppointmentDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	app, err := ac.makeAppointmentUseCase.Execute(dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":                   app.Id(),
		"date":                 app.Date().UTC(),
		"client_name":          app.ClientName(),
		"client_government_id": app.ClientGovernmentId(),
	})
}

func (ac *AppointmentController) Cancel(c *gin.Context) {
	id := c.Param("id")

	dto := usecases.CancelAppointmentDto{Id: id}

	err := ac.cancelAppointmentUseCase.Execute(dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (ac *AppointmentController) ListAvailableHours(c *gin.Context) {
	hrs, err := ac.listAvailableAppointmentHoursUseCase.Execute()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, hrs)
}

func NewAppointmentController() *AppointmentController {
	repo := repositories.NewAppointmentRepositoryInMemory()
	service := services.NewAppointmentHoursService(3, 8)

	return &AppointmentController{
		findAppointmentUseCase:               *usecases.NewFindAppointmentUseCase(&repo),
		makeAppointmentUseCase:               *usecases.NewMakeAppoinmentUseCase(&repo, &service),
		cancelAppointmentUseCase:             *usecases.NewCancelAppointmentUseCase(&repo),
		listAvailableAppointmentHoursUseCase: *usecases.NewListAvailableAppointmentHoursUseCase(&repo, &service),
	}
}
