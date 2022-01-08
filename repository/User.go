package repository

import (
	"database/sql"
	"log"

	"github.com/justjundana/go-crud-mysql/models"
)

type UserRepository interface {
	FindByEmail(email string) (models.User, error)
	GetUsers() ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
	GetUser(id int) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(user models.User) (models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) FindByEmail(email string) (models.User, error) {
	row := r.db.QueryRow(`SELECT id, email, password FROM users WHERE email = ?;`, email)
	var user models.User

	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

// get all users
func (r *userRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	rows, err := r.db.Query(`SELECT id, name, email FROM users ORDER BY id ASC`)
	if err != nil {
		log.Fatalf("Error")
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			log.Fatalf("Error")
		}

		users = append(users, user)
	}

	return users, nil
}

// create new user
func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	query := `INSERT INTO users (name, email, password) VALUES (?, ?, ?)`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return user, err
	}

	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Email, user.Password)

	return user, nil
}

// get user by id
func (r *userRepository) GetUser(id int) (models.User, error) {
	var user models.User

	row := r.db.QueryRow(`SELECT id, name, email FROM users WHERE id = ?`, id)

	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Update User
func (r *userRepository) UpdateUser(user models.User) (models.User, error) {
	query := `UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return user, err
	}

	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Email, user.Password, user.ID)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Delete User
func (r *userRepository) DeleteUser(user models.User) (models.User, error) {
	query := `DELETE FROM users WHERE id = ?`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return user, err
	}

	defer statement.Close()

	_, err = statement.Exec(user.ID)
	if err != nil {
		return user, err
	}

	return user, nil
}
