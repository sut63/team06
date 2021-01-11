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

// AppointmentresultsController defines the struct for the appointmentresults controller
type AppointmentresultsController struct {
	client *ent.Client
	router gin.IRouter
}

// Appointmentresults defines the struct for the appointmentresults controller
type Appointmentresults struct {
	Patient        int
	Nurse          int
	Doctor         int
	Room           int
	AddtimeAppoint string
	AddtimeSave    string
}

// CreateAppointmentresults handles POST requests for adding appointmentresults entities
// @Summary Create appointmentresults
// @Description Create appointmentresults
// @ID create-appointmentresults
// @Accept   json
// @Produce  json
// @Param appointmentresults body ent.Appointmentresults true "Appointmentresults entity"
// @Success 200 {object} ent.Appointmentresults
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /appointmentresultss [post]
func (ctl *AppointmentresultsController) CreateAppointmentresults(c *gin.Context) {
	obj := Appointmentresults{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "appointmentresults binding failed",
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

	timeappoint, err := time.Parse(time.RFC3339, obj.AddtimeAppoint)
	timesave, err := time.Parse(time.RFC3339, obj.AddtimeSave)

	a, err := ctl.client.AppointmentResults.
		Create().
		SetAddtimeAppoint(timeappoint).
		SetAddtimeSave(timesave).
		SetAppointmentResultsToPatient(p).
		SetAppointmentResultsToNurse(n).
		SetAppointmentResultsToDoctor(d).
		SetAppointmentResultsToRoom(r).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, a)
}

// GetAppointmentresults handles GET requests to retrieve a appointmentresults entity
// @Summary Get a appointmentresults entity by ID
// @Description get appointmentresults by ID
// @ID get-appointmentresults
// @Produce  json
// @Param id path int true "Appointmentresults ID"
// @Success 200 {object} ent.Appointmentresults
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /appointmentresultss/{id} [get]
func (ctl *AppointmentresultsController) GetAppointmentresults(c *gin.Context) {
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

// ListAppointmentresults handles request to get a list of appointmentresults entities
// @Summary List appointmentresults entities
// @Description list appointmentresults entities
// @ID list-appointmentresults
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Appointmentresults
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /appointmentresultss [get]
func (ctl *AppointmentresultsController) ListAppointmentresults(c *gin.Context) {
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

// DeleteAppointmentresults handles DELETE requests to delete a appointmentresults entity
// @Summary Delete a appointmentresults entity by ID
// @Description get appointmentresults by ID
// @ID delete-appointmentresults
// @Produce  json
// @Param id path int true "Appointmentresults ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /appointmentresultss/{id} [delete]
func (ctl *AppointmentresultsController) DeleteAppointmentresults(c *gin.Context) {
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

// UpdateAppointmentresults handles PUT requests to update a appointmentresults entity
// @Summary Update a appointmentresults entity by ID
// @Description update appointmentresults by ID
// @ID update-appointmentresults
// @Accept   json
// @Produce  json
// @Param id path int true "Appointmentresults ID"
// @Param appointmentresults body ent.Appointmentresults true "Appointmentresults entity"
// @Success 200 {object} ent.Appointmentresults
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /appointmentresultss/{id} [put]
func (ctl *AppointmentresultsController) UpdateAppointmentresults(c *gin.Context) {
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
			"error": "Appointmentresults binding failed",
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

// NewAppointmentresultsController creates and registers handles for the appointmentresults controller
func NewAppointmentresultsController(router gin.IRouter, client *ent.Client) *AppointmentresultsController {
	ptc := &AppointmentresultsController{
		client: client,
		router: router,
	}

	ptc.register()

	return ptc

}

func (ctl *AppointmentresultsController) register() {
	appointmentresultss := ctl.router.Group("/appointmentresults")

	appointmentresultss.POST("", ctl.CreateAppointmentresults)
	appointmentresultss.GET("", ctl.ListAppointmentresults)
	appointmentresultss.PUT(":id", ctl.UpdateAppointmentresults)
	appointmentresultss.GET(":id", ctl.GetAppointmentresults)
	appointmentresultss.DELETE(":id", ctl.DeleteAppointmentresults)

}
