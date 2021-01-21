package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/tanapon395/playlist-video/ent/customer"
	"github.com/tanapon395/playlist-video/ent/fix"
	"github.com/tanapon395/playlist-video/ent/fixbrand"
	"github.com/tanapon395/playlist-video/ent/fixcomtype"
	"github.com/tanapon395/playlist-video/ent/personal"

	"github.com/gin-gonic/gin"
	"github.com/tanapon395/playlist-video/ent"
)

// FixController defines the struct for the fix controller
type FixController struct {
	client *ent.Client
	router gin.IRouter
}

// Fix defines the struct for the fix controller
type Fix struct {
	Productnumber string
	Problemtype   string
	Date          string
	Queue         string
	Fixbrand      int
	Personal      int
	Customer      int
	Fixcomtype    int
}

// CreateFix handles POST requests for adding fix entities
// @Summary Create fix
// @Description Create fix
// @ID create-fix
// @Accept   json
// @Produce  json
// @Param fix body Fix true "Fix entity"
// @Success 200 {object} Fix
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /fixs [post]
func (ctl *FixController) CreateFix(c *gin.Context) {
	obj := Fix{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "fix binding failed",
		})
		return
	}

	fb, err := ctl.client.Fixbrand.
		Query().
		Where(fixbrand.IDEQ(int(obj.Fixbrand))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "fixbrand not found",
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

	cm, err := ctl.client.Customer.
		Query().
		Where(customer.IDEQ(int(obj.Customer))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "customer not found",
		})
		return
	}

	ft, err := ctl.client.Fixcomtype.
		Query().
		Where(fixcomtype.IDEQ(int(obj.Fixcomtype))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "fixcomtype not found",
		})
		return
	}
	t1 := time.Now()
	t2 := t1.Format("2006-01-02T15:04:05Z07:00")
	time1, err := time.Parse(time.RFC3339, t2)
	//date, err := time.Parse(time.RFC3339, obj.Date)
	f, err := ctl.client.Fix.
		Create().
		SetProductnumber(obj.Productnumber).
		SetProblemtype(obj.Problemtype).
		SetDate(time1).
		SetQueue(obj.Queue).
		SetFixbrand(fb).
		SetPersonal(p).
		SetCustomer(cm).
		SetFixcomtype(ft).
		Save(context.Background())
		if err != nil {
			c.JSON(400, gin.H{
				"status": false,
				"error":  err,
			})
			return
		}
	
		c.JSON(200, gin.H{
			"status": true,
			"data":   f,
		})
}

// GetFix handles GET requests to retrieve a fix entity
// @Summary Get a fix entity by ID
// @Description get fix by ID
// @ID get-fix
// @Produce  json
// @Param id path int true "Fix ID"
// @Success 200 {object} ent.Fix
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /fixs/{id} [get]
func (ctl *FixController) GetFix(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	f, err := ctl.client.Fix.
		Query().
		Where(fix.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, f)
}

// ListFix handles request to get a list of fix entities
// @Summary List fix entities
// @Description list fix entities
// @ID list-fix
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Fix
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /fixs [get]
func (ctl *FixController) ListFix(c *gin.Context) {
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

	fixs, err := ctl.client.Fix.
		Query().
		WithFixbrand().
		WithPersonal().
		WithCustomer().
		WithFixcomtype().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, fixs)
}

// DeleteFix handles DELETE requests to delete a fix entity
// @Summary Delete a fix entity by ID
// @Description get fix by ID
// @ID delete-fix
// @Produce  json
// @Param id path int true "Fix ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /fixs/{id} [delete]
func (ctl *FixController) DeleteFix(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Fix.
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

// UpdateFix handles PUT requests to update a fix entity
// @Summary Update a fix entity by ID
// @Description update fix by ID
// @ID update-fix
// @Accept   json
// @Produce  json
// @Param id path int true "Fix ID"
// @Param fix body ent.Fix true "Fix entity"
// @Success 200 {object} ent.Fix
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /fixs/{id} [put]
func (ctl *FixController) UpdateFix(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Fix{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "fix binding failed",
		})
		return
	}
	obj.ID = int(id)
	f, err := ctl.client.Fix.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, f)
}

// NewFixController creates and registers handles for the Fix controller
func NewFixController(router gin.IRouter, client *ent.Client) *FixController {
	fc := &FixController{
		client: client,
		router: router,
	}

	fc.register()

	return fc

}

func (ctl *FixController) register() {
	fixs := ctl.router.Group("/fixs")

	// CRUD
	fixs.POST("", ctl.CreateFix)
	fixs.GET(":id", ctl.GetFix)
	fixs.GET("", ctl.ListFix)
	fixs.DELETE(":id", ctl.DeleteFix)
	fixs.PUT(":id", ctl.UpdateFix)

}
