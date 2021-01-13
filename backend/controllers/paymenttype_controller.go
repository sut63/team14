package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tanapon395/playlist-video/ent"
	"github.com/tanapon395/playlist-video/ent/paymenttype"
)

// PaymentTypeController defines the struct for the user controller
type PaymentTypeController struct {
	client *ent.Client
	router gin.IRouter
}

type PaymentType struct {
	Typename string
}

// CreatePaymentType handles POST requests for adding paymenttype entities
// @Summary Create paymenttype
// @Description Create paymenttype
// @ID create-paymenttype
// @Accept   json
// @Produce  json
// @Param paymenttype body ent.PaymentType true "PaymentType entity"
// @Success 200 {object} ent.PaymentType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paymenttypes [post]
func (ctl *PaymentTypeController) CreatePaymentType(c *gin.Context) {
	obj := ent.PaymentType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "paymenttype binding failed",
		})
		return
	}

	pt, err := ctl.client.PaymentType.
		Create().
		SetTypename(obj.Typename).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, pt)
}

// GetPaymentType handles GET requests to retrieve a paymenttype entity
// @Summary Get a paymenttype entity by ID
// @Description get paymenttype by ID
// @ID get-paymenttype
// @Produce  json
// @Param id path int true "PaymentType ID"
// @Success 200 {object} ent.PaymentType
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paymenttypes/{id} [get]
func (ctl *PaymentTypeController) GetPaymentType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	pt, err := ctl.client.PaymentType.
		Query().
		Where(paymenttype.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, pt)
}

// ListPaymentType handles request to get a list of paymenttype entities
// @Summary List paymenttype entities
// @Description list paymenttype entities
// @ID list-paymenttype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.PaymentType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paymenttypes [get]
func (ctl *PaymentTypeController) ListPaymentType(c *gin.Context) {
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

	paymenttypes, err := ctl.client.PaymentType.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, paymenttypes)
}

// DeletePaymentType handles DELETE requests to delete a paymenttype entity
// @Summary Delete a paymenttype entity by ID
// @Description get paymenttype by ID
// @ID delete-paymenttype
// @Produce  json
// @Param id path int true "PaymentType ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paymenttypes/{id} [delete]
func (ctl *PaymentTypeController) DeletePaymentType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.PaymentType.
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

// UpdatePaymentType handles PUT requests to update a paymenttype entity
// @Summary Update a paymenttype entity by ID
// @Description update paymenttype by ID
// @ID update-paymenttype
// @Accept   json
// @Produce  json
// @Param id path int true "PaymentType ID"
// @Param paymenttype body ent.PaymentType true "PaymentType entity"
// @Success 200 {object} ent.PaymentType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paymenttypes/{id} [put]
func (ctl *PaymentTypeController) UpdatePaymentType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.PaymentType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "paymenttype binding failed",
		})
		return
	}
	obj.ID = int(id)
	pt, err := ctl.client.PaymentType.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, pt)
}

// NewPaymentTypeController creates and registers handles for the paymenttype controller
func NewPaymentTypeController(router gin.IRouter, client *ent.Client) *PaymentTypeController {
	ptc := &PaymentTypeController{
		client: client,
		router: router,
	}
	ptc.register()
	return ptc
}

// InitPaymentTypeController registers routes to the main engine
func (ctl *PaymentTypeController) register() {
	paymenttypes := ctl.router.Group("/paymenttypes")

	paymenttypes.GET("", ctl.ListPaymentType)

	// CRUD
	paymenttypes.POST("", ctl.CreatePaymentType)
	paymenttypes.GET(":id", ctl.GetPaymentType)
	paymenttypes.PUT(":id", ctl.UpdatePaymentType)
	paymenttypes.DELETE(":id", ctl.DeletePaymentType)
}
