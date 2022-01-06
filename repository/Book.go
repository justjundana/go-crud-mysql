package repository

import (
	"crud-database/models"
	"database/sql"
	"fmt"
	"log"
)

type BookRepository interface {
	GetBooks() ([]models.Book, error)
	CreateBook(book models.Book) (models.Book, error)
	GetBook(id int) (models.Book, error)
	UpdateBook(book models.Book) (models.Book, error)
	DeleteBook(book models.Book) (models.Book, error)
}

type bookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *bookRepository {
	return &bookRepository{db}
}

// get all Books
func (r *bookRepository) GetBooks() ([]models.Book, error) {
	var books []models.Book
	rows, err := r.db.Query(`SELECT id, code, title, description, author, publisher FROM books ORDER BY id ASC`)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Println(rows)

	defer rows.Close()

	for rows.Next() {
		var book models.Book

		err = rows.Scan(&book.ID, &book.Code, &book.Title, &book.Description, &book.Author, &book.Publisher)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		books = append(books, book)
	}

	return books, nil
}

// create new Book
func (r *bookRepository) CreateBook(book models.Book) (models.Book, error) {
	query := `INSERT INTO books (code, title, description, author, publisher) VALUES (?, ?, ?, ?, ?)`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return book, err
	}

	defer statement.Close()

	_, err = statement.Exec(book.Code, book.Title, book.Description, book.Author, book.Publisher)

	return book, nil
}

// get Book by id
func (r *bookRepository) GetBook(id int) (models.Book, error) {
	var book models.Book

	row := r.db.QueryRow(`SELECT id, code, title, description, author, publisher FROM books WHERE id = ?`, id)

	err := row.Scan(&book.ID, &book.Code, &book.Title, &book.Description, &book.Author, &book.Publisher)
	if err != nil {
		return book, err
	}

	return book, nil
}

// Update Book
func (r *bookRepository) UpdateBook(book models.Book) (models.Book, error) {
	query := `UPDATE books SET title = ?, description = ?, author = ?, publisher = ? WHERE id = ?`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return book, err
	}

	defer statement.Close()

	_, err = statement.Exec(book.Title, book.Description, book.Author, book.Publisher, book.ID)
	if err != nil {
		return book, err
	}

	return book, nil
}

// Delete Book
func (r *bookRepository) DeleteBook(book models.Book) (models.Book, error) {
	query := `DELETE FROM books WHERE id = ?`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return book, err
	}

	defer statement.Close()

	_, err = statement.Exec(book.ID)
	if err != nil {
		return book, err
	}

	return book, nil
}
