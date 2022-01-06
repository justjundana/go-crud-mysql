package repository

import (
	"crud-database/models"
	"database/sql"
	"fmt"
	"log"
)

type UserRepository interface {
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

// get all users
func (r *userRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	rows, err := r.db.Query(`SELECT id, name, email FROM users ORDER BY id ASC`)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Println(rows)

	defer rows.Close()

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
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
