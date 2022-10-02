package types

import (
	"errors"
)

type Note struct {
	UUID     string `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	UserUUID string `json:"user_id"`
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
		return errors.New("body is too long (max 1499 characters)")
	}
	if n.UserUUID == "" {
		return errors.New("user is required")
	}
	return nil
}
