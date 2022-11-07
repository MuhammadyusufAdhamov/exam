package main

import (
	"database/sql"
	"fmt"
	"time"
)

type DBProductManager struct {
	db *sql.DB
}

func NewDBProductManager(db *sql.DB) *DBProductManager {
	return &DBProductManager{
		db: db,
	}
}

type Products struct {
	Id        int64
	Name      string
	Price     float64
	ImageUrl  string
	CreatedAt time.Time
	Images    []*ProductImages
}

type ProductImages struct {
	Id             int32
	ImageUrl       string
	SequenceNumber int32
}

type GetProductsParams struct {
	Limit  int32
	Page   int32
	Search string
}

type GetProductsResponse struct {
	Product []*Products
	Count   int32
}

func (m *DBProductManager) CreateProducts(product *Products) (int64, error) {
	var productId int64
	query := `
		insert into products (name, price, image_url) 
		values ($1,$2,$3)
		returning id
	`

	row := m.db.QueryRow(
		query,
		product.Name,
		product.Price,
		product.ImageUrl,
	)

	err := row.Scan(&productId)
	if err != nil {
		return 0, err
	}

	queryInsertImage := `
		insert into product_images (product_id,image_url,sequence_number)
		values ($1,$2,$3)
`
	for _, image := range product.Images {
		_, err := m.db.Exec(
			queryInsertImage,
			productId,
			image.ImageUrl,
			image.SequenceNumber,
		)
		if err != nil {
			return 0, err
		}
	}

	return productId, nil
}

func (m *DBProductManager) GetProducts(id int64) (*Products, error) {
	var product Products

	product.Images = make([]*ProductImages, 0)
	query := `
		select
		   id,
		   name,
		   price,
		   image_url,
		   created_at
		from products
		where id=$1
`

	row := m.db.QueryRow(query, id)

	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.ImageUrl,
		&product.CreatedAt,
	)

	if err != nil {
		return nil, err
	}
	queryImages := `
		select 
			id,
			image_url,
			sequence_number
		from product_images
		where product_id=$1
`
	rows, err := m.db.Query(queryImages, id)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var image ProductImages

		err := rows.Scan(
			&image.Id,
			&image.ImageUrl,
			&image.SequenceNumber,
		)
		if err != nil {
			return nil, err
		}
		product.Images = append(product.Images, &image)
	}

	return &product, nil
}

func (m *DBProductManager) GetAllProducts(params *GetProductsParams) (*GetProductsResponse, error) {
	var result GetProductsResponse

	result.Product = make([]*Products, 0)

	filter := ""
	if params.Search != "" {
		filter = fmt.Sprintf("Where name ilike '%s'",
			"%s"+params.Search+"%")
	}

	query := `
		select
		    id,
		    name,
		    price,
			image_url,
		    created_at
		from products
		` + filter + `
		order by created_at desc limit $1 offset $2
`

	offset := (params.Page - 1) * params.Limit
	rows, err := m.db.Query(query, params.Limit, offset)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var product Products

		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.ImageUrl,
			&product.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		result.Product = append(result.Product, &product)
	}

	return &result, nil
}

func (m *DBProductManager) UpdateProduct(product *Products) error {
	query := `
		update products set
		                    name=$1,
		                    price=$2,
		                    image_url=$3
		where id=$4
`
	result, err := m.db.Exec(
		query,
		product.Name,
		product.Price,
		product.ImageUrl,
		product.Id,
	)

	if err != nil {
		return err
	}

	rowsCount, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsCount == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (m *DBProductManager) DeleteProduct(id int64) error {
	queryDeleteImages := `delete from product_images where product_id=$1`

	_, err := m.db.Exec(queryDeleteImages, id)
	if err != nil {
		return err
	}

	queryDelete := `delete from products where id=$1`

	result, err := m.db.Exec(queryDelete, id)
	if err != nil {
		return err
	}

	rowsCount, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsCount == 0 {
		return sql.ErrNoRows
	}

	return nil
}
