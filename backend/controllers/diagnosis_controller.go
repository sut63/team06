package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team06/app/ent"
	"github.com/team06/app/ent/diagnosis"
	"github.com/team06/app/ent/doctor"
	"github.com/team06/app/ent/patient"
	"github.com/team06/app/ent/treatmenttype"
)

// DiagnosisController defines the struct for the deoartment controller
type DiagnosisController struct {
	client *ent.Client
	router gin.IRouter
}

// Diagnosis defines the struct for the Diagnosis controller
type Diagnosis struct {
	Symptom       string
	Opinionresult string
	diagnosisDate string
	TreatmentType int
	Patient       int
	Doctor        int
	Note          string
}

// CreateDiagnosis handles POST requests for adding Diagnosis entities
// @Summary Create Diagnosis
// @Description Create Diagnosis
// @ID create-Diagnosis
// @Accept   json
// @Produce  json
// @Param Diagnosis body ent.Diagnosis true "Diagnosis entity"
// @Success 200 {object} ent.Diagnosis
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Diagnosiss [post]
func (ctl *DiagnosisController) CreateDiagnosis(c *gin.Context) {
	obj := Diagnosis{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Diagnosis binding failed",
		})
		return
	}

	patient, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(obj.Patient))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Patient not found",
		})
		return
	}

	doctor, err := ctl.client.Doctor.
		Query().
		Where(doctor.IDEQ(int(obj.Doctor))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Doctor not found",
		})
		return
	}

	ty, err := ctl.client.TreatmentType.
		Query().
		Where(treatmenttype.IDEQ(int(obj.TreatmentType))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Treatment type not found",
		})
		return
	}
	diagnosisdate, err := time.Parse(time.RFC3339, obj.diagnosisDate)

	Diagnosiss, err := ctl.client.Diagnosis.
		Create().
		SetSymptom(obj.Symptom).
		SetDiagnosisDate(diagnosisdate).
		SetOpinionresult(obj.Opinionresult).
		SetPatient(patient).
		SetType(ty).
		SetNote(obj.Note).
		SetDoctorName(doctor).
		Save(context.Background())

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   Diagnosiss,
	})
}

// GetDiagnosis handles GET requests to retrieve a Diagnosis entity
// @Summary Get a Diagnosis entity by ID
// @Description get Diagnosis by ID
// @ID get-Diagnosis
// @Produce  json
// @Param id path int true "Diagnosis ID"
// @Success 200 {object} ent.Diagnosis
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Diagnosiss/{id} [get]
func (ctl *DiagnosisController) GetDiagnosis(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	Diagnosiss, err := ctl.client.Diagnosis.
		Query().
		Where(diagnosis.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, Diagnosiss)
}

// ListDiagnosis handles request to get a list of Diagnosis entities
// @Summary List Diagnosis entities
// @Description list Diagnosis entities
// @ID list-Diagnosis
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Diagnosis
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Diagnosiss [get]
func (ctl *DiagnosisController) ListDiagnosis(c *gin.Context) {
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

	Diagnosiss, err := ctl.client.Diagnosis.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, Diagnosiss)
}

// DeleteDiagnosis handles DELETE requests to delete a Diagnosis entity
// @Summary Delete a Diagnosis entity by ID
// @Description get Diagnosis by ID
// @ID delete-Diagnosis
// @Produce  json
// @Param id path int true "Diagnosis ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Diagnosiss/{id} [delete]
func (ctl *DiagnosisController) DeleteDiagnosis(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Diagnosis.
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

// UpdateDiagnosis handles PUT requests to update a Diagnosis entity
// @Summary Update a Diagnosis entity by ID
// @Description update Diagnosis by ID
// @ID update-Diagnosis
// @Accept   json
// @Produce  json
// @Param id path int true "Diagnosis ID"
// @Param Diagnosis body ent.Diagnosis true "Diagnosis entity"
// @Success 200 {object} ent.Diagnosis
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Diagnosiss/{id} [put]
func (ctl *DiagnosisController) UpdateDiagnosis(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Diagnosis{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Diagnosis binding failed",
		})
		return
	}

	obj.ID = int(id)
	Diagnosiss, err := ctl.client.Diagnosis.
		UpdateOne(&obj).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, Diagnosiss)
}

// NewDiagnosisController creates and registers handles for the Diagnosis controller
func NewDiagnosisController(router gin.IRouter, client *ent.Client) *DiagnosisController {
	DiagnosisController := &DiagnosisController{
		client: client,
		router: router,
	}

	DiagnosisController.register()
	return DiagnosisController
}

// InitDiagnosisController registers routes to the main engine
func (ctl *DiagnosisController) register() {
	Diagnosiss := ctl.router.Group("/Diagnosiss")
	Diagnosiss.GET("", ctl.ListDiagnosis)

	// CRUD
	Diagnosiss.POST("", ctl.CreateDiagnosis)
	Diagnosiss.GET(":id", ctl.GetDiagnosis)
	Diagnosiss.PUT(":id", ctl.UpdateDiagnosis)
	Diagnosiss.DELETE(":id", ctl.DeleteDiagnosis)
}
