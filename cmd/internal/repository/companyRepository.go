package repository

import (
	"database/sql"
	"fmt"

	"github.com/Prettyletto/service-notifier/cmd/internal/db"
	"github.com/Prettyletto/service-notifier/cmd/internal/model"
)

type CompanyRepository interface {
	Exists(id string) (bool, error)
	SaveCompany(company *model.Company) error
	FindAllCompanies() ([]model.Company, error)
	FindCompanyById(id string) (*model.Company, error)
	UpdateCompany(id string, company *model.Company) error
	DeleteCompany(id string) error
}

type companyRepository struct {
	database *db.DataBase
}

func NewCompanyRepo(database *db.DataBase) CompanyRepository {
	return &companyRepository{database: database}
}

func (r *companyRepository) Exists(id string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM company WHERE id = ?)`
	var exists bool
	err := r.database.DB.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed do query the company id: %w", err)
	}
	return exists, nil
}

func (r *companyRepository) SaveCompany(company *model.Company) error {
	excStmt := `INSERT INTO company (id,name) VALUES(?,?)`

	_, err := r.database.DB.Exec(excStmt,
		company.ID,
		company.Name)

	if err != nil {
		return fmt.Errorf("Failed to insert company: %w", err)
	}

	return nil
}

func (r *companyRepository) FindAllCompanies() ([]model.Company, error) {

	rows, err := r.database.DB.Query(`SELECT id,name FROM company`)
	if err != nil {
		return nil, fmt.Errorf("error querying companies: %w", err)
	}
	defer rows.Close()

	var companies []model.Company

	for rows.Next() {
		var company model.Company
		if err := rows.Scan(&company.ID, &company.Name); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		companies = append(companies, company)
	}
	return companies, nil
}

func (r *companyRepository) FindCompanyById(id string) (*model.Company, error) {
	query := `SELECT id,name FROM company WHERE id = ?`
	row := r.database.DB.QueryRow(query, id)

	var company model.Company
	err := row.Scan(&company.ID, &company.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no company found")
		}
		return nil, fmt.Errorf("error on query: %q", err)
	}

	return &company, nil
}

func (r *companyRepository) UpdateCompany(id string, company *model.Company) error {
	updateStmt := `UPDATE company SET name = ? WHERE ID = ?`
	result, err := r.database.DB.Exec(updateStmt, company.Name, id)
	if err != nil {
		return fmt.Errorf("failed to update company")
	}

	updatedRows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve rows from database: %w", err)
	}
	if updatedRows == 0 {
		return fmt.Errorf("no company found with id: %s", id)
	}

	return nil
}

func (r *companyRepository) DeleteCompany(id string) error {
	query := `DELETE FROM company WHERE id = ?`
	result, err := r.database.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete company")
	}

	deletedRows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to delete rows: %w", err)
	}
	if deletedRows == 0 {
		return fmt.Errorf("no company found with id %s", id)
	}

	return nil
}
