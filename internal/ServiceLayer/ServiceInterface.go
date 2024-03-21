package servicelayer

import (
	"database/sql"
	"errors"

	dblayer "github.com/Adil-9/online_store/internal/DBlayer"
	structures "github.com/Adil-9/online_store/internal/Structures"
)

var (
	ErrEmailTaken            = errors.New("email already in use")
	ErrUserExists            = errors.New("user with this name or email already exists")
	ErrUserNameTaken         = errors.New("user name taken")
	ErrInvalidPassword       = errors.New("password does not meet requirements")
	ErrInvalidNameOrEmail    = errors.New("invalid name or email")
	ErrInvalidNameOrPassword = errors.New("invalid name or password")
)

type Authorization interface {
	//returns ErrUserExists || ErrInvalidNameOrEmail || ErrInvalidPassword
	//returns other if any, nil if OK
	CheckUserSignUp(user structures.User) error

	//returns ErrInvalidNameOrPassword if invalid name or password,
	//returns err if other, nil if passwords are same
	CheckUserLogin(name, password string) error
}

//check Authorization implementation by Service struct in runtime

type AuthService struct {
	DBLAuth dblayer.Authorization
}

func newAuthService(db dblayer.Authorization) *AuthService {
	return &AuthService{
		DBLAuth: db,
	}
}

func (a *AuthService) CheckUserSignUp(user structures.User) error {
	userExists, err := a.DBLAuth.UserExists(user.Name, user.Email)
	if err != nil {
		return err
	} else if userExists {
		return ErrUserExists
	}

	validNameEmail := checkNameEmailIsValid(user.Name, user.Email)
	if !validNameEmail {
		return ErrInvalidNameOrEmail
	}
	validPassword := checkPasswordIsValid(user.Password, user.RePassword)
	if !validPassword {
		return ErrInvalidPassword
	}

	cryptPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}

	err = a.DBLAuth.AddUser(user.Name, user.Email, cryptPassword)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthService) CheckUserLogin(name, password string) error {
	dbPasswordHash, err := a.DBLAuth.GetUserLogin(name)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrInvalidNameOrPassword
		}
		return err
	}
	if !checkPasswordHash(password, dbPasswordHash) {
		return ErrInvalidNameOrPassword
	}
	return nil
}
