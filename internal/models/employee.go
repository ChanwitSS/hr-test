package models

import (
	"hr/pkg/app"
	"time"
)

type Employee struct {
	EmployeeId int    `gorm:"primary_key" json:"employee_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Tel        string `json:"tel"`
	Gender     string `json:"gender"`
	Position   string `json:"position"`
	Level      string `json:"level"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type QueryEmployee struct {
	app.Query
}

func FindEmployees(query QueryEmployee) (*[]Employee, error) {
	var (
		employees []Employee
		limit     = query.Take
		offset    = (query.Page - 1) * query.Take
	)

	err := db.
		Table("employees").
		Order("created_at DESC").
		Find(&employees).
		Offset(offset).
		Limit(limit).
		Error

	if err != nil {
		return nil, err
	}

	return &employees, nil
}

func FindEmployee(employeeId string) (*Employee, error) {
	var employee Employee
	err := db.
		Table("employees").
		Where("employees.employee_id = ?", employeeId).
		First(&employee).
		Error

	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func CreateEmployee(employee Employee) (*Employee, error) {
	if err := db.Table("employees").Create(&employee).Error; err != nil {
		return nil, err
	}

	return &employee, nil
}

func UpdateEmployee(employeeId string, employee Employee) (*Employee, error) {
	if err := db.Table("employees").Where("employee_id = ?", employeeId).Updates(&employee).Error; err != nil {
		return nil, err
	}

	return &employee, nil
}

func DeleteEmployee(id string) error {
	if err := db.Table("employees").Where("employee_id = ?", id).Delete(&Employee{}).Error; err != nil {
		return err
	}

	return nil
}
