package dblayer

import "database/sql"

type Authorization interface {
	AddUser(username, email, password string) error
	GetUserLogin(username string) (string, error)
	// GetUser(username, email string) error
	UserExists(username, email string) (bool, error)
	// DeleteUser(username, email, password string) error
	// ChangeUserEmail(username, email, password string) error
	// ChangeUserName(username, email, password string) error
}

//check interface implementation by the AuthQuerier struct at runtime

type AuthQuerier struct {
	DB *sql.DB
}

func authNew(db *sql.DB) *AuthQuerier {
	return &AuthQuerier{
		DB: db,
	}
}

func (aq *AuthQuerier) AddUser(username, email, password string) error {
	stmt := `
				INSERT INTO users (name, email, password)
				VALUES ($1, $2, $3)
			`
	_, err := aq.DB.Exec(stmt, username, email, password)
	if err != nil {
		return err
	}
	return nil
}

func (aq *AuthQuerier) UserExists(username, email string) (bool, error) {
	var exists bool
	stmt := `SELECT EXISTS ( SELECT id FROM users WHERE name = $1 OR email = $2 LIMIT 1 )`
	err := aq.DB.QueryRow(stmt, username, email).Scan(&exists)
	if err != nil {
		return exists, err
	}
	return exists, nil
}

func (aq *AuthQuerier) GetUserLogin(name string) (string, error) {
	var password string

	stmt := `SELECT password FROM users WHERE name = $1 LIMIT 1 `
	err := aq.DB.QueryRow(stmt, name).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return password, nil
		}
		return password, err
	}
	return password, nil
}
