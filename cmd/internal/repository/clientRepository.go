package repository

import (
	"database/sql"
	"fmt"

	"github.com/Prettyletto/service-notifier/cmd/internal/db"
	"github.com/Prettyletto/service-notifier/cmd/internal/model"
)

type ClientRepository interface {
	SaveClient(client *model.Client) error
	FindAllClients() ([]model.Client, error)
	FindClientById(id string) (*model.Client, error)
	UpdateClient(id string, client *model.Client) error
	DeleteClient(id string) error
}

type clientRepository struct {
	database          *db.DataBase
	companyRepository CompanyRepository
}

func NewClientRepo(database *db.DataBase, companyRepo CompanyRepository) ClientRepository {
	return &clientRepository{database: database, companyRepository: companyRepo}
}

func (r *clientRepository) SaveClient(client *model.Client) error {
	exists, exterr := r.companyRepository.Exists(client.CompanyID)
	if exterr != nil {
		return exterr
	}
	if !exists {
		return fmt.Errorf("company with id %s doesn't exists", client.CompanyID)
	}

	excStmt := `INSERT INTO client (id,name,dob,email,phone,address,company_id) VALUES(?,?,?,?,?,?,?)`

	_, err := r.database.DB.Exec(excStmt,
		client.ID,
		client.Name,
		client.DOB,
		client.Email,
		client.Phone,
		client.Address,
		client.CompanyID)

	if err != nil {
		return fmt.Errorf("Failed to insert client: %w", err)
	}

	return nil
}

func (r *clientRepository) FindAllClients() ([]model.Client, error) {
	rows, err := r.database.DB.Query(`SELECT id,name,dob,email,phone, address,company_id FROM client`)
	if err != nil {
		return nil, fmt.Errorf("error querying clients: %w", err)
	}
	defer rows.Close()

	var clients []model.Client

	for rows.Next() {
		var client model.Client
		if err := rows.Scan(&client.ID,
			&client.Name,
			&client.DOB,
			&client.Email,
			&client.Phone,
			&client.Address,
			&client.CompanyID); err != nil {

			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		clients = append(clients, client)
	}
	return clients, nil
}

func (r *clientRepository) FindClientById(id string) (*model.Client, error) {
	query := `SELECT id,name,dob,email,phone,address,company_id FROM client WHERE id = ?`
	row := r.database.DB.QueryRow(query, id)

	var client model.Client
	err := row.Scan(&client.ID, &client.Name, &client.DOB, &client.Email, &client.Phone, &client.Address, &client.CompanyID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no client found")
		}
		return nil, fmt.Errorf("error on query: %q", err)
	}
	return &client, nil
}

func (r *clientRepository) UpdateClient(id string, client *model.Client) error {
	query := `UPDATE client SET name = ?, dob = ?,email = ?, phone = ?, address = ?, company_id = ? WHERE id = ?`
	result, err := r.database.DB.Exec(query, client.Name, client.DOB, client.Email, client.Phone, client.Address, client.CompanyID, id)
	if err != nil {
		return fmt.Errorf("failed to update client")
	}

	updatedRows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve rows: %w", err)
	}
	if updatedRows == 0 {
		return fmt.Errorf("no client found with id: %s", id)
	}

	return nil
}

func (r *clientRepository) DeleteClient(id string) error {
	query := `DELETE FROM client WHERE id = ?`
	result, err := r.database.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete client")
	}

	deletedRows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to delete rows: %w", err)
	}
	if deletedRows == 0 {
		return fmt.Errorf("no client found with id %s", id)
	}

	return nil
}
