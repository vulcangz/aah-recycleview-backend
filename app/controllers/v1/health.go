package v1

import (
	aah "aahframe.work"
)

// ValueController is kickstart sample for API implementation.
type HealthController struct {
	BaseController
}

// Index method is application's root API endpoint.
func (c *HealthController) RdbHealth() {
	h, err := Repo.RdbHealthCheck()
	if err != nil {
		c.Reply().Ok().JSON(aah.Error{
			Reason:  err,
			Message: "It seems that the database is not working properly. Please check it!",
			Code:    500,
		})
	}

	c.Reply().Ok().JSON(h)
}
