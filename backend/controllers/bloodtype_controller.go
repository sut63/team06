package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/team06/app/ent/bloodtype"

	"github.com/gin-gonic/gin"
	"github.com/team06/app/ent"
)

// BloodTypeController defines the struct for the bloodtype controller
type BloodTypeController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateBloodType handles POST requests for adding bloodtype entities
// @Summary Create bloodtype
// @Description Create bloodtype
// @ID create-bloodtype
// @Accept   json
// @Produce  json
// @Param bloodtype body ent.BloodType true "BloodType entity"
// @Success 200 {object} ent.BloodType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bloodtypes [post]
func (ctl *BloodTypeController) CreateBloodType(c *gin.Context) {
	obj := ent.BloodType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "bloodtype binding failed",
		})
		return
	}

	b, err := ctl.client.BloodType.
		Create().
		SetBlood(obj.Blood).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, b)
}

// GetBloodType handles GET requests to retrieve a bloodtype entity
// @Summary Get a bloodtype entity by ID
// @Description get bloodtype by ID
// @ID get-bloodtype
// @Produce  json
// @Param id path int true "BloodType ID"
// @Success 200 {object} ent.BloodType
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bloodtypes/{id} [get]
func (ctl *BloodTypeController) GetBloodType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	b, err := ctl.client.BloodType.
		Query().
		Where(bloodtype.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, b)
}

// ListBloodType handles request to get a list of bloodtype entities
// @Summary List bloodtype entities
// @Description list bloodtype entities
// @ID list-bloodtype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.BloodType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bloodtypes [get]
func (ctl *BloodTypeController) ListBloodType(c *gin.Context) {
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

	bloodtypes, err := ctl.client.BloodType.
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

	c.JSON(200, bloodtypes)
}

// DeleteBloodType handles DELETE requests to delete a bloodtype entity
// @Summary Delete a bloodtype entity by ID
// @Description get bloodtype by ID
// @ID delete-bloodtype
// @Produce  json
// @Param id path int true "BloodType ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bloodtypes/{id} [delete]
func (ctl *BloodTypeController) DeleteBloodType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.BloodType.
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

// UpdateBloodType handles PUT requests to update a bloodtype entity
// @Summary Update a bloodtype entity by ID
// @Description update bloodtype by ID
// @ID update-bloodtype
// @Accept   json
// @Produce  json
// @Param id path int true "BloodType ID"
// @Param bloodtype body ent.BloodType true "BloodType entity"
// @Success 200 {object} ent.BloodType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bloodtypes/{id} [put]
func (ctl *BloodTypeController) UpdateBloodType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.BloodType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "bloodtype binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	b, err := ctl.client.BloodType.
		UpdateOneID(int(id)).
		SetBlood(obj.Blood).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "update failed",
		})
		return
	}

	c.JSON(200, b)
}

// NewBloodTypeController creates and registers handles for the bloodtype controller
func NewBloodTypeController(router gin.IRouter, client *ent.Client) *BloodTypeController {
	bc := &BloodTypeController{
		client: client,
		router: router,
	}

	bc.register()

	return bc

}

func (ctl *BloodTypeController) register() {
	bloodtypes := ctl.router.Group("/bloodtypes")

	bloodtypes.GET("", ctl.ListBloodType)

	// CRUD
	bloodtypes.POST("", ctl.CreateBloodType)
	bloodtypes.GET(":id", ctl.GetBloodType)
	bloodtypes.PUT(":id", ctl.UpdateBloodType)
	bloodtypes.DELETE(":id", ctl.DeleteBloodType)
}
