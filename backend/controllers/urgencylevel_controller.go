package controllers

import (
	"context"
	"fmt"
	"strconv"
	"github.com/team06/app/ent"
	"github.com/team06/app/ent/urgencylevel"
	"github.com/gin-gonic/gin"
 )
 
 
 // UrgencyLevelController defines the struct for the deoartment controller
 type UrgencyLevelController struct {
	client *ent.Client
	router gin.IRouter
 }


// CreateUrgencyLevel handles POST requests for adding urgencylevel entities
// @Summary Create urgencylevel
// @Description Create urgencylevel
// @ID create-urgencylevel
// @Accept   json
// @Produce  json
// @Param urgencylevel body ent.UrgencyLevel true "UrgencyLevel entity"
// @Success 200 {object} ent.UrgencyLevel
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /urgencylevels [post]
func (ctl *UrgencyLevelController) CreateUrgencyLevel(c *gin.Context) {
	obj := ent.UrgencyLevel{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "urgencylevel binding failed",
		})
		return
	}

	urgencylevels, err := ctl.client.UrgencyLevel.
		Create().
		SetUrgencyName(obj.UrgencyName).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, urgencylevels)
 }


// GetUrgencyLevel handles GET requests to retrieve a urgencylevel entity
// @Summary Get a urgencylevel entity by ID
// @Description get urgencylevel by ID
// @ID get-urgencylevel
// @Produce  json
// @Param id path int true "UrgencyLevel ID"
// @Success 200 {object} ent.UrgencyLevel
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /urgencylevels/{id} [get]
func (ctl *UrgencyLevelController) GetUrgencyLevel(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
 
	urgencylevels, err := ctl.client.UrgencyLevel.
		Query().
		Where(urgencylevel.IDEQ(int(id))).
		Only(context.Background())

	if err != nil { 
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
 
	c.JSON(200, urgencylevels)
 }


// ListUrgencyLevel handles request to get a list of urgencylevel entities
// @Summary List urgencylevel entities
// @Description list urgencylevel entities
// @ID list-urgencylevel
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.UrgencyLevel
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /urgencylevels [get]
func (ctl *UrgencyLevelController) ListUrgencyLevel(c *gin.Context) {
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
 
	urgencylevels, err := ctl.client.UrgencyLevel.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
			if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
 
	c.JSON(200, urgencylevels)
 }


// DeleteUrgencyLevel handles DELETE requests to delete a urgencylevel entity
// @Summary Delete a urgencylevel entity by ID
// @Description get urgencylevel by ID
// @ID delete-urgencylevel
// @Produce  json
// @Param id path int true "UrgencyLevel ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /urgencylevels/{id} [delete]
func (ctl *UrgencyLevelController) DeleteUrgencyLevel(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
 
	err = ctl.client.UrgencyLevel.
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


// UpdateUrgencyLevel handles PUT requests to update a urgencylevel entity
// @Summary Update a urgencylevel entity by ID
// @Description update urgencylevel by ID
// @ID update-urgencylevel
// @Accept   json
// @Produce  json
// @Param id path int true "UrgencyLevel ID"
// @Param urgencylevel body ent.UrgencyLevel true "UrgencyLevel entity"
// @Success 200 {object} ent.UrgencyLevel
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /urgencylevels/{id} [put]
func (ctl *UrgencyLevelController) UpdateUrgencyLevel(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
 
	obj := ent.UrgencyLevel{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "urgencylevel binding failed",
		})
		return
	}

	obj.ID = int(id) 
	urgencylevels, err := ctl.client.UrgencyLevel.
		UpdateOne(&obj).
		Save(context.Background())

	if err != nil { 
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
 
	c.JSON(200, urgencylevels)
 }



// NewUrgencyLevelController creates and registers handles for the urgencylevel controller
func NewUrgencyLevelController(router gin.IRouter, client *ent.Client) *UrgencyLevelController {
	urgencylevelcontroller := &UrgencyLevelController{
		client: client,
		router: router,
	}

	urgencylevelcontroller.register()
	return urgencylevelcontroller
 }
 
// InitUrgencyLevelController registers routes to the main engine 
func (ctl *UrgencyLevelController) register() {
	urgencylevels := ctl.router.Group("/urgencylevels")
	urgencylevels.GET("", ctl.ListUrgencyLevel)
 
	// CRUD
	urgencylevels.POST("", ctl.CreateUrgencyLevel)
	urgencylevels.GET(":id", ctl.GetUrgencyLevel)
	urgencylevels.PUT(":id", ctl.UpdateUrgencyLevel)
	urgencylevels.DELETE(":id", ctl.DeleteUrgencyLevel)
 }