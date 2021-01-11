package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team06/app/ent"
	"github.com/team06/app/ent/doctor"
)

// DoctorController struct
type DoctorController struct {
	client *ent.Client
	router gin.IRouter
}


// CreateDoctor handles POST requests for adding Doctor entities
// @Summary Create doctor
// @Description Create doctor
// @ID create-doctor
// @Accept   json
// @Produce  json
// @Param Doctor body Doctor true "Doctor entity"
// @Success 200 {object} ent.Doctor
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctor [post]
func (ctl *DoctorController) CreateDoctor(c *gin.Context) {
	obj := ent.Doctor{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Doctor binding failed",
		})
		return
	}


	doctor, err := ctl.client.Doctor.
		Create().
		SetDoctorName(obj.DoctorName).
		SetDoctorUsername(obj.DoctorUsername).
		SetDoctorPassword(obj.DoctorPassword).
		Save(context.Background())
	fmt.Println(err)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, doctor)
}

// GetDoctor handles GET requests to retrieve a Doctor entity
// @Summary Get a Doctor entity by ID
// @Description get Doctor by ID
// @ID get-Doctor
// @Produce  json
// @Param id path int true "Doctor ID"
// @Success 200 {object} ent.Doctor
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctor/{id} [get]
func (ctl *DoctorController) GetDoctor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	doctor, err := ctl.client.Doctor.
		Query().
		Where(doctor.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, doctor)
}
// ListDoctor handles request to get a list of Doctor entities
// @Summary List Doctor entities
// @Description list Doctor entities
// @ID list-Doctor
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Doctor
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctor [get]
func (ctl *DoctorController) ListDoctor(c *gin.Context) {
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

	doctor, err := ctl.client.Doctor.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, doctor)
}

// NewDoctorController creates and registers handles for the Doctor controller
func NewDoctorController(router gin.IRouter, client *ent.Client) *DoctorController {
	dc := &DoctorController{
		client: client,
		router: router,
	}
	dc.register()

	return dc

}

func (ctl *DoctorController) register() {
	doctors := ctl.router.Group("/doctor")

	doctors.POST("", ctl.CreateDoctor)
	doctors.GET("", ctl.ListDoctor)
	doctors.GET(":id", ctl.GetDoctor)
	

}
