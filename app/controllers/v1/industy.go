// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// aahframework.org/examples source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"fmt"

	"aah-recycleview-backend/app/models"

	aah "aahframe.work"
)

// IndustryController implements the industry APi endpoints.
type IndustryController struct {
	*aah.Context
}

// List method returns all the industries available.
func (c *IndustryController) List() {
	industries := models.AllIndustries()
	c.Log().Infof("%v industries found", len(industries))
	c.Reply().JSON(aah.Data{
		"industries": industries,
	})
}

// Create method to create a industry via JSON.
func (c *IndustryController) Create(industry *models.Industry) {
	c.Log().Infof("Industry Info: %+v", *industry)
	id := models.CreateIndustry(industry)
	newResourceURL := fmt.Sprintf("%s:%v", c.Req.Scheme, c.RouteURL("retrieve_industry", id))
	c.Reply().Created().
		Header("Location", newResourceURL).
		JSON(aah.Data{
			"id": id,
		})
}

// Retrieve method retunrs single industry details for given industry ID.
func (c *IndustryController) Retrieve(id int64) {
	c.Log().Infof("Retrieving industry, ID: %v", id)
	industry, err := models.GetIndustry(id)
	if err != nil {
		c.Log().Errorf("Industry ID %v, %v", id, err)
		c.Reply().NotFound().JSON(aah.Data{
			"error": err.Error(),
		})
		return
	}

	c.Reply().JSON(models.ToIndustryAlias(industry))
}

// Update method updates the industry with given content.
func (c *IndustryController) Update(id int64, industry *models.Industry) {
	c.Log().Infof("Updating industry: %v", id)
	if industry.IndustryID == 0 {
		industry.IndustryID = id
	}

	if err := models.UpdateIndustry(industry); err != nil {
		c.Log().Errorf("Industry ID %v, %v", id, err)
		c.Reply().NotFound().JSON(aah.Data{
			"error": err.Error(),
		})
		return
	}

	c.Reply().NoContent()
}

// Delete method deletes the industry of given industry ID.
func (c *IndustryController) Delete(id int64) {
	c.Log().Infof("Deleting industry, ID: %v", id)
	if err := models.DeleteIndustry(id); err != nil {
		c.Log().Errorf("Industry ID %v, %v", id, err)
		c.Reply().NotFound().JSON(aah.Data{
			"error": err.Error(),
		})
		return
	}

	c.Reply().NoContent()
}
