package repository

import (
	"database/sql"

	"clean-architecture/entity"
)

//UserMySQL mysql repo
type UserMySQL struct {
	db *sql.DB
}

//NewUserMySQL create new repository
func NewUserMySQL(db *sql.DB) *UserMySQL {
	return &UserMySQL{
		db: db,
	}
}

//Get an user
func (r *UserMySQL) Get(id int64) (*entity.User, error) {
	return getUser(id, r.db)
}

func getUser(id int64, db *sql.DB) (*entity.User, error) {
	stmt, err := db.Prepare(`select id, email, first_name, last_name, created_at from user where id = ?`)
	if err != nil {
		return nil, err
	}
	var user entity.User
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.CreatedAt)
	}
	return &user, nil
}
