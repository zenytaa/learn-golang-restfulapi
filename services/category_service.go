package services

import (
	"context"
	"golang_restfulapi/model/web"
)

// create contract
// dont return category repository so didnt expose table database
// business logic

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	DeleteById(ctx context.Context, categoryId int)
	DeleteAll(ctx context.Context)
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
