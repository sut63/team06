package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team06/app/ent"
	"github.com/team06/app/ent/proceduretype"
)

// ProceduretypeController defines the struct for the proceduretype controller
type ProceduretypeController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateProceduretype handles POST requests for adding proceduretype entities
// @Summary Create proceduretype
// @Description Create proceduretype
// @ID create-proceduretype
// @Accept   json
// @Produce  json
// @Param pharmacist body ent.proceduretype true "proceduretype entity"
// @Success 200 {object} ent.proceduretype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /proceduretype [post]
func (ctl *ProceduretypeController) CreateProceduretype(c *gin.Context) {
	obj := ent.ProcedureType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Proceduretype binding failed",
		})
		return
	}

	procedure, err := ctl.client.ProcedureType.
		Create().
		SetProcedureType(obj.ProcedureType).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, procedure)
}

// GetProceduretype handles GET requests to retrieve a proceduretype entity
// @Summary Get a proceduretype entity by ID
// @Description get proceduretype by ID
// @ID get-proceduretype
// @Produce  json
// @Param id path int true "proceduretype ID"
// @Success 200 {object} ent.Proceduretype
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /proceduretype/{id} [get]
func (ctl *ProceduretypeController) GetProceduretype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	procedure, err := ctl.client.ProcedureType.
		Query().
		Where(proceduretype.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, procedure)
}

// ListProceduretype handles request to get a list of proceduretype entities
// @Summary List proceduretype entities
// @Description list proceduretype entities
// @ID list-proceduretype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Proceduretype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /proceduretype [get]
func (ctl *ProceduretypeController) ListProceduretype(c *gin.Context) {
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

	proceduretype, err := ctl.client.ProcedureType.
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

	c.JSON(200, proceduretype)
}

// DeleteProceduretype handles DELETE requests to delete a proceduretype entity
// @Summary Delete a proceduretype entity by ID
// @Description get proceduretype by ID
// @ID delete-proceduretype
// @Produce  json
// @Param id path int true "proceduretype ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /proceduretype/{id} [delete]
func (ctl *ProceduretypeController) DeleteProceduretype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.ProcedureType.
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

// UpdateProceduretype handles PUT requests to update a proceduretype entity
// @Summary Update a proceduretype entity by ID
// @Description update proceduretype by ID
// @ID update-proceduretype
// @Accept   json
// @Produce  json
// @Param id path int true "proceduretype ID"
// @Param pharmacist body ent.Proceduretype true "proceduretype entity"
// @Success 200 {object} ent.Proceduretype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /proceduretype/{id} [put]
func (ctl *ProceduretypeController) UpdateProceduretype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.ProcedureType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Proceduretype binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	procedure, err := ctl.client.ProcedureType.
		UpdateOneID(int(id)).
		SetProcedureType(obj.ProcedureType).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "update failed",
		})
		return
	}

	c.JSON(200, procedure)
}

// NewProceduretypeController creates and registers handles for the proceduretype controller
func NewProceduretypeController(router gin.IRouter, client *ent.Client) *ProceduretypeController {
	prc := &ProceduretypeController{
		client: client,
		router: router,
	}

	prc.register()

	return prc

}

func (ctl *ProceduretypeController) register() {
	proceduretypes := ctl.router.Group("/proceduretype")

	proceduretypes.GET("", ctl.ListProceduretype)

	// CRUD
	proceduretypes.POST("", ctl.CreateProceduretype)
	proceduretypes.GET(":id", ctl.GetProceduretype)
	proceduretypes.PUT(":id", ctl.UpdateProceduretype)
	proceduretypes.DELETE(":id", ctl.DeleteProceduretype)
}
