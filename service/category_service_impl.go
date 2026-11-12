package service

import (
	"belajar-golang-restful-api/exception"
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/domain"
	"belajar-golang-restful-api/model/web"
	"belajar-golang-restful-api/repository"
	"context"
	"database/sql"
	"errors"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	beginTx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(beginTx)

	categories := service.CategoryRepository.FindAll(ctx, beginTx)

	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, helper.ToCategoryResponse(category))
	}

	return categoryResponses
}

func (service *CategoryServiceImpl) FindByID(
	ctx context.Context,
	categoryId int,
) web.CategoryResponse {

	beginTx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(beginTx)

	category, err := service.CategoryRepository.FindByID(
		ctx,
		beginTx,
		categoryId,
	)

	if errors.Is(err, sql.ErrNoRows) {
		panic(exception.NewNotFoundHandler("Category not found"))
	}

	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	beginTx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(beginTx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, beginTx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	beginTx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(beginTx)

	category, err := service.CategoryRepository.FindByID(ctx, beginTx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundHandler(err.Error()))
	}

	category.Name = request.Name

	category = service.CategoryRepository.Update(ctx, beginTx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	beginTx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(beginTx)

	category, err := service.CategoryRepository.FindByID(ctx, beginTx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundHandler(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, beginTx, category)
}
