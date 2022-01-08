package service

import (
	"time"

	"github.com/justjundana/go-crud-mysql/helper"
	"github.com/justjundana/go-crud-mysql/models"
	"github.com/justjundana/go-crud-mysql/repository"
)

type BookService interface {
	GetBooksService() ([]models.Book, error)
	GetBookService(id int) (models.Book, error)
	CreateBookService(input helper.CreateBookRequest) (models.Book, error)
	UpdateBookService(id int, input helper.EditBookRequest) (models.Book, error)
	DeleteBookService(id int) (models.Book, error)
}

type bookService struct {
	repository repository.BookRepository
}

func NewBookService(repository repository.BookRepository) *bookService {
	return &bookService{repository}
}

func (s *bookService) GetBooksService() ([]models.Book, error) {
	books, err := s.repository.GetBooks()
	if err != nil {
		return books, err
	}
	return books, nil
}

func (s *bookService) GetBookService(id int) (models.Book, error) {
	book, err := s.repository.GetBook(id)
	if err != nil {
		return book, err
	}
	return book, nil
}

func (s *bookService) CreateBookService(input helper.CreateBookRequest) (models.Book, error) {
	book := models.Book{}
	book.Code = input.Code
	book.Title = input.Title
	book.Description = input.Description
	book.Author = input.Author
	book.Publisher = input.Publisher
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	createBook, err := s.repository.CreateBook(book)
	if err != nil {
		return createBook, err
	}
	return createBook, nil
}

func (s *bookService) UpdateBookService(id int, input helper.EditBookRequest) (models.Book, error) {
	book, err := s.repository.GetBook(id)
	if err != nil {
		return book, err
	}

	book.Title = input.Title
	book.Description = input.Description
	book.Author = input.Author
	book.Publisher = input.Publisher
	book.UpdatedAt = time.Now()

	updateBook, err := s.repository.UpdateBook(book)
	if err != nil {
		return updateBook, err
	}
	return updateBook, nil
}

func (s *bookService) DeleteBookService(id int) (models.Book, error) {
	bookID, err := s.GetBookService(id)
	if err != nil {
		return bookID, err
	}

	deleteBook, err := s.repository.DeleteBook(bookID)
	if err != nil {
		return deleteBook, err
	} else {
		return deleteBook, nil
	}
}
