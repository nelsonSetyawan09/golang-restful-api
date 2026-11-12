package service

import (
	"belajar-golang-restful-api/model/web"
	"context"
)

type CategoryService interface {
	FindAll(ctx context.Context) []web.CategoryResponse
	FindByID(ctx context.Context, categoryId int) web.CategoryResponse
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
}
