package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/raynaaliyya/ctrl-wh-st/models"
)

type ProductWhRepo interface {
	GetAll() any
	GetById(id int) any
	GetByName(name string) (*models.ProductWh, error)
	Create(newProduct *models.ProductWh) (*models.ProductWh, error)
	Update(product *models.ProductWh) string
	Delete(id int) string
}

type productWhRepo struct {
	db *sql.DB
}

func NewProductWhRepo(db *sql.DB) ProductWhRepo {
	repo := new(productWhRepo)
	repo.db = db

	return repo
}

func (r *productWhRepo) GetAll() any {
	var products []models.ProductWh

	query := "SELECT id, product_name, price, product_category, stock FROM product_wh"

	rows, err := r.db.Query(query)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var product models.ProductWh

		if err := rows.Scan(&product.ID, &product.ProductName, &product.Price, &product.ProductCategory, &product.Stock); err != nil {
			log.Println(err)
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	if len(products) == 0 {
		return "no data"
	}

	return products
}

func (r *productWhRepo) GetById(id int) any {
	var productInDb models.ProductWh

	query := "SELECT id, product_name, price, product_category, stock FROM product_wh WHERE id = $1"
	row := r.db.QueryRow(query, id)

	err := row.Scan(&productInDb.ID, &productInDb.ProductName, &productInDb.Price, &productInDb.ProductCategory, &productInDb.Stock)

	if err != nil {
		log.Println(err)
	}

	if productInDb.ID == 0 {
		return "product not found"
	}

	return productInDb
}

func (r *productWhRepo) GetByName(name string) (*models.ProductWh, error) {
	var productInDb models.ProductWh

	query := "SELECT id, product_name, price, product_category, stock FROM product_wh WHERE product_name = $1"
	row := r.db.QueryRow(query, name)

	err := row.Scan(&productInDb.ID, &productInDb.ProductName, &productInDb.Price, &productInDb.ProductCategory, &productInDb.Stock)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found")
		}
		return nil, err
	}

	return &productInDb, nil
}

func (r *productWhRepo) Create(newProduct *models.ProductWh) (*models.ProductWh, error) {
	query := "INSERT INTO product_wh(id, product_name, price, product_category, stock) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.db.Exec(query, newProduct.ID, newProduct.ProductName, newProduct.Price, newProduct.ProductCategory, newProduct.Stock)

	if err != nil {
		log.Println(err)
		return &models.ProductWh{}, fmt.Errorf("failed to create product")
	}

	return newProduct, nil
}

func (r *productWhRepo) Update(product *models.ProductWh) string {
	res, err := r.GetByName(product.ProductName) //respon

	// jika tidak ada, return pesan
	if err != nil {
		return err.Error()
	}

	// jika ada maka update user
	query := "UPDATE product_wh SET id = $1, price = $2, product_category = $3, stock = $4 WHERE product_name = $5"
	_, err = r.db.Exec(query, product.ID, product.Price, product.ProductCategory, product.Stock, product.ProductName)

	if err != nil {
		log.Println(err)
		return "failed to update product"
	}

	return fmt.Sprintf("product %s updated successfully", res.ProductName)
}

func (r *productWhRepo) Delete(id int) string {
	res := r.GetById(id)

	// jika tidak ada, return pesan
	if res == "product not found" {
		return res.(string)
	}

	// jika ada, delete user
	query := "DELETE FROM product_wh WHERE id = $1"
	_, err := r.db.Exec(query, id)

	if err != nil {
		log.Println(err)
		return "failed to delete product"
	}

	return fmt.Sprintf("product with id %d deleted successfully", id)
}
