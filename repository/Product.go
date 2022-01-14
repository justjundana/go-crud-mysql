package repository

import (
	"database/sql"
	"log"

	"github.com/justjundana/go-crud-mysql/models"
)

type ProductRepositoryInterface interface {
	GetProducts() ([]models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	GetProduct(id int) (models.Product, error)
	UpdateProduct(product models.Product) (models.Product, error)
	DeleteProduct(product models.Product) (models.Product, error)
}

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db}
}

// get all Products
func (r *ProductRepository) GetProducts() ([]models.Product, error) {
	var products []models.Product
	rows, err := r.db.Query(`SELECT  products.id, products.name, products.description, products.price, products.user_id, users.name, users.email FROM products  JOIN users ON products.user_id = users.id ORDER BY products.id ASC`)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var product models.Product

		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.User.ID, &product.User.Name, &product.User.Email)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		products = append(products, product)
	}

	return products, nil
}

// create new Product
func (r *ProductRepository) CreateProduct(product models.Product) (models.Product, error) {
	query := `INSERT INTO products (name, description, price, user_id) VALUES (?, ?, ?, ?)`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return product, err
	}

	defer statement.Close()

	_, err = statement.Exec(product.Name, product.Description, product.Price, product.User.ID)
	if err != nil {
		return product, err
	}

	return product, nil
}

// get Product by id
func (r *ProductRepository) GetProduct(id int) (models.Product, error) {
	var product models.Product

	row := r.db.QueryRow(`SELECT products.id, products.name, products.description, products.price, products.user_id, users.name, users.email  FROM products  JOIN users ON products.user_id = users.id  WHERE products.id = ?`, id)

	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.User.ID, &product.User.Name, &product.User.Email)
	if err != nil {
		return product, err
	}

	return product, nil
}

// Update Product
func (r *ProductRepository) UpdateProduct(product models.Product) (models.Product, error) {
	query := `UPDATE products SET name = ?, description = ?, price = ?, user_id = ? WHERE id = ?`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return product, err
	}

	defer statement.Close()

	_, err = statement.Exec(product.Name, product.Description, product.Price, product.User.ID, product.ID)
	if err != nil {
		return product, err
	}

	return product, nil
}

// Delete Product
func (r *ProductRepository) DeleteProduct(product models.Product) (models.Product, error) {
	query := `DELETE FROM products WHERE id = ?`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return product, err
	}

	defer statement.Close()

	_, err = statement.Exec(product.ID)
	if err != nil {
		return product, err
	}

	return product, nil
}
