package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tanapon395/playlist-video/ent"
	"github.com/tanapon395/playlist-video/ent/fixbrand"
)

type FixbrandController struct {
	client *ent.Client
	router gin.IRouter
}

type Fixbrand struct {
	Fixbrandname string
}

// CreateFixbrand handles POST requests for adding fixbrand entities
// @Summary Create fixbrand
// @Description Create fixbrand
// @ID create-fixbrand
// @Accept   json
// @Produce  json
// @Param fixbrand body ent.Fixbrand true "Fixbrand entity"
// @Success 200 {object} ent.Fixbrand
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /fixbrands [post]
func (ctl *FixbrandController) CreateFixbrand(c *gin.Context) {
	obj := Fixbrand{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "fixbrand binding failed",
		})
		return
	}

	ft, err := ctl.client.Fixbrand.
		Create().
		SetFixbrandname(obj.Fixbrandname).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, ft)
}

// GetFixbrand handles GET requests to retrieve a fixbrand entity
// @Summary Get a fixbrand entity by ID
// @Description get fixbrand by ID
// @ID get-fixbrand
// @Produce  json
// @Param id path int true "Fixbrand ID"
// @Success 200 {object} ent.Fixbrand
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /fixbrands/{id} [get]
func (ctl *FixbrandController) GetFixbrand(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	fb, err := ctl.client.Fixbrand.
		Query().
		Where(fixbrand.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, fb)
}

// ListFixbrand handles request to get a list of fixbrand entities
// @Summary List fixbrand entities
// @Description list fixbrand entities
// @ID list-fixbrand
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Fixbrand
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /fixbrands [get]
func (ctl *FixbrandController) ListFixbrand(c *gin.Context) {
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

	fixbrands, err := ctl.client.Fixbrand.
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

	c.JSON(200, fixbrands)
}

// DeleteFixbrand handles DELETE requests to delete a fixbrand entity
// @Summary Delete a fixbrand entity by ID
// @Description get fixbrand by ID
// @ID delete-fixbrand
// @Produce  json
// @Param id path int true "Fixbrand ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /fixbrands/{id} [delete]
func (ctl *FixbrandController) DeleteFixbrand(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Fixbrand.
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

// NewFixbrandController creates and registers handles for the fixbrand controller
func NewFixbrandController(router gin.IRouter, client *ent.Client) *FixbrandController {
	ftc := &FixbrandController{
		client: client,
		router: router,
	}

	ftc.register()

	return ftc

}

func (ctl *FixbrandController) register() {
	fixbrands := ctl.router.Group("/fixbrands")

	fixbrands.POST("", ctl.CreateFixbrand)
	fixbrands.GET(":id", ctl.GetFixbrand)
	fixbrands.GET("", ctl.ListFixbrand)
	fixbrands.DELETE("", ctl.DeleteFixbrand)

}
