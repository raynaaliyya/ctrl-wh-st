package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/raynaaliyya/ctrl-wh-st/models"
)

type WarehouseTeamRepo interface {
	GetAll() any
	GetById(id int) any
	GetByName(name string) (*models.WarehouseTeam, error)
	Create(newEmployee *models.WarehouseTeam) (*models.WarehouseTeam, error)
	Update(employee *models.WarehouseTeam) string
	Delete(id int) string
}

type warehouseTeamRepo struct {
	db *sql.DB
}

func NewWarehouseTeamRepo(db *sql.DB) WarehouseTeamRepo {
	repo := new(warehouseTeamRepo)
	repo.db = db

	return repo
}

func (r *warehouseTeamRepo) GetAll() any {
	var employees []models.WarehouseTeam

	query := "SELECT id, name, email, password, phone, photo FROM admin_wh"

	rows, err := r.db.Query(query)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var employee models.WarehouseTeam

		if err := rows.Scan(&employee.ID, &employee.Name, &employee.Email, &employee.Password, &employee.Phone, &employee.Photo); err != nil {
			log.Println(err)
		}

		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	if len(employees) == 0 {
		return "no data"
	}

	return employees
}

func (r *warehouseTeamRepo) GetById(id int) any {
	var employeeInDb models.WarehouseTeam

	query := "SELECT id, name, email, password, phone, photo FROM admin_wh WHERE id = $1"
	row := r.db.QueryRow(query, id)

	err := row.Scan(&employeeInDb.ID, &employeeInDb.Name, &employeeInDb.Email, &employeeInDb.Password, &employeeInDb.Phone, &employeeInDb.Photo)

	if err != nil {
		log.Println(err)
	}

	if employeeInDb.ID == 0 {
		return "employee not found"
	}

	return employeeInDb
}

func (r *warehouseTeamRepo) GetByName(name string) (*models.WarehouseTeam, error) {
	var employeeInDb models.WarehouseTeam

	query := "SELECT id, name, email, password, phone, photo FROM admin_wh WHERE name = $1"
	row := r.db.QueryRow(query, name)

	err := row.Scan(&employeeInDb.ID, &employeeInDb.Name, &employeeInDb.Email, &employeeInDb.Password, &employeeInDb.Phone, &employeeInDb.Photo)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("employee not found")
		}
		return nil, err
	}

	return &employeeInDb, nil
}

func (r *warehouseTeamRepo) Create(newEmployee *models.WarehouseTeam) (*models.WarehouseTeam, error) {
	query := "INSERT INTO admin_wh(id, name, email, password, phone, photo) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := r.db.Exec(query, newEmployee.ID, newEmployee.Name, newEmployee.Email, newEmployee.Password, newEmployee.Phone, newEmployee.Photo)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to create employee")
	}

	return newEmployee, nil
}

func (r *warehouseTeamRepo) Update(employee *models.WarehouseTeam) string {
	res, err := r.GetByName(employee.Name) //respon

	// jika tidak ada, return pesan
	if err != nil {
		return err.Error()
	}

	// jika ada maka update user
	query := "UPDATE admin_wh SET id = $1, name = $2, email = $3, password = $4, phone = $5, photo = $6 WHERE name = $7"
	_, err = r.db.Exec(query, employee.ID, employee.Name, employee.Email, employee.Password, employee.Phone, employee.Photo, res.Name)

	if err != nil {
		log.Println(err)
		return "failed to update employee"
	}

	return fmt.Sprintf("employee %s updated successfully", res.Name)
}

func (r *warehouseTeamRepo) Delete(id int) string {
	res := r.GetById(id)

	// jika tidak ada, return pesan
	if res == "employee not found" {
		return res.(string)
	}

	// jika ada, delete user
	query := "DELETE FROM admin_wh WHERE id = $1"
	_, err := r.db.Exec(query, id)

	if err != nil {
		log.Println(err)
		return "failed to delete employee"
	}

	return fmt.Sprintf("employee with id %d deleted successfully", id)
}
