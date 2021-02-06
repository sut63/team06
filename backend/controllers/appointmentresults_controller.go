package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team06/app/ent"
	"github.com/team06/app/ent/appointmentresults"
	"github.com/team06/app/ent/doctor"
	"github.com/team06/app/ent/nurse"
	"github.com/team06/app/ent/patient"
	"github.com/team06/app/ent/room"
)

// AppointmentResultsController defines the struct for the AppointmentResults controller
type AppointmentResultsController struct {
	client *ent.Client
	router gin.IRouter
}

// AppointmentResults defines the struct for the AppointmentResults controller
type AppointmentResults struct {
	Patient     int
	Nurse       int
	Doctor      int
	Room        int
	Cause       string
	Advice      string
	Hour        string
	Minute      string
	DateAppoint string
	TimeAppoint string
	AddtimeSave string
}

// CreateAppointmentResults handles POST requests for adding AppointmentResults entities
// @Summary Create AppointmentResults
// @Description Create AppointmentResults
// @ID create-AppointmentResults
// @Accept   json
// @Produce  json
// @Param AppointmentResults body ent.AppointmentResults true "AppointmentResults entity"
// @Success 200 {object} ent.AppointmentResults
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /appointmentresultss [post]
func (ctl *AppointmentResultsController) CreateAppointmentResults(c *gin.Context) {
	obj := AppointmentResults{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "AppointmentResults binding failed",
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

	n, err := ctl.client.Nurse.
		Query().
		Where(nurse.IDEQ(int(obj.Nurse))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "nurse not found",
		})
		return
	}

	d, err := ctl.client.Doctor.
		Query().
		Where(doctor.IDEQ(int(obj.Doctor))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "doctor not found",
		})
		return
	}

	r, err := ctl.client.Room.
		Query().
		Where(room.IDEQ(int(obj.Room))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "room type not found",
		})
		return
	}

	save, err := time.Parse(time.RFC3339, obj.AddtimeSave)
	date, err := time.Parse(time.RFC3339, obj.DateAppoint)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "date appointment not found",
		})
		return
	}
	time, err := time.Parse(time.RFC3339, obj.TimeAppoint)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "time appointment not found",
		})
		return
	}
	hour, err := strconv.Atoi(obj.Hour)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "hour not found",
		})
		return
	}
	minute, err := strconv.Atoi(obj.Minute)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "minute not found",
		})
		return
	}
	a, err := ctl.client.AppointmentResults.
		Create().
		SetAppointmentResultsToPatient(p).
		SetAppointmentResultsToNurse(n).
		SetAppointmentResultsToDoctor(d).
		SetAppointmentResultsToRoom(r).
		SetCauseAppoint(obj.Cause).
		SetAdvice(obj.Advice).
		SetHourBeforeAppoint(hour).
		SetMinuteBeforeAppoint(minute).
		SetDateAppoint(date).
		SetTimeAppoint(time).
		SetAddtimeSave(save).
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
		"data":   a,
	})
}

// GetAppointmentResults handles GET requests to retrieve a appointmentresults entity
// @Summary Get a appointmentresults entity by ID
// @Description get appointmentresults by ID
// @ID get-appointmentresults
// @Produce  json
// @Param id path int true "AppointmentResults ID"
// @Success 200 {array} ent.AppointmentResults
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /appointmentresultss/{id} [get]
func (ctl *AppointmentResultsController) GetAppointmentResults(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	a, err := ctl.client.AppointmentResults.
		Query().
		Where(appointmentresults.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, a)
}

// ListAppointmentResults handles request to get a list of appointmentresults entities
// @Summary List appointmentresults entities
// @Description list appointmentresults entities
// @ID list-appointmentresults
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.AppointmentResults
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /appointmentresultss [get]
func (ctl *AppointmentResultsController) ListAppointmentResults(c *gin.Context) {
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

	appointmentresultss, err := ctl.client.AppointmentResults.
		Query().
		WithAppointmentResultsToPatient().
		WithAppointmentResultsToNurse().
		WithAppointmentResultsToDoctor().
		WithAppointmentResultsToRoom().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, appointmentresultss)
}

// DeleteAppointmentResults handles DELETE requests to delete a appointmentresults entity
// @Summary Delete a appointmentresults entity by ID
// @Description get appointmentresults by ID
// @ID delete-appointmentresults
// @Produce  json
// @Param id path int true "AppointmentResults ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /appointmentresultss/{id} [delete]
func (ctl *AppointmentResultsController) DeleteAppointmentResults(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.AppointmentResults.
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

// UpdateAppointmentResults handles PUT requests to update a appointmentresults entity
// @Summary Update a appointmentresults entity by ID
// @Description update appointmentresults by ID
// @ID update-appointmentresults
// @Accept   json
// @Produce  json
// @Param id path int true "AppointmentResults ID"
// @Param appointmentresults body ent.AppointmentResults true "AppointmentResults entity"
// @Success 200 {object} ent.AppointmentResults
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /appointmentresultss/{id} [put]
func (ctl *AppointmentResultsController) UpdateAppointmentResults(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.AppointmentResults{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "AppointmentResults binding failed",
		})
		return
	}
	obj.ID = int(id)
	a, err := ctl.client.AppointmentResults.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, a)
}

// NewAppointmentResultsController creates and registers handles for the appointmentresults controller
func NewAppointmentResultsController(router gin.IRouter, client *ent.Client) *AppointmentResultsController {
	ptc := &AppointmentResultsController{
		client: client,
		router: router,
	}

	ptc.register()

	return ptc

}

func (ctl *AppointmentResultsController) register() {
	appointmentresultss := ctl.router.Group("/appointmentresultss")

	appointmentresultss.POST("", ctl.CreateAppointmentResults)
	appointmentresultss.GET("", ctl.ListAppointmentResults)
	appointmentresultss.PUT(":id", ctl.UpdateAppointmentResults)
	appointmentresultss.GET(":id", ctl.GetAppointmentResults)
	appointmentresultss.DELETE(":id", ctl.DeleteAppointmentResults)

}
