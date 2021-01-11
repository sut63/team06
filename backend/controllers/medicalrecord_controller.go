package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/team06/app/ent/medicalrecord"

	"github.com/gin-gonic/gin"
	"github.com/team06/app/ent"
)

// MedicalRecordController defines the struct for the medicalrecord controller
type MedicalRecordController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateMedicalRecord handles POST requests for adding medicalrecord entities
// @Summary Create medicalrecord
// @Description Create medicalrecord
// @ID create-medicalrecord
// @Accept   json
// @Produce  json
// @Param medicalrecord body ent.MedicalRecord true "MedicalRecord entity"
// @Success 200 {object} ent.MedicalRecord
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalrecords [post]
func (ctl *MedicalRecordController) CreateMedicalRecord(c *gin.Context) {
	obj := ent.MedicalRecord{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "medicalrecord binding failed",
		})
		return
	}

	m, err := ctl.client.MedicalRecord.
		Create().
		SetEmail(obj.Email).
		SetName(obj.Name).
		SetPassword(obj.Password).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, m)
}

// GetMedicalRecord handles GET requests to retrieve a medicalrecord entity
// @Summary Get a medicalrecord entity by ID
// @Description get medicalrecord by ID
// @ID get-medicalrecord
// @Produce  json
// @Param id path int true "MedicalRecord ID"
// @Success 200 {object} ent.MedicalRecord
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalrecords/{id} [get]
func (ctl *MedicalRecordController) GetMedicalRecord(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	m, err := ctl.client.MedicalRecord.
		Query().
		Where(medicalrecord.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, m)
}

// ListMedicalRecord handles request to get a list of medicalrecord entities
// @Summary List medicalrecord entities
// @Description list medicalrecord entities
// @ID list-medicalrecord
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.MedicalRecord
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalrecords [get]
func (ctl *MedicalRecordController) ListMedicalRecord(c *gin.Context) {
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

	medicalrecords, err := ctl.client.MedicalRecord.
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

	c.JSON(200, medicalrecords)
}

// DeleteMedicalRecord handles DELETE requests to delete a medicalrecord entity
// @Summary Delete a medicalrecord entity by ID
// @Description get medicalrecord by ID
// @ID delete-medicalrecord
// @Produce  json
// @Param id path int true "MedicalRecord ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalrecords/{id} [delete]
func (ctl *MedicalRecordController) DeleteMedicalRecord(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.MedicalRecord.
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

// UpdateMedicalRecord handles PUT requests to update a medicalrecord entity
// @Summary Update a medicalrecord entity by ID
// @Description update medicalrecord by ID
// @ID update-medicalrecord
// @Accept   json
// @Produce  json
// @Param id path int true "MedicalRecord ID"
// @Param medicalrecord body ent.MedicalRecord true "MedicalRecord entity"
// @Success 200 {object} ent.MedicalRecord
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalrecords/{id} [put]
func (ctl *MedicalRecordController) UpdateMedicalRecord(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.MedicalRecord{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "medicalrecord binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	m, err := ctl.client.MedicalRecord.
		UpdateOneID(int(id)).
		SetEmail(obj.Email).
		SetName(obj.Name).
		SetPassword(obj.Password).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "update failed",
		})
		return
	}

	c.JSON(200, m)
}

// NewMedicalRecordController creates and registers handles for the medicalrecord controller
func NewMedicalRecordController(router gin.IRouter, client *ent.Client) *MedicalRecordController {
	mc := &MedicalRecordController{
		client: client,
		router: router,
	}

	mc.register()

	return mc

}

func (ctl *MedicalRecordController) register() {
	medicalrecords := ctl.router.Group("/medicalrecords")

	medicalrecords.GET("", ctl.ListMedicalRecord)

	// CRUD
	medicalrecords.POST("", ctl.CreateMedicalRecord)
	medicalrecords.GET(":id", ctl.GetMedicalRecord)
	medicalrecords.PUT(":id", ctl.UpdateMedicalRecord)
	medicalrecords.DELETE(":id", ctl.DeleteMedicalRecord)
}
