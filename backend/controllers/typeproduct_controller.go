package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tanapon395/playlist-video/ent"
	"github.com/tanapon395/playlist-video/ent/typeproduct"
)

type TypeproductController struct {
	client *ent.Client
	router gin.IRouter
}

type Typeproduct struct {
	Typeproductname	 string
}

// CreateTypeproduct handles POST requests for adding typeproduct entities
// @Summary Create typeproduct
// @Description Create typeproduct
// @ID create-typeproduct
// @Accept   json
// @Produce  json
// @Param typeproduct body ent.Typeproduct true "typeproduct entity"
// @Success 200 {object} ent.Typeproduct
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /typeproducts [post]
func (ctl *TypeproductController) CreateTypeproduct(c *gin.Context) {
	obj := Typeproduct{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "typeproduct binding failed",
		})
		return
	}

	t, err := ctl.client.Typeproduct.
		Create().
		SetTypeproductname(obj.Typeproductname).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, t)
}

// GetTypeproduct handles GET requests to retrieve a typeproduct entity
// @Summary Get a typeproduct entity by ID
// @Description get typeproduct by ID
// @ID get-typeproduct
// @Produce  json
// @Param id path int true "Typeproduct ID"
// @Success 200 {object} ent.Typeproduct
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /typeproducts/{id} [get]
func (ctl *TypeproductController) GetTypeproduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	t, err := ctl.client.Typeproduct.
		Query().
		Where(typeproduct.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, t)
}

// ListTypeproduct handles request to get a list of typeproduct entities
// @Summary List typeproduct entities
// @Description list typeproduct entities
// @ID list-typeproduct
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Typeproduct
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /typeproducts [get]
func (ctl *TypeproductController) ListTypeproduct(c *gin.Context) {
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

	typeproducts, err := ctl.client.Typeproduct.
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

	c.JSON(200, typeproducts)
}

// DeleteTypeproduct handles DELETE requests to delete a typeproduct entity
// @Summary Delete a typeproduct entity by ID
// @Description get typeproduct by ID
// @ID delete-typeproduct
// @Produce  json
// @Param id path int true "Typeproduct ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /typeproduct/{id} [delete]
func (ctl *TypeproductController) DeleteTypeproduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Typeproduct.
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

// NewTypeproductController creates and registers handles for the typeproduct controller
func NewTypeproductController(router gin.IRouter, client *ent.Client) *TypeproductController {
	tc := &TypeproductController{
		client: client,
		router: router,
	}

	tc.register()

	return tc

}

func (ctl *TypeproductController) register() {
	typeproducts := ctl.router.Group("/typeproducts")

	typeproducts.POST("", ctl.CreateTypeproduct)
	typeproducts.GET(":id", ctl.GetTypeproduct)
	typeproducts.GET("", ctl.ListTypeproduct)
	typeproducts.DELETE("", ctl.DeleteTypeproduct)

}