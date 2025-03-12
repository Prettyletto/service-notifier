package repository

import (
	"database/sql"
	"fmt"

	"github.com/Prettyletto/service-notifier/cmd/internal/db"
	"github.com/Prettyletto/service-notifier/cmd/internal/model"
)

type ServiceRepository interface {
	SaveService(service *model.Service) error
	FindAllServices() ([]model.Service, error)
	FindServiceById(id string) (*model.Service, error)
	UpdateService(id string, service *model.Service) error
	DeleteService(id string) error
}

type serviceRepository struct {
	database *db.DataBase
}

func NewServiceRepo(database *db.DataBase) ServiceRepository {
	return &serviceRepository{database: database}
}

func (r *serviceRepository) SaveService(service *model.Service) error {
	excStmt := `INSERT INTO service (id,name,due_days,priority,max_reschedules,company_id) VALUES(?,?,?,?,?,?)`

	_, err := r.database.DB.Exec(excStmt,
		service.ID,
		service.Name,
		service.DueDays,
		service.Priority,
		service.MaxReschedules,
		service.CompanyID)

	if err != nil {
		return fmt.Errorf("failed to insert service: %w", err)
	}

	return nil
}

func (r *serviceRepository) FindAllServices() ([]model.Service, error) {

	rows, err := r.database.DB.Query(`SELECT id,name,due_days,priority,max_reschedules,company_id FROM service`)
	if err != nil {
		return nil, fmt.Errorf("error querying services: %w", err)
	}
	defer rows.Close()

	var services []model.Service

	for rows.Next() {
		var service model.Service
		if err := rows.Scan(&service.ID,
			&service.Name,
			&service.DueDays,
			&service.Priority,
			&service.MaxReschedules,
			&service.CompanyID); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		services = append(services, service)
	}
	return services, nil
}

func (r *serviceRepository) FindServiceById(id string) (*model.Service, error) {
	query := `SELECT id,name,due_days,priority,max_reschedules,company_id FROM service WHERE id = ?`
	row := r.database.DB.QueryRow(query, id)

	var service model.Service
	err := row.Scan(&service.ID,
		&service.Name,
		&service.DueDays,
		&service.Priority,
		&service.MaxReschedules,
		&service.CompanyID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no service found")
		}
		return nil, fmt.Errorf("error on query: %q", err)
	}

	return &service, nil
}

func (r *serviceRepository) UpdateService(id string, service *model.Service) error {
	updateStmt := `
	UPDATE service SET 
	name = ?,
	due_days = ?,
	priority = ?,
	max_reschedules = ?,
	company_id = ? 
	WHERE ID = ?`
	result, err := r.database.DB.Exec(updateStmt, service.Name, service.DueDays, service.Priority, service.MaxReschedules, service.CompanyID, id)
	if err != nil {
		return fmt.Errorf("failed to update service")
	}

	updatedRows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve rows from database: %w", err)
	}
	if updatedRows == 0 {
		return fmt.Errorf("no service found with id: %s", id)
	}

	return nil
}

func (r *serviceRepository) DeleteService(id string) error {
	query := `DELETE FROM service WHERE id = ?`
	result, err := r.database.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete service")
	}

	deletedRows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to delete rows: %w", err)
	}
	if deletedRows == 0 {
		return fmt.Errorf("no service found with id %s", id)
	}

	return nil
}
