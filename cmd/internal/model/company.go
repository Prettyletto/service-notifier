package model

import (
	"fmt"

	"github.com/google/uuid"
)

type Company struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c *Company) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("Company name cannot be empty")
	}
	return nil
}

func (c *Company) GenerateID() {
	c.ID = uuid.New().String()
}
