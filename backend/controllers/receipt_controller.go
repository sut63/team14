package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tanapon395/playlist-video/ent"
	"github.com/tanapon395/playlist-video/ent/adminrepair"
	"github.com/tanapon395/playlist-video/ent/paymenttype"
	"github.com/tanapon395/playlist-video/ent/personal"
	"github.com/tanapon395/playlist-video/ent/receipt"
)

// ReceiptController defines the struct for the receipt controller
type ReceiptController struct {
	client *ent.Client
	router gin.IRouter
}

type Receipt struct {
	Cusidentification string
	Customername      string
	Phonenumber       string
	Added             string
	Personal          int
	PaymentType       int
	Adminrepair       int
}

// CreateReceipt handles POST requests for adding receipt entities
// @Summary Create receipt
// @Description Create receipt
// @ID create-receipt
// @Accept   json
// @Produce  json
// @Param receipt body ent.Receipt true "Receipt entity"
// @Success 200 {object} ent.Receipt
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /receipts [post]
func (ctl *ReceiptController) CreateReceipt(c *gin.Context) {
	obj := Receipt{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "receipt binding failed",
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
	pt, err := ctl.client.PaymentType.
		Query().
		Where(paymenttype.IDEQ(int(obj.PaymentType))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "paymenttype not found",
		})
		return
	}
	a, err := ctl.client.Adminrepair.
		Query().
		Where(adminrepair.IDEQ(int(obj.Adminrepair))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "adminrepair not found",
		})
		return
	}

	times, err := time.Parse(time.RFC3339, obj.Added)
	r, err := ctl.client.Receipt.
		Create().
		SetCusidentification(obj.Cusidentification).
		SetCustomername(obj.Customername).
		SetPhonenumber(obj.Phonenumber).
		SetAddedTime(times).
		SetPersonal(p).
		SetAdminrepair(a).
		SetPaymenttype(pt).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, r)
}

// GetReceipt handles GET requests to retrieve a receipt entity
// @Summary Get a receipt entity by ID
// @Description get receipt by ID
// @ID get-receipt
// @Produce  json
// @Param id path int true "Receipt ID"
// @Success 200 {object} ent.Receipt
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /receipts/{id} [get]
func (ctl *ReceiptController) GetReceipt(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	r, err := ctl.client.Receipt.
		Query().
		Where(receipt.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, r)
}

// ListReceipt handles request to get a list of receipt entities
// @Summary List receipt entities
// @Description list receipt entities
// @ID list-receipt
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Receipt
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /receipts [get]
func (ctl *ReceiptController) ListReceipt(c *gin.Context) {
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

	receipts, err := ctl.client.Receipt.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, receipts)
}

// DeleteReceipt handles DELETE requests to delete a receipt entity
// @Summary Delete a receipt entity by ID
// @Description get receipt by ID
// @ID delete-receipt
// @Produce  json
// @Param id path int true "Receipt ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /receipts/{id} [delete]
func (ctl *ReceiptController) DeleteReceipt(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Receipt.
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

// UpdateReceipt handles PUT requests to update a receipt entity
// @Summary Update a receipt entity by ID
// @Description update receipt by ID
// @ID update-receipt
// @Accept   json
// @Produce  json
// @Param id path int true "receipt ID"
// @Param receipt body ent.Receipt true "Receipt entity"
// @Success 200 {object} ent.Receipt
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /receipts/{id} [put]
func (ctl *ReceiptController) UpdateReceipt(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Receipt{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "receipt binding failed",
		})
		return
	}
	obj.ID = int(id)
	r, err := ctl.client.Receipt.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, r)
}

// NewReceiptController creates and registers handles for the receipt controller
func NewReceiptController(router gin.IRouter, client *ent.Client) *ReceiptController {
	rc := &ReceiptController{
		client: client,
		router: router,
	}
	rc.register()
	return rc
}

// InitReceiptController registers routes to the main engine
func (ctl *ReceiptController) register() {
	receipts := ctl.router.Group("/receipts")

	receipts.GET("", ctl.ListReceipt)

	// CRUD
	receipts.POST("", ctl.CreateReceipt)
	receipts.GET(":id", ctl.GetReceipt)
	receipts.PUT(":id", ctl.UpdateReceipt)
	receipts.DELETE(":id", ctl.DeleteReceipt)
}