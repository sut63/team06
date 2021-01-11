package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/team06/app/ent"
	"github.com/team06/app/ent/righttotreatmenttype"
	"github.com/gin-gonic/gin"
)

// RightToTreatmentTypeController defines the struct for the righttotreatmenttype controller
type RightToTreatmentTypeController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateRightToTreatmentType handles POST requests for adding righttotreatmenttype entities
// @Summary Create righttotreatmenttype
// @Description Create righttotreatmenttype
// @ID create-righttotreatmenttype
// @Accept   json
// @Produce  json
// @Param righttotreatmenttype body ent.RightToTreatmentType true "RightToTreatmentType entity"
// @Success 200 {object} ent.RightToTreatmentType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /righttotreatmenttypes [post]
func (ctl *RightToTreatmentTypeController) CreateRightToTreatmentType(c *gin.Context) {
	obj := ent.RightToTreatmentType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "righttotreatmenttype binding failed",
		})
		return
	}

	t, err := ctl.client.RightToTreatmentType.
		Create().
		SetTypeName(obj.TypeName).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, t)
}


// GetRightToTreatmentType handles GET requests to retrieve a righttotreatmenttype entity
// @Summary Get a righttotreatmenttype entity by ID
// @Description get righttotreatmenttype by ID
// @ID get-righttotreatmenttype
// @Produce  json
// @Param id path int true "RightToTreatmentType ID"
// @Success 200 {object} ent.RightToTreatmentType
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /righttotreatmenttypes/{id} [get]
func (ctl *RightToTreatmentTypeController) GetRightToTreatmentType(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {

		c.JSON(400, gin.H{

			"error": err.Error(),
		})

		return

	}

	t, err := ctl.client.RightToTreatmentType.
		Query().
		Where(righttotreatmenttype.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {

		c.JSON(404, gin.H{

			"error": err.Error(),
		})

		return

	}

	c.JSON(200, t)

}

// ListRightToTreatmentType handles request to get a list of righttotreatmenttype entities
// @Summary List righttotreatmenttype entities
// @Description list righttotreatmenttype entities
// @ID list-righttotreatmenttype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.RightToTreatmentType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /righttotreatmenttypes [get]
func (ctl *RightToTreatmentTypeController) ListRightToTreatmentType(c *gin.Context) {

	limitQuery := c.Query("limit")

	limit := 10

	if limitQuery != "" {

		limit64, err := strconv.ParseInt(limitQuery, 10, 64)

		if err == nil {
			limit = int(limit64)
		}

	}

	offsetQuery := c.Query("offset")

	offset := 0

	if offsetQuery != "" {

		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)

		if err == nil {
			offset = int(offset64)
		}

	}

	RightToTreatmentTypes, err := ctl.client.RightToTreatmentType.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {

		c.JSON(400, gin.H{"error": err.Error()})

		return

	}

	c.JSON(200, RightToTreatmentTypes)

}

// DeleteRightToTreatmentType handles DELETE requests to delete a righttotreatmenttype entity
// @Summary Delete a righttotreatmenttype entity by ID
// @Description get righttotreatmenttype by ID
// @ID delete-righttotreatmenttype
// @Produce  json
// @Param id path int true "righttotreatmenttype ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /righttotreatmenttypes/{id} [delete]
func (ctl *RightToTreatmentTypeController) DeleteRightToTreatmentType(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {

		c.JSON(400, gin.H{

			"error": err.Error(),
		})

		return

	}

	err = ctl.client.RightToTreatmentType.
		DeleteOneID(int(id)).
		Exec(context.Background())

	if err != nil {

		c.JSON(404, gin.H{

			"error": err.Error(),
		})

		return

	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})

}

// UpdateRightToTreatmentType handles PUT requests to update a righttotreatmenttype entity
// @Summary Update a righttotreatmenttype entity by ID
// @Description update righttotreatmenttype by ID
// @ID update-righttotreatmenttype
// @Accept   json
// @Produce  json
// @Param id path int true "righttotreatmenttype ID"
// @Param righttotreatmenttype body ent.RightToTreatmentType true "RightToTreatmentType entity"
// @Success 200 {object} ent.RightToTreatmentType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /righttotreatmenttypes/{id} [put]
func (ctl *RightToTreatmentTypeController) UpdateRightToTreatmentType(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {

		c.JSON(400, gin.H{

			"error": err.Error(),
		})

		return

	}

	obj := ent.RightToTreatmentType{}

	if err := c.ShouldBind(&obj); err != nil {

		c.JSON(400, gin.H{

			"error": "righttotreatmenttype binding failed",
		})

		return

	}

	obj.ID = int(id)

	t, err := ctl.client.RightToTreatmentType.
		UpdateOne(&obj).
		Save(context.Background())

	if err != nil {

		c.JSON(400, gin.H{"error": "update failed"})

		return

	}

	c.JSON(200, t)

}

// NewRightToTreatmentTypeController creates and registers handles for the righttotreatmenttype controller
func NewRightToTreatmentTypeController(router gin.IRouter, client *ent.Client) *RightToTreatmentTypeController {

	tc := &RightToTreatmentTypeController{

		client: client,

		router: router,
	}

	tc.register()

	return tc

}

// InitRightToTreatmentTypeController registers routes to the main engine
func (ctl *RightToTreatmentTypeController) register() {
	righttotreatmenttypes := ctl.router.Group("/righttotreatmenttypes")
	righttotreatmenttypes.GET("", ctl.ListRightToTreatmentType)
	// CRUD
	righttotreatmenttypes.POST("", ctl.CreateRightToTreatmentType)
	righttotreatmenttypes.GET(":id", ctl.GetRightToTreatmentType)
	righttotreatmenttypes.PUT(":id", ctl.UpdateRightToTreatmentType)
	righttotreatmenttypes.DELETE(":id", ctl.DeleteRightToTreatmentType)
}