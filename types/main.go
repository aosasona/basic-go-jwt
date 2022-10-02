package types

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type User struct {
	UUID            string `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	Password        string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirm_password,omitempty"`
}

type Note struct {
	UUID     string `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	UserUUID string `json:"user_id"`
}

func (u *User) Validate() (bool, error) {

	emailRegex := `^([a-zA-Z0-9_\-\.]+)@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.)|(([a-zA-Z0-9\-]+\.)+))([a-zA-Z]{2,4}|[0-9]{1,3})(\]?)$`

	if u.FirstName == "" || len(u.FirstName) < 2 || len(u.FirstName) > 100 {
		return false, errors.New("first name is invalid or missing")
	}
	if u.LastName == "" || len(u.LastName) < 2 || len(u.LastName) > 100 {
		return false, errors.New("last name is invalid or missing")
	}
	if u.Email == "" || len(u.Email) < 6 || len(u.Email) > 100 || !regexp.MustCompile(emailRegex).MatchString(u.Email) {
		return false, errors.New("e-mail is invalid or missing")
	}
	if u.Password != u.ConfirmPassword {
		return false, errors.New("passwords do not match")
	}
	if len(u.Password) < 6 {
		return false, errors.New("password must be at least 6 characters")
	}
	return true, nil
}

func (u *User) ValidateLoginData() (bool, error) {
	emailRegex := `^([a-zA-Z0-9_\-\.]+)@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.)|(([a-zA-Z0-9\-]+\.)+))([a-zA-Z]{2,4}|[0-9]{1,3})(\]?)$`
	if u.Email == "" || len(u.Email) < 6 || len(u.Email) > 100 || !regexp.MustCompile(emailRegex).MatchString(u.Email) {
		return false, errors.New("e-mail is invalid or missing")
	}
	if u.Password == "" {
		return false, errors.New("password is required")
	}
	if len(u.Password) < 6 {
		return false, errors.New("password must be at least 6 characters")
	}
	return true, nil
}

func (u *User) HashPassword() (string, error) {
	bytePassword := []byte(u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("something went wrong")
	}
	return string(hashedPassword), nil
}

func (u *User) ComparePassword(hashedPassword string) error {
	bytePassword := []byte(u.Password)
	byteHashedPassword := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
	if err != nil {
		return err
	}
	return nil
}

func (n *Note) Validate() error {
	if n.Title == "" {
		return errors.New("title is required")
	}
	if len(n.Title) < 2 {
		return errors.New("title must be at least 2 characters")
	}
	if len(n.Title) > 100 {
		return errors.New("title is too long (max 100 characters)")
	}
	if n.Body == "" {
		return errors.New("body is required")
	}
	if len(n.Body) < 2 {
		return errors.New("body must be at least 2 characters")
	}
	if len(n.Body) > 1499 {
		return errors.New("body is too long (max 1000 characters)")
	}
	if n.UserUUID == "" {
		return errors.New("user is required")
	}
	return nil
}
