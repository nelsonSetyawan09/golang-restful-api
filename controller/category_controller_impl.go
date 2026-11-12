package controller

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/web"
	"belajar-golang-restful-api/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	categoryResponses := controller.CategoryService.FindAll(r.Context())

	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "Categories found successfully",
		Data:    categoryResponses,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *CategoryControllerImpl) FindByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryID := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryID)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindByID(r.Context(), id)

	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "Category found successfully",
		Data:    categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var categoryCreateRequest web.CategoryCreateRequest
	helper.ReadFromRequestBody(r, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(r.Context(), categoryCreateRequest)

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "success",
		Message: "Category created successfully",
		Data:    categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var categoryUpdateRequest web.CategoryUpdateRequest
	helper.ReadFromRequestBody(r, &categoryUpdateRequest)

	categoryID := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryID)
	helper.PanicIfError(err)
	categoryUpdateRequest.ID = id

	categoryResponse := controller.CategoryService.Update(r.Context(), categoryUpdateRequest)

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "success",
		Message: "Category updated successfully",
		Data:    categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryID := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryID)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(r.Context(), id)

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "success",
		Message: "Category delete successfully",
	}

	helper.WriteToResponseBody(w, webResponse)
}
