package services

import (
	"context"
	"database/sql"
	"golang_restfulapi/exception"
	"golang_restfulapi/helper"
	"golang_restfulapi/model/entity"
	"golang_restfulapi/model/web"
	"golang_restfulapi/repository"

	"github.com/go-playground/validator"
)

// business logic

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

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	// validation for create and update category
	err := service.Validate.Struct(request)
	helper.IfErrorPanic(err)

	// start transactional sql
	tx, err := service.DB.Begin()
	helper.IfErrorPanic(err)
	defer helper.CommitOrRollback(tx)

	// create category from entity repository contract
	category := entity.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Create(ctx, tx, category)

	// return category should be convert from category to CategoryResponse
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	// validation for create and update category
	err := service.Validate.Struct(request)
	helper.IfErrorPanic(err)

	tx, err := service.DB.Begin()
	helper.IfErrorPanic(err)
	defer helper.CommitOrRollback(tx)

	// validation for id
	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// update name
	category.Name = request.Name

	service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) DeleteById(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.IfErrorPanic(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.DeleteById(ctx, tx, category)
}

func (service *CategoryServiceImpl) DeleteAll(ctx context.Context) {
	tx, err := service.DB.Begin()
	helper.IfErrorPanic(err)
	defer helper.CommitOrRollback(tx)

	service.CategoryRepository.DeleteAll(ctx, tx)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.IfErrorPanic(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.IfErrorPanic(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindALl(ctx, tx)
	return helper.ToCategoryResponses(categories)
}
