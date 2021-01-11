package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team06/app/ent"
	"github.com/team06/app/ent/hospital"
	"github.com/team06/app/ent/patient"
	"github.com/team06/app/ent/righttotreatment"
	"github.com/team06/app/ent/righttotreatmenttype"
)

// RightToTreatmentController defines the struct for the righttotreatment controller
type RightToTreatmentController struct {
	client *ent.Client
	router gin.IRouter
}

// RightToTreatment defines the struct for the righttotreatment controller
type RightToTreatment struct {
	Hospital             int
	Patient              int
	RightToTreatmentType int
	Starttime            string
	Endtime              string
}

// CreateRightToTreatment handles POST requests for adding righttotreatment entities
// @Summary Create righttotreatment
// @Description Create righttotreatment
// @ID create-righttotreatment
// @Accept   json
// @Produce  json
// @Param righttotreatment body ent.RightToTreatment true "RightToTreatment entity"
// @Success 200 {object} ent.RightToTreatment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /righttotreatments [post]
func (ctl *RightToTreatmentController) CreateRightToTreatment(c *gin.Context) {
	obj := RightToTreatment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "righttotreatment binding failed",
		})
		return
	}

	h, err := ctl.client.Hospital.
		Query().
		Where(hospital.IDEQ(int(obj.Hospital))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "hospital not found",
		})
		return
	}

	p, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(obj.Patient))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "patient not found",
		})
		return
	}

	t, err := ctl.client.RightToTreatmentType.
		Query().
		Where(righttotreatmenttype.IDEQ(int(obj.RightToTreatmentType))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "righttotreatment not found",
		})
		return
	}

	s, err := time.Parse(time.RFC3339, obj.Starttime)
	e, err := time.Parse(time.RFC3339, obj.Endtime)

	r, err := ctl.client.RightToTreatment.
		Create().
		SetStarttime(s).
		SetEndtime(e).
		SetHospital(h).
		SetPatient(p).
		SetRightToTreatmentType(t).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, r)
}

// GetRightToTreatment handles GET requests to retrieve a righttotreatment entity
// @Summary Get a righttotreatment entity by ID
// @Description get righttotreatment by ID
// @ID get-righttotreatment
// @Produce  json
// @Param id path int true "RightToTreatment ID"
// @Success 200 {object} ent.RightToTreatment
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /righttotreatments/{id} [get]
func (ctl *RightToTreatmentController) GetRightToTreatment(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {

		c.JSON(400, gin.H{

			"error": err.Error(),
		})

		return

	}

	r, err := ctl.client.RightToTreatment.
		Query().
		Where(righttotreatment.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {

		c.JSON(404, gin.H{

			"error": err.Error(),
		})

		return

	}

	c.JSON(200, r)

}

// ListRightToTreatment handles request to get a list of righttotreatment entities
// @Summary List righttotreatment entities
// @Description list righttotreatment entities
// @ID list-righttotreatment
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.RightToTreatment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /righttotreatments [get]
func (ctl *RightToTreatmentController) ListRightToTreatment(c *gin.Context) {

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

	RightToTreatments, err := ctl.client.RightToTreatment.
		Query().
		WithHospital().
		WithPatient().
		WithRightToTreatmentType().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {

		c.JSON(400, gin.H{"error": err.Error()})

		return

	}

	c.JSON(200, RightToTreatments)

}

// DeleteRightToTreatment handles DELETE requests to delete a righttotreatment entity
// @Summary Delete a righttotreatment entity by ID
// @Description get righttotreatment by ID
// @ID delete-righttotreatment
// @Produce  json
// @Param id path int true "RightToTreatment ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /righttotreatments/{id} [delete]
func (ctl *RightToTreatmentController) DeleteRightToTreatment(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {

		c.JSON(400, gin.H{

			"error": err.Error(),
		})

		return

	}

	err = ctl.client.RightToTreatment.
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

// UpdateRightToTreatment handles PUT requests to update a righttotreatment entity
// @Summary Update a righttotreatment entity by ID
// @Description update righttotreatment by ID
// @ID update-righttotreatment
// @Accept   json
// @Produce  json
// @Param id path int true "righttotreatment ID"
// @Param righttotreatment body ent.RightToTreatment true "RightToTreatment entity"
// @Success 200 {object} ent.RightToTreatment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /righttotreatments/{id} [put]
func (ctl *RightToTreatmentController) UpdateRightToTreatment(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {

		c.JSON(400, gin.H{

			"error": err.Error(),
		})

		return

	}

	obj := ent.RightToTreatment{}

	if err := c.ShouldBind(&obj); err != nil {

		c.JSON(400, gin.H{

			"error": "righttotreatment binding failed",
		})

		return

	}

	obj.ID = int(id)

	r, err := ctl.client.RightToTreatment.
		UpdateOne(&obj).
		Save(context.Background())

	if err != nil {

		c.JSON(400, gin.H{"error": "update failed"})

		return

	}

	c.JSON(200, r)

}

// NewRightToTreatmentController creates and registers handles for the righttotreatment controller
func NewRightToTreatmentController(router gin.IRouter, client *ent.Client) *RightToTreatmentController {

	rc := &RightToTreatmentController{

		client: client,

		router: router,
	}

	rc.register()

	return rc

}

// InitRightToTreatmentController registers routes to the main engine
func (ctl *RightToTreatmentController) register() {
	righttotreatments := ctl.router.Group("/righttotreatments")
	righttotreatments.GET("", ctl.ListRightToTreatment)
	// CRUD
	righttotreatments.POST("", ctl.CreateRightToTreatment)
	righttotreatments.GET(":id", ctl.GetRightToTreatment)
	righttotreatments.PUT(":id", ctl.UpdateRightToTreatment)
	righttotreatments.DELETE(":id", ctl.DeleteRightToTreatment)
}
