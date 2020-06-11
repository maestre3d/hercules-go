package domain

import (
	"errors"
	"strings"
	"time"

	gonanoid "github.com/matoous/go-nanoid"
)

type Client struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	LastName   string    `json:"last_name"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	Active     bool      `json:"active"`
}

func NewClient(name, lastName string) *Client {
	id, err := gonanoid.Nanoid(16)
	if err != nil {
		id = ""
	}

	return &Client{
		ID:         id,
		Name:       sanitizeName(name),
		LastName:   sanitizeName(lastName),
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Active:     true,
	}
}

// sanitizeName Clean name, output from 'lorem' -> 'Lorem'
func sanitizeName(field string) string {
	if len(field) == 0 {
		return ""
	}

	fieldSlice := strings.Split(field, "")
	field = ""

	fieldSlice[0] = strings.ToUpper(fieldSlice[0])
	for _, f := range fieldSlice {
		field += f
	}

	return field
}

func (c Client) IsValid() error {
	switch {
	case len(c.Name) == 0 || len(c.Name) > 255:
		return errors.New("invalid name")
	case len(c.LastName) == 0 || len(c.LastName) > 255:
		return errors.New("invalid last name")
	}

	return nil
}
