package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/tanapon395/playlist-video/ent"
	"github.com/tanapon395/playlist-video/ent/brand"
	"github.com/tanapon395/playlist-video/ent/typeproduct"
	"github.com/tanapon395/playlist-video/ent/personal"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	client *ent.Client
	router gin.IRouter
}

type Product struct {
	Productname               string
	Numberofproduct		      string
	Price		      		  string
	Brand			          int
	Typeproduct 			  int
	Personal     	 		  int
}

// CreateProduct handles POST requests for adding product entities
// @Summary Create product
// @Description Create product
// @ID create-product
// @Accept   json
// @Produce  json
// @Param product body Product true "Product entity"
// @Success 200 {object} Product
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /products [post]
func (ctl *ProductController) CreateProduct(c *gin.Context) {
	obj := Product{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "product binding failed",
		})
		return
	}

	b, err := ctl.client.Brand.
		Query().
		Where(brand.IDEQ(int(obj.Brand))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "brand not found",
		})
		return
	}

	t, err := ctl.client.Typeproduct.
		Query().
		Where(typeproduct.IDEQ(int(obj.Typeproduct))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "typeproduct not found",
		})
		return
	}

	p, err := ctl.client.Personal.
		Query().
		Where(personal.IDEQ(int(obj.Personal))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "personal not found",
		})
		return
	}

	pr, err := ctl.client.Product.
		Create().
		SetProductname(obj.Productname).
		SetNumberofproduct(obj.Numberofproduct).
		SetPrice(obj.Price).
		SetBrand(b).
		SetTypeproduct(t).
		SetPersonal(p).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, pr)
}

// DeleteProduct handles DELETE requests to delete a product entity
// @Summary Delete a product entity by ID
// @Description get product by ID
// @ID delete-product
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /products/{id} [delete]
func (ctl *ProductController) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Product.
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

// ListProduct handles request to get a list of product entities
// @Summary List product entities
// @Description list product entities
// @ID list-product
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Product
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /products [get]
func (ctl *ProductController) ListProduct(c *gin.Context) {
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

	products, err := ctl.client.Product.
		Query().
		WithBrand().
		WithTypeproduct().
		WithPersonal().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, products)
}

// NewProductController creates and registers handles for the product controller
func NewProductController(router gin.IRouter, client *ent.Client) *ProductController {
	pic := &ProductController{
		client: client,
		router: router,
	}

	pic.register()

	return pic

}

func (ctl *ProductController) register() {
	products := ctl.router.Group("/products")

	products.POST("", ctl.CreateProduct)
	products.GET("", ctl.ListProduct)
	products.DELETE(":id", ctl.DeleteProduct)

}
