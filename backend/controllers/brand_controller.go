package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tanapon395/playlist-video/ent"
	"github.com/tanapon395/playlist-video/ent/brand"
)

type BrandController struct {
	client *ent.Client
	router gin.IRouter
}

type Brand struct {
	Brandname string
}

// CreateBrand handles POST requests for adding brand entities
// @Summary Create brand
// @Description Create brand
// @ID create-brand
// @Accept   json
// @Produce  json
// @Param brand body ent.Brand true "Brand entity"
// @Success 200 {object} ent.Brand
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /brands [post]
func (ctl *BrandController) CreateBrand(c *gin.Context) {
	obj := Brand{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "brand binding failed",
		})
		return
	}

	b, err := ctl.client.Brand.
		Create().
		SetBrandname(obj.Brandname).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, b)
}

// GetBrand handles GET requests to retrieve a brand entity
// @Summary Get a brand entity by ID
// @Description get brand by ID
// @ID get-brand
// @Produce  json
// @Param id path int true "Brand ID"
// @Success 200 {object} ent.Brand
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /brands/{id} [get]
func (ctl *BrandController) GetBrand(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	b, err := ctl.client.Brand.
		Query().
		Where(brand.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, b)
}

// ListBrand handles request to get a list of brand entities
// @Summary List brand entities
// @Description list brand entities
// @ID list-brand
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Brand
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /brands [get]
func (ctl *BrandController) ListBrand(c *gin.Context) {
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

	brands, err := ctl.client.Brand.
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

	c.JSON(200, brands)
}

// DeleteBrand handles DELETE requests to delete a brand entity
// @Summary Delete a brand entity by ID
// @Description get brand by ID
// @ID delete-brand
// @Produce  json
// @Param id path int true "Brand ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /brand/{id} [delete]
func (ctl *BrandController) DeleteBrand(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Brand.
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

// NewBrandController creates and registers handles for the brand controller
func NewBrandController(router gin.IRouter, client *ent.Client) *BrandController {
	bc := &BrandController{
		client: client,
		router: router,
	}

	bc.register()

	return bc

}

func (ctl *BrandController) register() {
	brands := ctl.router.Group("/brands")

	brands.POST("", ctl.CreateBrand)
	brands.GET(":id", ctl.GetBrand)
	brands.GET("", ctl.ListBrand)
	brands.DELETE("", ctl.DeleteBrand)

}