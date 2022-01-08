package helper

import (
	"github.com/justjundana/go-crud-mysql/models"
)

type AuthFormat struct {
	Token string `json:"token"`
}

func FormatAuth(user models.User, token string) AuthFormat {
	formatter := AuthFormat{
		Token: token,
	}
	return formatter
}

type UserFormatter struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FormatUser(user models.User) UserFormatter {
	formatter := UserFormatter{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	return formatter
}

type BookFormatter struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
}

func FormatBook(book models.Book) BookFormatter {
	formatter := BookFormatter{
		ID:          book.ID,
		Code:        book.Code,
		Title:       book.Title,
		Description: book.Description,
		Author:      book.Author,
		Publisher:   book.Publisher,
	}
	return formatter
}
