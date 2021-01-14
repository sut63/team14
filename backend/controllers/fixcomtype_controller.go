package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tanapon395/playlist-video/ent"
	"github.com/tanapon395/playlist-video/ent/fixcomtype"
)

type FixcomtypeController struct {
	client *ent.Client
	router gin.IRouter
}

type Fixcomtype struct {
	Fixcomtypename string
}

// CreateFixcomtype handles POST requests for adding fixcomtype entities
// @Summary Create fixcomtype
// @Description Create fixcomtype
// @ID create-fixcomtype
// @Accept   json
// @Produce  json
// @Param fixcomtype body Fixcomtype true "Fixcomtype entity"
// @Success 200 {object} Fixcomtype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /fixcomtypes [post]
func (ctl *FixcomtypeController) CreateFixcomtype(c *gin.Context) {
	obj := Fixcomtype{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Fixcomtype binding failed",
		})
		return
	}

	ft, err := ctl.client.Fixcomtype.
		Create().
		SetFixcomtypename(obj.Fixcomtypename).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, ft)
}

// GetFixcomtype handles GET requests to retrieve a fixcomtype entity
// @Summary Get a fixcomtype entity by ID
// @Description get fixcomtype by ID
// @ID get-fixcomtype
// @Produce  json
// @Param id path int true "Fixcomtype ID"
// @Success 200 {object} ent.Fixcomtype
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /fixcomtypes/{id} [get]
func (ctl *FixcomtypeController) GetFixcomtype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	ft, err := ctl.client.Fixcomtype.
		Query().
		Where(fixcomtype.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, ft)
}

// ListFixcomtype handles request to get a list of fixcomtype entities
// @Summary List fixcomtype entities
// @Description list fixcomtype entities
// @ID list-fixcomtype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Fixcomtype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /fixcomtypes [get]
func (ctl *FixcomtypeController) ListFixcomtype(c *gin.Context) {
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

	fixcomtypes, err := ctl.client.Fixcomtype.
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

	c.JSON(200, fixcomtypes)
}

// DeleteFixcomtype handles DELETE requests to delete a fixcomtype entity
// @Summary Delete a fixcomtype entity by ID
// @Description get fixcomtype by ID
// @ID delete-fixcomtype
// @Produce  json
// @Param id path int true "Fixcomtype ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /fixcomtypes/{id} [delete]
func (ctl *FixcomtypeController) DeleteFixcomtype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Fixcomtype.
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

// UpdateFixcomtype handles PUT requests to update a fixcomtype entity
// @Summary Update a fixcomtype entity by ID
// @Description update fixcomtype by ID
// @ID update-fixcomtype
// @Accept   json
// @Produce  json
// @Param id path int true "Fixcomtype ID"
// @Param fixcomtype body ent.Fixcomtype true "Fixcomtype entity"
// @Success 200 {object} ent.Fixcomtype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /fixcomtypes/{id} [put]
func (ctl *FixcomtypeController) UpdateFixcomtype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Fixcomtype{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "fixcomtype binding failed",
		})
		return
	}
	obj.ID = int(id)
	ft, err := ctl.client.Fixcomtype.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, ft)
}

// NewFixcomtypeController creates and registers handles for the fixcomtype controller
func NewFixcomtypeController(router gin.IRouter, client *ent.Client) *FixcomtypeController {
	ftc := &FixcomtypeController{
		client: client,
		router: router,
	}

	ftc.register()

	return ftc

}

// InitFixcomtypeController registers routes to the main engine
func (ctl *FixcomtypeController) register() {
	fixcomtypes := ctl.router.Group("/fixcomtypes")

	fixcomtypes.GET("", ctl.ListFixcomtype)

	// CRUD
	fixcomtypes.POST("", ctl.CreateFixcomtype)
	fixcomtypes.GET(":id", ctl.GetFixcomtype)
	fixcomtypes.PUT(":id", ctl.UpdateFixcomtype)
	fixcomtypes.DELETE(":id", ctl.DeleteFixcomtype)
}
