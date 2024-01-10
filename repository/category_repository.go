package repository

import (
	"context"
	"database/sql"
	"golang_restfulapi/model/entity"
)

// create contract use interface

type CategoryRepository interface {
	Create(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category
	Update(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category
	DeleteById(ctx context.Context, tx *sql.Tx, category entity.Category)
	DeleteAll(ctx context.Context, tx *sql.Tx)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (entity.Category, error)
	FindALl(ctx context.Context, tx *sql.Tx) []entity.Category
}
