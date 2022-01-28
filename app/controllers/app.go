package controllers

import (
	"aahframe.work"

	"aah-recycleview-backend/app/models"
)

// AppController struct application controller
type AppController struct {
	*aah.Context
}

// Index method is application's root API endpoint.
func (c *AppController) Index() {
	c.Reply().Ok().JSON(models.Greet{
		Message: "Welcome to aah framework - API application",
	})
}
