package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tanapon395/playlist-video/ent"
	"github.com/tanapon395/playlist-video/ent/adminrepair"
	"github.com/tanapon395/playlist-video/ent/fix"
	"github.com/tanapon395/playlist-video/ent/personal"
	"github.com/tanapon395/playlist-video/ent/product"
)

// AdminrepairController defines the struct for the adminrepair controller
type AdminrepairController struct {
	client *ent.Client
	router gin.IRouter
}

type Adminrepair struct {
	Personal          int
	Fix               int
	Product           int
	Equipmentdamate   string
	Numberrepair      string
	Repairinformation string
}

// CreateAdminrepair handles POST requests for adding adminrepair entities
// @Summary Create adminrepair
// @Description Create adminrepair
// @ID create-adminrepair
// @Accept   json
// @Produce  json
// @Param adminrepair body Adminrepair true "Adminrepair entity"
// @Success 200 {object} ent.Adminrepair
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /adminrepairs [post]
func (ctl *AdminrepairController) CreateAdminrepair(c *gin.Context) {
	obj := Adminrepair{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "adminrepair binding failed",
		})
		return
	}

	p, err := ctl.client.Personal.
		Query().
		Where(personal.IDEQ(obj.Personal)).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "personal not found",
		})
		return
	}

	f, err := ctl.client.Fix.
		Query().
		Where(fix.IDEQ(int(obj.Fix))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "fix not found",
		})
		return
	}

	pd, err := ctl.client.Product.
		Query().
		Where(product.IDEQ(int(obj.Product))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "product not found",
		})
		return
	}

	a, err := ctl.client.Adminrepair.
		Create().
		SetNumberrepair(obj.Numberrepair).
		SetEquipmentdamate(obj.Equipmentdamate).
		SetRepairinformation(obj.Repairinformation).
		SetAdminrepairPersonal(p).
		SetAdminrepairFix(f).
		SetAdminrepairProduct(pd).
		Save(context.Background())

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"error":  a,
	})
}

// GetAdminrepair handles GET requests to retrieve a adminrepair entity
// @Summary Get a adminrepair entity by ID
// @Description get adminrepair by ID
// @ID get-adminrepair
// @Produce  json
// @Param id path int true "Adminrepair ID"
// @Success 200 {object} ent.Adminrepair
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /adminrepairs/{id} [get]
func (ctl *AdminrepairController) GetAdminrepair(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	a, err := ctl.client.Adminrepair.
		Query().
		WithAdminrepairPersonal().
		WithAdminrepairFix().
		WithAdminrepairProduct().
		Where(adminrepair.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, a)
}

// ListAdminrepair handles request to get a list of adminrepair entities
// @Summary List adminrepair entities
// @Description list adminrepair entities
// @ID list-adminrepair
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Adminrepair
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /adminrepairs [get]
func (ctl *AdminrepairController) ListAdminrepair(c *gin.Context) {
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

	adminrepairs, err := ctl.client.Adminrepair.
		Query().
		WithAdminrepairPersonal().
		WithAdminrepairFix().
		WithAdminrepairProduct().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, adminrepairs)
}

// DeleteAdminrepair handles DELETE requests to delete a adminrepair entity
// @Summary Delete a adminrepair entity by ID
// @Description get adminrepair by ID
// @ID delete-adminrepair
// @Produce  json
// @Param id path int true "Adminrepair ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /adminrepairs/{id} [delete]
func (ctl *AdminrepairController) DeleteAdminrepair(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Adminrepair.
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

// UpdateAdminrepair handles PUT requests to update a adminrepair entity
// @Summary Update a adminrepair entity by ID
// @Description update adminrepair by ID
// @ID update-adminrepair
// @Accept   json
// @Produce  json
// @Param id path int true "Adminrepair ID"
// @Param adminrepair body ent.Adminrepair true "Adminrepair entity"
// @Success 200 {object} ent.Adminrepair
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /adminrepairs/{id} [put]
func (ctl *AdminrepairController) UpdateAdminrepair(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Adminrepair{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "adminrepair binding failed",
		})
		return
	}
	obj.ID = int(id)
	a, err := ctl.client.Adminrepair.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, a)
}

// NewAdminrepairController creates and registers handles for the adminrepair controller
func NewAdminrepairController(router gin.IRouter, client *ent.Client) *AdminrepairController {
	ac := &AdminrepairController{
		client: client,
		router: router,
	}
	ac.register()
	return ac
}

// InitAdminrepairController registers routes to the main engine
func (ctl *AdminrepairController) register() {
	adminrepairs := ctl.router.Group("/adminrepairs")

	adminrepairs.GET("", ctl.ListAdminrepair)
	adminrepairs.POST("", ctl.CreateAdminrepair)
	adminrepairs.GET(":id", ctl.GetAdminrepair)
	adminrepairs.PUT(":id", ctl.UpdateAdminrepair)
	adminrepairs.DELETE(":id", ctl.DeleteAdminrepair)
}
