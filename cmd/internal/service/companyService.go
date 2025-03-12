package service

import (
	"github.com/Prettyletto/service-notifier/cmd/internal/model"
	"github.com/Prettyletto/service-notifier/cmd/internal/repository"
)

type CompanyService interface {
	CreateCompany(company *model.Company) error
	RetrieveAllCompanies() ([]model.Company, error)
	RetrieveCompanyById(id string) (*model.Company, error)
	UpdateCompany(id string, company *model.Company) (*model.Company, error)
	DeleteCompany(id string) error
}

type companyService struct {
	repo repository.CompanyRepository
}

func NewCompanyService(repo repository.CompanyRepository) CompanyService {
	return &companyService{repo: repo}
}

func (s *companyService) CreateCompany(company *model.Company) error {

	if err := company.Validate(); err != nil {
		return err
	}

	company.GenerateID()

	return s.repo.SaveCompany(company)
}

func (s *companyService) RetrieveAllCompanies() ([]model.Company, error) {
	return s.repo.FindAllCompanies()
}

func (s *companyService) RetrieveCompanyById(id string) (*model.Company, error) {

	return s.repo.FindCompanyById(id)
}

func (s *companyService) UpdateCompany(id string, company *model.Company) (*model.Company, error) {
	if err := company.Validate(); err != nil {
		return nil, err
	}

	err := s.repo.UpdateCompany(id, company)
	if err != nil {
		return nil, err
	}
	return company, err
}

func (s *companyService) DeleteCompany(id string) error {
	return s.repo.DeleteCompany(id)
}
