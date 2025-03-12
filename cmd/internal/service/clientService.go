package service

import (
	"github.com/Prettyletto/service-notifier/cmd/internal/model"
	"github.com/Prettyletto/service-notifier/cmd/internal/repository"
)

type ClientService interface {
	CreateClient(client *model.Client) error
	RetrieveAllClients() ([]model.Client, error)
	RetrieveClientById(id string) (*model.Client, error)
	UpdateClient(id string, client *model.Client) (*model.Client, error)
	DeleteClient(id string) error
}

type clientService struct {
	repo repository.ClientRepository
}

func NewClientService(repo repository.ClientRepository) ClientService {
	return &clientService{repo: repo}
}

func (s *clientService) CreateClient(client *model.Client) error {

	if err := client.Validate(); err != nil {
		return err
	}

	client.GenerateID()

	if err := client.ParseDOB(); err != nil {
		return err
	}

	return s.repo.SaveClient(client)
}

func (s *clientService) RetrieveAllClients() ([]model.Client, error) {
	return s.repo.FindAllClients()
}

func (s *clientService) RetrieveClientById(id string) (*model.Client, error) {

	return s.repo.FindClientById(id)
}

func (s *clientService) UpdateClient(id string, client *model.Client) (*model.Client, error) {
	if err := client.Validate(); err != nil {
		return nil, err
	}

	if err := client.ParseDOB(); err != nil {
		return nil, err
	}

	err := s.repo.UpdateClient(id, client)
	if err != nil {
		return nil, err
	}

	return client, err
}

func (s *clientService) DeleteClient(id string) error {
	return s.repo.DeleteClient(id)
}
