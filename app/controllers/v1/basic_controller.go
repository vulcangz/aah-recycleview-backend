package v1

import (
	"aah-recycleview-backend/app/database"
	"aah-recycleview-backend/app/repository"

	aah "aahframe.work"
)

var Repo repository.Repository

// BaseController for admin controllers.
// Created for any common abstraction of admin controllers.
type BaseController struct {
	*aah.Context
}

// Before method is an interceptor for admin path.
func (c *BaseController) Before() {
	c.Log().Info("BaseController before interceptor is called")
	var err error
	Repo, err = repository.NewSormRepository(database.SDB())

	if err != nil {
		c.Reply().InternalServerError().Text("500 Internal Server Error")
		c.Abort()
		return
	}
}
