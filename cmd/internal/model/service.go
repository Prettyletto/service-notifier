package model

import (
	"fmt"

	"github.com/google/uuid"
)

type Service struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	DueDays        int    `json:"due_days"`
	Priority       int    `json:"priority"`
	MaxReschedules int    `json:"max_reschedules"`
	CompanyID      string `json:"company_id"`
}

func (s *Service) Validate() error {

	if s.Name == "" {
		return fmt.Errorf("service name cannot be empty")
	}
	if s.DueDays < 0 {
		return fmt.Errorf("service dues days cannot be less than 0")
	}
	if s.Priority < 0 {
		return fmt.Errorf("service priority cannot be less than 0")
	}
	if s.MaxReschedules < 0 {
		return fmt.Errorf("service reschedules cannot be less than 0")
	}
	if s.CompanyID == "" {
		return fmt.Errorf("service company id cannot be empty")
	}

	return nil
}

func (s *Service) GenerateID() {
	s.ID = uuid.New().String()
}
