package data

import (
	"fmt"
	"github.com/rvalessandro/mf-backend/datasources/mysql"
	"github.com/rvalessandro/mf-backend/modules/categories/domain"
	"github.com/rvalessandro/mf-backend/utils/errors"
	"github.com/rvalessandro/mf-backend/utils/mysql_util"
)

const (
	queryFindCategory   = `SELECT id, name, created_at, updated_at FROM categories;`
	queryGetCategory    = `SELECT id, name, created_at, updated_at from categories where id=?;`
	queryCreateCategory = `
		INSERT INTO categories (name, created_at, updated_at)
		VALUES (
			?, ?, ?
		)`
	queryUpdateCategory = `UPDATE categories SET name=?, updated_at=? WHERE id=?`
	queryDeleteCategory = `DELETE FROM categories WHERE id=$1`
)

func Find() ([]domain.Category, *errors.RestErr) {
	stmt, err := mysql.Client.Prepare(queryFindCategory)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer rows.Close()

	results := make([]domain.Category, 0)
	for rows.Next() {
		var category domain.Category
		err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, errors.NewErrInternalServer(err.Error())
		}
		results = append(results, category)
	}

	if len(results) == 0 {
		return nil, errors.NewErrNotFound(fmt.Sprintf("No users found"))
	}

	return results, nil
}

func Get(categoryID int64) (*domain.Category, *errors.RestErr) {
	category := domain.Category{}
	stmt, err := mysql.Client.Prepare(queryGetCategory)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(categoryID)
	err = result.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		return nil, mysql_util.ParseError(err)
	}

	return &category, nil
}

func Create(categoryParam domain.CreateCategoryParams) (*domain.Category, *errors.RestErr) {
	stmt, err := mysql.Client.Prepare(queryCreateCategory)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	queryRes, err := stmt.Exec(categoryParam.Name, categoryParam.CreatedAt, categoryParam.UpdatedAt)
	if err != nil {
		return nil, mysql_util.ParseError(err)
	}

	newID, err := queryRes.LastInsertId()

	return Get(newID)
}

func Update(id int64, categoryParam domain.CreateCategoryParams) (*domain.Category, *errors.RestErr) {
	stmt, err := mysql.Client.Prepare(queryUpdateCategory)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(categoryParam.Name, categoryParam.UpdatedAt, id)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}

	return Get(id)
}

func Delete(categoryID int64) *errors.RestErr {
	stmt, err := mysql.Client.Prepare(queryDeleteCategory)
	if err != nil {
		return errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(categoryID)
	if err != nil {
		return errors.NewErrInternalServer(err.Error())
	}

	return nil
}
