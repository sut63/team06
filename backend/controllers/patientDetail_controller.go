package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team06/app/ent"
	"github.com/team06/app/ent/bloodtype"
	"github.com/team06/app/ent/gender"
	"github.com/team06/app/ent/patient"
	"github.com/team06/app/ent/patientdetail"
	"github.com/team06/app/ent/prefix"
)

// PatientDetailController defines the struct for the patient-detail controller
type PatientDetailController struct {
	client *ent.Client
	router gin.IRouter
}

// PatientDetail defines the struct for the patient-detail controller
type PatientDetail struct {
	Prefix    int
	Gender    int
	BloodType int
	Patient   int
}

// CreatePatientDetail handles POST requests for adding patient-detail entities
// @Summary Create patient-detail
// @Description Create patient-detail
// @ID create-patient-detail
// @Accept   json
// @Produce  json
// @Param patient-detail body ent.PatientDetail true "PatientDetail entity"
// @Success 200 {object} ent.PatientDetail
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patient-details [post]
func (ctl *PatientDetailController) CreatePatientDetail(c *gin.Context) {
	obj := PatientDetail{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "patient detail binding failed",
		})
		return
	}

	pf, err := ctl.client.Prefix.
		Query().
		Where(prefix.IDEQ(int(obj.Prefix))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "prefix not found",
		})
		return
	}

	g, err := ctl.client.Gender.
		Query().
		Where(gender.IDEQ(int(obj.Gender))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "gender not found",
		})
		return
	}

	b, err := ctl.client.BloodType.
		Query().
		Where(bloodtype.IDEQ(int(obj.BloodType))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "bloodtype not found",
		})
		return
	}

	pt, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(obj.Patient))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "patient not found",
		})
		return
	}

	pd, err := ctl.client.PatientDetail.
		Create().
		SetPrefix(pf).
		SetGender(g).
		SetBloodtype(b).
		SetPatient(pt).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"patient": true,
		"data":    pd,
	})
}

// GetPatientDetail handles GET requests to retrieve a patientdetail entity
// @Summary Get a patientdetail entity by ID
// @Description get patientdetail by ID
// @ID get-patientdetail
// @Produce  json
// @Param id path int true "PatientDetail ID"
// @Success 200 {object} ent.PatientDetail
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientdetails/{id} [get]
func (ctl *PatientDetailController) GetPatientDetail(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	pd, err := ctl.client.PatientDetail.
		Query().
		Where(patientdetail.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, pd)
}

// ListPatientDetail handles request to get a list of patient-detail entities
// @Summary List patient-detail entities
// @Description list patient-detail entities
// @ID list-patient-detail
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.PatientDetail
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patient-details [get]
func (ctl *PatientDetailController) ListPatientDetail(c *gin.Context) {
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

	patientdetails, err := ctl.client.PatientDetail.
		Query().
		WithPrefix().
		WithGender().
		WithBloodtype().
		WithPatient().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, patientdetails)
}

// DeletePatientDetail handles DELETE requests to delete a patientdetail entity
// @Summary Delete a patientdetail entity by ID
// @Description get patientdetail by ID
// @ID delete-patientdetail
// @Produce  json
// @Param id path int true "PatientDetail ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientdetails/{id} [delete]
func (ctl *PatientDetailController) DeletePatientDetail(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.PatientDetail.
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

// UpdatePatientDetail handles PUT requests to update a patientdetail entity
// @Summary Update a patientdetail entity by ID
// @Description update patientdetail by ID
// @ID update-patientdetail
// @Accept   json
// @Produce  json
// @Param id path int true "PatientDetail ID"
// @Param patientdetail body ent.PatientDetail true "PatientDetail entity"
// @Success 200 {object} ent.PatientDetail
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientdetails/{id} [put]
func (ctl *PatientDetailController) UpdatePatientDetail(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.PatientDetail{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "PatientDetail binding failed",
		})
		return
	}
	obj.ID = int(id)
	pd, err := ctl.client.PatientDetail.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, pd)
}

// NewPatientDetailController creates and registers handles for the patient-detail controller
func NewPatientDetailController(router gin.IRouter, client *ent.Client) *PatientDetailController {
	pdc := &PatientDetailController{
		client: client,
		router: router,
	}

	pdc.register()

	return pdc

}

func (ctl *PatientDetailController) register() {
	patientdetails := ctl.router.Group("/patient-details")

	patientdetails.POST("", ctl.CreatePatientDetail)
	patientdetails.GET("", ctl.ListPatientDetail)
	patientdetails.PUT(":id", ctl.UpdatePatientDetail)
	patientdetails.GET(":id", ctl.GetPatientDetail)
	patientdetails.DELETE(":id", ctl.DeletePatientDetail)

}
