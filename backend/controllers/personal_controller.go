package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tanapon395/playlist-video/ent"
	"github.com/tanapon395/playlist-video/ent/department"
	"github.com/tanapon395/playlist-video/ent/gender"
	"github.com/tanapon395/playlist-video/ent/title"
)

type PersonalController struct {
	client *ent.Client
	router gin.IRouter
}

type Personal struct {
	Name       string
	Email      string
	Password   string
	Title      int
	Department int
	Gender     int
}

// CreatePersonal handles POST requests for adding personal entities
// @Summary Create personal
// @Description Create personal
// @ID create-personal
// @Accept   json
// @Produce  json
// @Param personal body Personal true "Personal entity"
// @Success 200 {object} Personal
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /personals [post]
func (ctl *PersonalController) CreatePersonal(c *gin.Context) {
	obj := Personal{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "personal binding failed",
		})
		return
	}

	nt, err := ctl.client.Title.
		Query().
		Where(title.IDEQ(int(obj.Title))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "title not found",
		})
		return
	}

	d, err := ctl.client.Department.
		Query().
		Where(department.IDEQ(int(obj.Department))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "department not found",
		})
		return
	}

	g, err := ctl.client.Gender.
		Query().
		Where(gender.IDEQ(int(obj.Gender))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "gender not found",
		})
		return
	}

	p, err := ctl.client.Personal.
		Create().
		SetName(obj.Name).
		SetEmail(obj.Email).
		SetPassword(obj.Password).
		SetTitle(nt).
		SetDepartment(d).
		SetGender(g).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, p)
}

// DeletePersonal handles DELETE requests to delete a personal entity
// @Summary Delete a personal entity by ID
// @Description get personal by ID
// @ID delete-personal
// @Produce  json
// @Param id path int true "Personal ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /personals/{id} [delete]
func (ctl *PersonalController) DeletePersonal(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Personal.
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

// ListPersonal handles request to get a list of personal entities
// @Summary List personal entities
// @Description list personal entities
// @ID list-personal
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Personal
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /personals [get]
func (ctl *PersonalController) ListPersonal(c *gin.Context) {
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

	personals, err := ctl.client.Personal.
		Query().
		WithTitle().
		WithDepartment().
		WithGender().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, personals)
}

// NewPersonalController creates and registers handles for the personal controller
func NewPersonalController(router gin.IRouter, client *ent.Client) *PersonalController {
	pvc := &PersonalController{
		client: client,
		router: router,
	}

	pvc.register()

	return pvc

}

func (ctl *PersonalController) register() {
	personals := ctl.router.Group("/personals")

	personals.POST("", ctl.CreatePersonal)
	personals.GET("", ctl.ListPersonal)
	personals.DELETE(":id", ctl.DeletePersonal)

}
