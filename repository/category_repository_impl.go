package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang_restfulapi/helper"
	"golang_restfulapi/model/entity"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	SQL := "INSERT INTO categories(name) VALUES (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.IfErrorPanic(err)
	// get last id
	id, err := result.LastInsertId()
	helper.IfErrorPanic(err)
	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	SQL := "UPDATE categories SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.IfErrorPanic(err)
	return category
}

func (repository *CategoryRepositoryImpl) DeleteById(ctx context.Context, tx *sql.Tx, category entity.Category) {
	SQL := "DELETE FROM categories WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.IfErrorPanic(err)
}

func (repository *CategoryRepositoryImpl) DeleteAll(ctx context.Context, tx *sql.Tx) {
	SQL := "TRUNCATE categories"
	_, err := tx.ExecContext(ctx, SQL)
	helper.IfErrorPanic(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (entity.Category, error) {
	SQL := "SELECT id, name FROM categories WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.IfErrorPanic(err)
	defer rows.Close()

	category := entity.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.IfErrorPanic(err)
		return category, nil
	} else {
		return category, errors.New("no category found")
	}

}

func (repository *CategoryRepositoryImpl) FindALl(ctx context.Context, tx *sql.Tx) []entity.Category {
	SQL := "SELECT id, name FROM categories"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.IfErrorPanic(err)
	defer rows.Close()

	var categories []entity.Category
	for rows.Next() {
		category := entity.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.IfErrorPanic(err)
		categories = append(categories, category)
	}
	return categories
}
