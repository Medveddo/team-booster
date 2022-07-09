package api

import (
	"net/http"

	"github.com/Medveddo/team-booster/backend/models"
	"github.com/Medveddo/team-booster/backend/storage"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type API struct {
	db storage.ISkillStorage
}

func NewAPI() *API {
	db := storage.NewSkillStorage()
	return &API{
		db: db,
	}
}

func (api *API) Close() error {
	err := api.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (api *API) GetSkills(c echo.Context) error {
	skills, err := api.db.GetSkills()
	if err != nil {
		c.Error(err)
	}
	return c.JSON(http.StatusOK, skills)
}

func (api *API) GetSkill(c echo.Context) error {
	id := c.Param("id")
	skillID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		msg := struct {
			Message string `json:"message"`
		}{
			Message: "Invalid ObjectID",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}
	skill, err := api.db.GetSkill(skillID)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, skill)
}

func (api *API) UpdateSkill(c echo.Context) error {
	id := c.Param("id")
	skill := new(models.Skill)
	if err := c.Bind(skill); err != nil {
		c.Logger().Error(err)
		return err
	}
	skillID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		msg := struct {
			Message string `json:"message"`
		}{
			Message: "Invalid ObjectID",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}
	err = api.db.UpdateSkill(skillID, skill)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, struct{}{})
}

func (api *API) Health(c echo.Context) error {
	response := struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	}

	return c.JSON(http.StatusOK, response)
}
