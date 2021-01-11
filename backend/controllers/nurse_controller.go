package controllers

import (
	"context"
	"fmt"
	"strconv"
	"github.com/team06/app/ent"
	"github.com/team06/app/ent/nurse"
	"github.com/gin-gonic/gin"
 )
 
 
 // NurseController defines the struct for the deoartment controller
 type NurseController struct {
	client *ent.Client
	router gin.IRouter
 }


// CreateNurse handles POST requests for adding nurse entities
// @Summary Create nurse
// @Description Create nurse
// @ID create-nurse
// @Accept   json
// @Produce  json
// @Param nurse body ent.Nurse true "Nurse entity"
// @Success 200 {object} ent.Nurse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /nurses [post]
func (ctl *NurseController) CreateNurse(c *gin.Context) {
	obj := ent.Nurse{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "nurse binding failed",
		})
		return
	}

	nurses, err := ctl.client.Nurse.
		Create().
		SetNurseName(obj.NurseName).
		SetNurseUsername(obj.NurseUsername).
		SetNursePassword(obj.NursePassword).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, nurses)
 }


// GetNurse handles GET requests to retrieve a nurse entity
// @Summary Get a nurse entity by ID
// @Description get nurse by ID
// @ID get-nurse
// @Produce  json
// @Param id path int true "Nurse ID"
// @Success 200 {object} ent.Nurse
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /nurses/{id} [get]
func (ctl *NurseController) GetNurse(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
 
	nurses, err := ctl.client.Nurse.
		Query().
		Where(nurse.IDEQ(int(id))).
		Only(context.Background())

	if err != nil { 
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
 
	c.JSON(200, nurses)
 }


// ListNurse handles request to get a list of nurse entities
// @Summary List nurse entities
// @Description list nurse entities
// @ID list-nurse
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Nurse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /nurses [get]
func (ctl *NurseController) ListNurse(c *gin.Context) {
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
 
	nurses, err := ctl.client.Nurse.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
			if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
 
	c.JSON(200, nurses)
 }


// DeleteNurse handles DELETE requests to delete a nurse entity
// @Summary Delete a nurse entity by ID
// @Description get nurse by ID
// @ID delete-nurse
// @Produce  json
// @Param id path int true "Nurse ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /nurses/{id} [delete]
func (ctl *NurseController) DeleteNurse(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
 
	err = ctl.client.Nurse.
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


// UpdateNurse handles PUT requests to update a nurse entity
// @Summary Update a nurse entity by ID
// @Description update nurse by ID
// @ID update-nurse
// @Accept   json
// @Produce  json
// @Param id path int true "Nurse ID"
// @Param nurse body ent.Nurse true "Nurse entity"
// @Success 200 {object} ent.Nurse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /nurses/{id} [put]
func (ctl *NurseController) UpdateNurse(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
 
	obj := ent.Nurse{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "nurse binding failed",
		})
		return
	}

	obj.ID = int(id) 
	nurses, err := ctl.client.Nurse.
		UpdateOne(&obj).
		Save(context.Background())

	if err != nil { 
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
 
	c.JSON(200, nurses)
 }



// NewNurseController creates and registers handles for the nurse controller
func NewNurseController(router gin.IRouter, client *ent.Client) *NurseController {
	nursecontroller := &NurseController{
		client: client,
		router: router,
	}

	nursecontroller.register()
	return nursecontroller
 }
 
// InitNurseController registers routes to the main engine 
func (ctl *NurseController) register() {
	nurses := ctl.router.Group("/nurses")
	nurses.GET("", ctl.ListNurse)
 
	// CRUD
	nurses.POST("", ctl.CreateNurse)
	nurses.GET(":id", ctl.GetNurse)
	nurses.PUT(":id", ctl.UpdateNurse)
	nurses.DELETE(":id", ctl.DeleteNurse)
 }