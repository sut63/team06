package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team06/app/ent"
	"github.com/team06/app/ent/bloodtype"
	"github.com/team06/app/ent/gender"
	"github.com/team06/app/ent/patient"
	"github.com/team06/app/ent/prefix"
)

// PatientController defines the struct for the patient controller
type PatientController struct {
	client *ent.Client
	router gin.IRouter
}

// Patient defines the struct for the patient controller
type Patient struct {
	PersonalID     int
	Prefix         int
	PatientName    string
	Age            int
	Gender         int
	BloodType      int
	HospitalNumber string
	DrugAllergy    string
	AddedDate      string
}

// CreatePatient handles POST requests for adding patient entities
// @Summary Create Patient
// @Description Create Patient
// @ID create-Patient
// @Accept   json
// @Produce  json
// @Param Patient body ent.Patient true "Patient entity"
// @Success 200 {object} ent.Patient
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patients [post]
func (ctl *PatientController) CreatePatient(c *gin.Context) {
	obj := Patient{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "patient binding failed",
		})
		return
	}

	prefix, err := ctl.client.Prefix.
		Query().
		Where(prefix.IDEQ(int(obj.Prefix))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "prefix not found",
		})
		return
	}

	gender, err := ctl.client.Gender.
		Query().
		Where(gender.IDEQ(int(obj.Gender))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "gender not found",
		})
		return
	}

	bloodtype, err := ctl.client.BloodType.
		Query().
		Where(bloodtype.IDEQ(int(obj.BloodType))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "blood type not found",
		})
		return
	}

	addeddate, err := time.Parse(time.RFC3339, obj.AddedDate)

	patients, err := ctl.client.Patient.
		Create().
		SetPersonalID(obj.PersonalID).
		SetPrefix(prefix).
		SetPatientName(obj.PatientName).
		SetAge(obj.Age).
		SetGender(gender).
		SetBloodtype(bloodtype).
		SetHospitalNumber(obj.HospitalNumber).
		SetDrugAllergy(obj.DrugAllergy).
		SetAddedDate(addeddate).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   patients,
	})
}

// GetPatient handles GET requests to retrieve a patient entity
// @Summary Get a patient entity by ID
// @Description get patient by ID
// @ID get-patient
// @Produce  json
// @Param id path int true "Patient ID"
// @Success 200 {object} ent.Patient
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patients/{id} [get]
func (ctl *PatientController) GetPatient(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	patients, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, patients)
}

// ListPatient handles request to get a list of patient entities
// @Summary List patient entities
// @Description list patient entities
// @ID list-patient
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Patient
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patients [get]
func (ctl *PatientController) ListPatient(c *gin.Context) {
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

	patients, err := ctl.client.Patient.
		Query().
		WithPrefix().
		WithGender().
		WithBloodtype().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, patients)
}

// DeletePatient handles DELETE requests to delete a patient entity
// @Summary Delete a patient entity by ID
// @Description get patient by ID
// @ID delete-patient
// @Produce  json
// @Param id path int true "Patient ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patients/{id} [delete]
func (ctl *PatientController) DeletePatient(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Patient.
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

// UpdatePatient handles PUT requests to update a patient entity
// @Summary Update a patient entity by ID
// @Description update patient by ID
// @ID update-patient
// @Accept   json
// @Produce  json
// @Param id path int true "Patient ID"
// @Param patient body ent.Patient true "Patient entity"
// @Success 200 {object} ent.Patient
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patients/{id} [put]
func (ctl *PatientController) UpdatePatient(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Patient{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Patient binding failed",
		})
		return
	}
	obj.ID = int(id)
	patients, err := ctl.client.Patient.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, patients)
}

// NewPatientController creates and registers handles for the patient controller
func NewPatientController(router gin.IRouter, client *ent.Client) *PatientController {
	patientcontroller := &PatientController{
		client: client,
		router: router,
	}

	patientcontroller.register()

	return patientcontroller

}

// InitPatientController registers routes to the main engine
func (ctl *PatientController) register() {
	patients := ctl.router.Group("/patients")
	patients.GET("", ctl.ListPatient)

	//CRUD
	patients.POST("", ctl.CreatePatient)
	patients.PUT(":id", ctl.UpdatePatient)
	patients.GET(":id", ctl.GetPatient)
	patients.DELETE(":id", ctl.DeletePatient)

}
