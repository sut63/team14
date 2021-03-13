package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tanapon395/playlist-video/ent"
	"github.com/tanapon395/playlist-video/ent/department"
	"github.com/tanapon395/playlist-video/ent/gender"
	"github.com/tanapon395/playlist-video/ent/personal"
	"github.com/tanapon395/playlist-video/ent/title"
)

type PersonalController struct {
	client *ent.Client
	router gin.IRouter
}

type Personal struct {
	Personalname string
	Email        string
	Password     string
	Title        int
	Department   int
	Gender       int
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
		SetPersonalname(obj.Personalname).
		SetEmail(obj.Email).
		SetPassword(obj.Password).
		SetTitle(nt).
		SetDepartment(d).
		SetGender(g).
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
		"data":   p,
	})
}

// GetPersonal handles GET requests to retrieve a personal entity
// @Summary Get a personal entity by ID
// @Description get personal by ID
// @ID get-personal
// @Produce  json
// @Param id path int true "Personal ID"
// @Success 200 {object} ent.Personal
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /personals/{id} [get]
func (ctl *PersonalController) GetPersonal(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	p, err := ctl.client.Personal.
		Query().
		Where(personal.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, p)
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

// UpdatePersonal handles PUT requests to update a personal entity
// @Summary Update a personal entity by ID
// @Description update personal by ID
// @ID update-personal
// @Accept   json
// @Produce  json
// @Param id path int true "Personal ID"
// @Param personal body ent.Personal true "Personal entity"
// @Success 200 {object} ent.Personal
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /personals/{id} [put]
func (ctl *PersonalController) UpdatePersonal(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Personal{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "personal binding failed",
		})
		return
	}
	obj.ID = int(id)
	p, err := ctl.client.Personal.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, p)
}

// GetPersonalSearch handles GET requests to retrieve a Personal entity
// @Summary Get a personal entity by Search
// @Description get personal by Search
// @ID get-personal-search
// @Produce  json
// @Param personal query string false "Personal Search"
// @Success 200 {object} ent.Personal
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /searchpersonals [get]
func (ctl *PersonalController) GetPersonalSearch(c *gin.Context) {
	psearch := c.Query("personal")

	ps, err := ctl.client.Personal.
		Query().
		WithDepartment().
		WithGender().
		WithTitle().
		Where(personal.PersonalnameContains(psearch)).
		All(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": ps,
	})
}

// NewPersonalController creates and registers handles for the personal controller
func NewPersonalController(router gin.IRouter, client *ent.Client) *PersonalController {
	pc := &PersonalController{
		client: client,
		router: router,
	}
	pc.register()
	return pc
}

// InitPersonalController registers routes to the main engine
func (ctl *PersonalController) register() {
	personals := ctl.router.Group("/personals")

	personals.GET("", ctl.ListPersonal)

	searchpersonals := ctl.router.Group("/searchpersonals")
	searchpersonals.GET("", ctl.GetPersonalSearch)

	// CRUD
	personals.POST("", ctl.CreatePersonal)
	personals.GET(":id", ctl.GetPersonal)
	personals.PUT(":id", ctl.UpdatePersonal)
	personals.DELETE(":id", ctl.DeletePersonal)

}
