package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"github.com/team06/app/ent"
	"github.com/gin-gonic/gin"

	"github.com/team06/app/ent/triageresult"
	"github.com/team06/app/ent/department"
	"github.com/team06/app/ent/urgencylevel"
	"github.com/team06/app/ent/patient"
	"github.com/team06/app/ent/nurse"
 )


 // TriageResultController defines the struct for the deoartment controller
 type TriageResultController struct {
	client *ent.Client
	router gin.IRouter
 }

// TriageResult defines the struct for the TriageResult controller
type TriageResult struct {
	Symptom    string
	TriageDate string
	Patient    int
	Nurse int
	Department int
	UrgencyLevel int
}


// CreateTriageResult handles POST requests for adding TriageResult entities
// @Summary Create TriageResult
// @Description Create TriageResult
// @ID create-TriageResult
// @Accept   json
// @Produce  json
// @Param TriageResult body ent.TriageResult true "TriageResult entity"
// @Success 200 {object} ent.TriageResult
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /TriageResults [post]
func (ctl *TriageResultController) CreateTriageResult(c *gin.Context) {
	obj := TriageResult{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "triageresult binding failed",
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

	nurse, err := ctl.client.Nurse.
		Query().
		Where(nurse.IDEQ(int(obj.Nurse))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Nurse not found",
		})
		return
	}

	department, err := ctl.client.Department.
		Query().
		Where(department.IDEQ(int(obj.Department))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Department not found",
		})
		return
	}

	urgencylevel, err := ctl.client.UrgencyLevel.
		Query().
		Where(urgencylevel.IDEQ(int(obj.UrgencyLevel))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "UrgencyLevel not found",
		})
		return
	}

	triagedate, err := time.Parse(time.RFC3339, obj.TriageDate)

	triageresults, err := ctl.client.TriageResult.
		Create().
		SetSymptom(obj.Symptom).
		SetTriageDate(triagedate).
		SetTriageResultToDepartment(department).
		SetTriageResultToNurse(nurse).
		SetTriageResultToPatient(patient).
		SetTriageResultToUrgencyLevel(urgencylevel).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   triageresults,
	})
 }


// GetTriageResult handles GET requests to retrieve a triageresult entity
// @Summary Get a triageresult entity by ID
// @Description get triageresult by ID
// @ID get-triageresult
// @Produce  json
// @Param id path int true "TriageResult ID"
// @Success 200 {object} ent.TriageResult
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /triageresults/{id} [get]
func (ctl *TriageResultController) GetTriageResult(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
 
	triageresults, err := ctl.client.TriageResult.
		Query().
		Where(triageresult.IDEQ(int(id))).
		Only(context.Background())

	if err != nil { 
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
 
	c.JSON(200, triageresults)
 }


// ListTriageResult handles request to get a list of triageresult entities
// @Summary List triageresult entities
// @Description list triageresult entities
// @ID list-triageresult
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.TriageResult
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /triageresults [get]
func (ctl *TriageResultController) ListTriageResult(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {limit = int(limit64)}
	}
 
	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {offset = int(offset64)}
	}
 
	triageresults, err := ctl.client.TriageResult.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
			if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
 
	c.JSON(200, triageresults)
 }


// DeleteTriageResult handles DELETE requests to delete a triageresult entity
// @Summary Delete a triageresult entity by ID
// @Description get triageresult by ID
// @ID delete-triageresult
// @Produce  json
// @Param id path int true "TriageResult ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /triageresults/{id} [delete]
func (ctl *TriageResultController) DeleteTriageResult(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
 
	err = ctl.client.TriageResult.
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


// UpdateTriageResult handles PUT requests to update a triageresult entity
// @Summary Update a triageresult entity by ID
// @Description update triageresult by ID
// @ID update-triageresult
// @Accept   json
// @Produce  json
// @Param id path int true "TriageResult ID"
// @Param triageresult body ent.TriageResult true "TriageResult entity"
// @Success 200 {object} ent.TriageResult
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /triageresults/{id} [put]
func (ctl *TriageResultController) UpdateTriageResult(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
 
	obj := ent.TriageResult{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "triageresult binding failed",
		})
		return
	}

	obj.ID = int(id) 
	triageresults, err := ctl.client.TriageResult.
		UpdateOne(&obj).
		Save(context.Background())

	if err != nil { 
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
 
	c.JSON(200, triageresults)
 }



// NewTriageResultController creates and registers handles for the triageresult controller
func NewTriageResultController(router gin.IRouter, client *ent.Client) *TriageResultController {
	triageresultcontroller := &TriageResultController{
		client: client,
		router: router,
	}

	triageresultcontroller.register()
	return triageresultcontroller
 }
 
// InitTriageResultController registers routes to the main engine 
func (ctl *TriageResultController) register() {
	triageresults := ctl.router.Group("/triageresults")
	triageresults.GET("", ctl.ListTriageResult)
 
	// CRUD
	triageresults.POST("", ctl.CreateTriageResult)
	triageresults.GET(":id", ctl.GetTriageResult)
	triageresults.PUT(":id", ctl.UpdateTriageResult)
	triageresults.DELETE(":id", ctl.DeleteTriageResult)
 }