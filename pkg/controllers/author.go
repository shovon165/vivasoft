package controllers

import (
	"book-crud/pkg/domain"
	"book-crud/pkg/types"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IAuthorController interface {
	CreateAuthor(e echo.Context) error
	GetAuthor(e echo.Context) error
	GetAuthors(e echo.Context) error
	UpdateAuthor(e echo.Context) error
	DeleteAuthor(e echo.Context) error
}

type AuthorController struct {
	authorSvc domain.IAuthorService
}

func NewAuthorController(authorSvc domain.IAuthorService) AuthorController {
	return AuthorController{
		authorSvc: authorSvc,
	}
}

func (controller *AuthorController) CreateAuthor(e echo.Context) error {
	request := &types.CreateAuthorRequest{}
	if err := e.Bind(request); err != nil {
		return e.JSON(http.StatusBadRequest, "invalid request body")
	}
	if err := request.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	if err := controller.authorSvc.CreateAuthor(request); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusCreated, "author was created successfully")
}

func (controller *AuthorController) GetAuthor(e echo.Context) error {
	tempAuthorID := e.Param("id")
	authorID, err := strconv.ParseUint(tempAuthorID, 0, 0)
	if err != nil && tempAuthorID != "" {
		return e.JSON(http.StatusBadRequest, "enter a valid author ID")
	}

	response, err := controller.authorSvc.GetAuthor(uint(authorID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, response)
}

func (controller *AuthorController) GetAuthors(e echo.Context) error {
	request := make(map[string]string)
	if e.QueryParam("id") != "" {
		if _, err := strconv.ParseUint(e.QueryParam("id"), 0, 0); err != nil {
			return e.JSON(http.StatusBadRequest, "enter a valid author ID")
		}
		request["ID"] = e.QueryParam("id")
	}
	if e.QueryParam("authorName") != "" {
		request["AuthorName"] = e.QueryParam("authorName")
	}
	if e.QueryParam("address") != "" {
		request["Address"] = e.QueryParam("address")
	}
	if e.QueryParam("phoneNumber") != "" {
		request["PhoneNumber"] = e.QueryParam("phoneNumber")
	}

	// pass the request to the service layer
	response, err := controller.authorSvc.GetAuthors(request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, response)
}

func (controller *AuthorController) UpdateAuthor(e echo.Context) error {
	request := &types.UpdateAuthorRequest{}
	if err := e.Bind(request); err != nil {
		return e.JSON(http.StatusBadRequest, "invalid request body")
	}

	if err := request.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	tempAuthorID := e.Param("id")
	authorID, err := strconv.ParseInt(tempAuthorID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "enter a valid author ID")
	}
	if err := controller.authorSvc.UpdateAuthor(uint(authorID), request); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, "author was updated successfully")
}

func (controller *AuthorController) DeleteAuthor(e echo.Context) error {
	tempAuthorID := e.Param("id")
	authorID, err := strconv.ParseInt(tempAuthorID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "enter a valid author ID")
	}
	if err := controller.authorSvc.DeleteAuthor(uint(authorID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "author and books of author was deleted successfully")
}
