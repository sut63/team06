package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/team06/app/ent"
	"github.com/team06/app/ent/doctor"
	"github.com/team06/app/ent/patient"
	"github.com/team06/app/ent/proceduretype"
	"github.com/gin-gonic/gin"
)

// MedicalprocedureController defines the struct for the medicalprocedure controller
type MedicalprocedureController struct {
	client *ent.Client
	router gin.IRouter
}

// Medicalprocedure defines the struct for the dispensemedicine
type Medicalprocedure struct {
	Doctors int
	Patients int
	Proceduretypes  int
	Addedtime  string
}

// CreateMedicalprocedure handles POST requests for adding medicalprocedure entities
// @Summary Create medicalprocedure
// @Description Create medicalprocedure
// @ID create-medicalprocedure
// @Accept   json
// @Produce  json
// @Param medicalprocedure body Medicalprocedure true "Medicalprocedure entity"
// @Success 200 {object} ent.Medicalprocedure
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalprocedure [post]
func (ctl *MedicalprocedureController) CreateMedicalprocedure(c *gin.Context) {
	obj := Medicalprocedure{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "medicalprocedure binding failed",
		})
		return
	}

	doctors, err := ctl.client.Doctor.
		Query().
		Where(doctor.IDEQ(int(obj.Doctors))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Doctor not found",
		})
		return
	}

	patient, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(obj.Patients))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Patient not found",
		})
		return
	}

	procedure, err := ctl.client.ProcedureType.
		Query().
		Where(proceduretype.IDEQ(int(obj.Proceduretypes))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ProcedureType not found",
		})
		return
	}
	times, err := time.Parse(time.RFC3339, obj.Addedtime)

	medicalprocedure, err := ctl.client.MedicalProcedure.
		Create().
		SetDoctor(doctors).
		SetPatient(patient).
		SetProcedureType(procedure).
		SetAddtime(times).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, medicalprocedure)
}

// ListMedicalprocedure handles request to get a list of medicalprocedure entities
// @Summary List medicalprocedure entities
// @Description list medicalprocedure entities
// @ID list-medicalprocedure
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Medicalprocedure
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalprocedure [get]
func (ctl *MedicalprocedureController) ListMedicalprocedure(c *gin.Context) {
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

	medicalprocedures, err := ctl.client.MedicalProcedure.
		Query().
		WithDoctor().
		WithPatient().
		WithProcedureType().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, medicalprocedures)
}

// DeleteMedicalprocedure handles DELETE requests to delete a medicalprocedure entity
// @Summary Delete a medicalprocedure entity by ID
// @Description get medicalprocedure by ID
// @ID delete-medicalprocedure
// @Produce  json
// @Param id path int true "Medicalprocedure ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalprocedure/{id} [delete]
func (ctl *MedicalprocedureController) DeleteMedicalprocedure(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.MedicalProcedure.
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

// NewMedicalprocedureController creates and registers handles for the medicalprocedure controller
func NewMedicalprocedureController(router gin.IRouter, client *ent.Client) *MedicalprocedureController {
	mpc := &MedicalprocedureController{
		client: client,
		router: router,
	}
	mpc.register()
	return mpc
}

// InitUserController registers routes to the main engine
func (ctl *MedicalprocedureController) register() {
	medicalprocedures := ctl.router.Group("/medicalprocedure")

	medicalprocedures.GET("", ctl.ListMedicalprocedure)

	// CRUD
	medicalprocedures.POST("", ctl.CreateMedicalprocedure)
	medicalprocedures.DELETE(":id", ctl.DeleteMedicalprocedure)
}
