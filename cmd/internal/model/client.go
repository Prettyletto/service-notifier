package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	DOB       string `json:"dob"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	CompanyID string `json:"company_id"`
}

func (c *Client) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("Name of the client cannot be empty")
	}
	if c.DOB == "" {
		return fmt.Errorf("DOB  of the client cannot be empty")
	}
	if c.Email == "" {
		return fmt.Errorf("Email of the client cannot be empty")
	}
	if c.Phone == "" {
		return fmt.Errorf("Phone of the client cannot be empty")
	}
	if c.Address == "" {
		return fmt.Errorf("Address of the client cannot be empty")
	}
	if c.CompanyID == "" {
		return fmt.Errorf("CompanyID of the client cannot be empty")
	}
	return nil
}

func (c *Client) ParseDOB() error {
	input := "01-02-2006"
	output := "2006-01-02"

	parseDOB, err := time.Parse(input, c.DOB)
	if err != nil {
		return fmt.Errorf("invalid DOB format, expecting MM-DD-YYYY")
	}

	c.DOB = parseDOB.Format(output)
	return nil
}

func (c *Client) GenerateID() {
	c.ID = uuid.New().String()
}
