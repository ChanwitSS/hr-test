package services

import (
	// "hr/internal/enums"
	"hr/internal/models"
	"hr/pkg/app/util"
)

func GetEmployees(query models.QueryEmployee) (*[]models.Employee, error) {
	employees, err := models.FindEmployees(query)
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func GetEmployee(employeeId string) (*models.Employee, error) {
	employee, err := models.FindEmployee(employeeId)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func CreateEmployee(createEmployee models.Employee) (*models.Employee, error) {
	result, err := models.CreateEmployee(models.Employee{
		FirstName: createEmployee.FirstName,
		LastName:  createEmployee.LastName,
		Email:     createEmployee.Email,
		Tel:       createEmployee.Tel,

		CreatedAt: util.GetLocalTime("Bangkok"),
		UpdatedAt: util.GetLocalTime("Bangkok"),
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateEmployee(employeeId string, createEmployee models.Employee) (*models.Employee, error) {
	employee, err := models.UpdateEmployee(employeeId, createEmployee)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func DeleteEmployee(employeeId string) error {
	err := models.DeleteEmployee(employeeId)
	if err != nil {
		return err
	}
	return nil
}
