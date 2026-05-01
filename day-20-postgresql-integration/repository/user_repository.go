package repository

import (
	"database/sql"
	"postgresql-integration/models"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	rows, err := r.DB.Query("SELECT id, name, age FROM users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Name, &u.Age)

		users = append(users, u)
	}

	return users, nil
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users(name, age) VALUES($1, $2)"

	_, err := r.DB.Exec(query, user.Name, user.Age)

	return err
}
