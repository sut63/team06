package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/team06/app/ent/treatmenttype"

	"github.com/gin-gonic/gin"
	"github.com/team06/app/ent"
)

// TreatmenttypeController defines the struct for the TreatmentType controller
type TreatmenttypeController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateTreatmentType handles POST requests for adding TreatmentType entities
// @Summary Create TreatmentType
// @Description Create TreatmentType
// @ID create-TreatmentType
// @Accept   json
// @Produce  json
// @Param TreatmentType body ent.TreatmentType true "TreatmentType entity"
// @Success 200 {object} ent.TreatmentType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /TreatmentTypes [post]
func (ctl *TreatmenttypeController) CreateTreatmentType(c *gin.Context) {
	obj := ent.TreatmentType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "TreatmentType binding failed",
		})
		return
	}

	t, err := ctl.client.TreatmentType.
		Create().
		SetType(obj.Type).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, t)
}

// GetTreatmentType handles GET requests to retrieve a TreatmentType entity
// @Summary Get a TreatmentType entity by ID
// @Description get TreatmentType by ID
// @ID get-TreatmentType
// @Produce  json
// @Param id path int true "TreatmentType ID"
// @Success 200 {object} ent.TreatmentType
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /TreatmentTypes/{id} [get]
func (ctl *TreatmenttypeController) GetTreatmentType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	t, err := ctl.client.TreatmentType.
		Query().
		Where(treatmenttype.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, t)
}

// ListTreatmentType handles request to get a list of TreatmentType entities
// @Summary List TreatmentType entities
// @Description list TreatmentType entities
// @ID list-TreatmentType
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.TreatmentType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /TreatmentTypes [get]
func (ctl *TreatmenttypeController) ListTreatmentType(c *gin.Context) {
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

	TreatmentTypes, err := ctl.client.TreatmentType.
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

	c.JSON(200, TreatmentTypes)
}

// DeleteTreatmentType handles DELETE requests to delete a TreatmentType entity
// @Summary Delete a TreatmentType entity by ID
// @Description get TreatmentType by ID
// @ID delete-TreatmentType
// @Produce  json
// @Param id path int true "TreatmentType ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /TreatmentTypes/{id} [delete]
func (ctl *TreatmenttypeController) DeleteTreatmentType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.TreatmentType.
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

// UpdateTreatmentType handles PUT requests to update a TreatmentType entity
// @Summary Update a TreatmentType entity by ID
// @Description update TreatmentType by ID
// @ID update-TreatmentType
// @Accept   json
// @Produce  json
// @Param id path int true "TreatmentType ID"
// @Param TreatmentType body ent.TreatmentType true "TreatmentType entity"
// @Success 200 {object} ent.TreatmentType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /TreatmentTypes/{id} [put]
func (ctl *TreatmenttypeController) UpdateTreatmentType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.TreatmentType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "TreatmentType binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	t, err := ctl.client.TreatmentType.
		UpdateOneID(int(id)).
		SetType(obj.Type).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "update failed",
		})
		return
	}

	c.JSON(200, t)
}

// NewTreatmenttypeController creates and registers handles for the TreatmentType controller
func NewTreatmenttypeController(router gin.IRouter, client *ent.Client) *TreatmenttypeController {
	tc := &TreatmenttypeController{
		client: client,
		router: router,
	}

	tc.register()

	return tc

}

func (ctl *TreatmenttypeController) register() {
	TreatmentTypes := ctl.router.Group("/TreatmentTypes")

	TreatmentTypes.GET("", ctl.ListTreatmentType)

	// CRUD
	TreatmentTypes.POST("", ctl.CreateTreatmentType)
	TreatmentTypes.GET(":id", ctl.GetTreatmentType)
	TreatmentTypes.PUT(":id", ctl.UpdateTreatmentType)
	TreatmentTypes.DELETE(":id", ctl.DeleteTreatmentType)
}
