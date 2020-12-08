package data

import (
	"fmt"
	"github.com/rvalessandro/mf-backend/datasources/mysql"
	"github.com/rvalessandro/mf-backend/modules/products/domain"
	"github.com/rvalessandro/mf-backend/utils/errors"
	"github.com/rvalessandro/mf-backend/utils/mysql_util"
)

const (
	queryFindProduct = `SELECT id, name, description, image_url, price, created_at, updated_at FROM products;`
	queryGetProduct  = `
		SELECT id, name, description, image_url, price, created_at, updated_at
		FROM products
		WHERE id = ?
		LIMIT 1;
	`
	queryCreateProduct = `
		INSERT INTO products (name, description, image_url, price, created_at, updated_at)
		VALUES (
			?, ?, ?, ?, ?, ?
		)
	`
	queryUpdateProduct = `
		UPDATE products
		SET name=?,
			description=?,
			image_url=?,
			price=?,
			updated_at=?
		WHERE id = ?
	`
	queryDeleteProduct = `DELETE FROM products WHERE id=?`
)

func Find() ([]domain.Product, *errors.RestErr) {
	stmt, err := mysql.Client.Prepare(queryFindProduct)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer rows.Close()

	results := make([]domain.Product, 0)
	for rows.Next() {
		var product domain.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.ImageURL, &product.Price, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, errors.NewErrInternalServer(err.Error())
		}
		results = append(results, product)
	}

	if len(results) == 0 {
		return nil, errors.NewErrNotFound(fmt.Sprintf("No products found"))
	}

	return results, nil
}

/**
 * TODO Generate product list
 */
func Get(id int64) (*domain.Product, *errors.RestErr) {
	product := domain.Product{}
	stmt, err := mysql.Client.Prepare(queryGetProduct)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(id)
	err = result.Scan(&product.ID, &product.Name, &product.Description, &product.ImageURL, &product.Price, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, mysql_util.ParseError(err)
	}

	return &product, nil
}

func Create(ProductParam domain.CreateProductParams) (*domain.Product, *errors.RestErr) {
	stmt, err := mysql.Client.Prepare(queryCreateProduct)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	queryRes, err := stmt.Exec(ProductParam.Name, ProductParam.Description, ProductParam.ImageURL, ProductParam.Price, ProductParam.CreatedAt, ProductParam.UpdatedAt)
	if err != nil {
		return nil, mysql_util.ParseError(err)
	}

	newID, err := queryRes.LastInsertId()

	return Get(newID)
}

func Update(id int64, ProductParam domain.CreateProductParams) (*domain.Product, *errors.RestErr) {
	stmt, err := mysql.Client.Prepare(queryUpdateProduct)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(ProductParam.Name, ProductParam.Description, ProductParam.ImageURL, ProductParam.Price, ProductParam.UpdatedAt, id)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}

	return Get(id)
}

func Delete(id int64) *errors.RestErr {
	stmt, err := mysql.Client.Prepare(queryDeleteProduct)
	if err != nil {
		return errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return errors.NewErrInternalServer(err.Error())
	}

	return nil
}
