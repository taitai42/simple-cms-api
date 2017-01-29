package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/taitai42/simple-cms-api/model"
)

type Headers struct {
	Authorization string
}

//GetHome define GetHome parameters
//swagger:params GetHomeAction
type GetHomeParams struct {
	Headers
}

// swagger:routes GET /home GetHomeAction
// CreateCard api entry point to create a card
func GetHome(c echo.Context) error {

	home := &model.Home{}
	if err := home.GetHome(); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, home)
}

type UpdateHomeBody struct {
	Title       string `json:"title"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
}

//params required by update home endpoint
//swagger:params UpdateHomeAction
type UpdateHomeParams struct {
	Headers

	//in: body
	Body UpdateHomeBody
}

//swagger:routes PUT /home UpdateHomeAction
func UpdateHome(c echo.Context) error {
	params := &UpdateHomeParams{}
	if err := c.Bind(&params.Body); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	home := &model.Home{}
	if err := home.GetHome(); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if params.Body.Title != "" {
		home.Title = params.Body.Title
	}

	if params.Body.ImageURL != "" {
		home.ImageURL = params.Body.ImageURL
	}

	if params.Body.Description != "" {
		home.Description = params.Body.Description
	}

	if err := home.UpdateHome(); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, home)
}


type CreateHomeBody struct {
	Title       string `json:"title"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
}

//params required by update home endpoint
//swagger:params CreateHomeAction
type CreateHomeParams struct {
	Headers

	//in: body
	Body CreateHomeBody
}

//swagger:routes POST /home CreateHomeAction
func CreateHome(c echo.Context) error {
	params := &CreateHomeParams{}
	if err := c.Bind(&params.Body); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	home := &model.Home{}
	if err := home.GetHome(); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if params.Body.Title != "" {
		home.Title = params.Body.Title
	}

	if params.Body.ImageURL != "" {
		home.ImageURL = params.Body.ImageURL
	}

	if params.Body.Description != "" {
		home.Description = params.Body.Description
	}

	if err := home.UpdateHome(); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, home)
}

