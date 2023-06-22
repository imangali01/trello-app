package repository

import (
	"fmt"

	"database/sql"

	"github.com/Imangali2002/trello-app/view"
)

type UserRepository struct {
	db *sql.DB
}

func UserRepositoryInit(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) RegisterUser(user view.UserCreate) error {
	_, err := r.db.Exec("INSERT INTO public.user (username, email) VALUES ($1, $2)", user.Username, user.Email)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetAllUsers() ([]view.UserInfo, error) {
	var users []view.UserInfo

	rows, err := r.db.Query(`SELECT id, username, email FROM public.user`)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user view.UserInfo

		err = rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}

func (r *UserRepository) GetUserByID(userID uint) (view.UserInfo, error) {
	var user view.UserInfo

	row := r.db.QueryRow(`SELECT id, username, email FROM public.user WHERE id = $1`, userID)

	err := row.Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user not found")
		}
		return user, err
	}

	return user, nil
}
