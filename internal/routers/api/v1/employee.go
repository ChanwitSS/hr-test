package v1

import (
	"hr/internal/models"
	"hr/internal/services"
	"hr/pkg/app"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetEmployees(c *gin.Context) {
	var (
		appG          = app.Gin{C: c}
		queryEmployee = models.QueryEmployee{}
	)

	if err := c.ShouldBindQuery(&queryEmployee); err != nil {
		appG.Response(http.StatusBadRequest, strings.Split(err.Error(), "\n"))
		return
	}

	employees, err := services.GetEmployees(queryEmployee)
	if err != nil {
		appG.Response(http.StatusInternalServerError, err)
		return
	}

	appG.Response(http.StatusOK, employees)
}

func GetEmployee(c *gin.Context) {
	var (
		appG       = app.Gin{C: c}
		employeeId = appG.C.Param("employeeId")
	)

	employee, err := services.GetEmployee(employeeId)
	if err != nil {
		appG.Response(http.StatusInternalServerError, err)
		return
	}
	appG.Response(http.StatusOK, employee)
}

func CreateEmployee(c *gin.Context) {
	var (
		appG           = app.Gin{C: c}
		createEmployee = models.Employee{}
	)

	if err := appG.C.ShouldBindJSON(&createEmployee); err != nil {
		appG.Response(http.StatusBadRequest, strings.Split(err.Error(), "\n"))
		return
	}

	employee, err := services.CreateEmployee(createEmployee)
	if err != nil {
		appG.Response(http.StatusInternalServerError, err)
		return
	}
	appG.Response(http.StatusOK, employee)
}

func UpdateEmployee(c *gin.Context) {
	var (
		appG           = app.Gin{C: c}
		employeeId     = appG.C.Param("employeeId")
		updateEmployee = models.Employee{}
	)

	if err := appG.C.ShouldBindJSON(&updateEmployee); err != nil {
		appG.Response(http.StatusBadRequest, strings.Split(err.Error(), "\n"))
		return
	}

	employee, err := services.UpdateEmployee(employeeId, updateEmployee)
	if err != nil {
		appG.Response(http.StatusInternalServerError, err)
		return
	}
	appG.Response(http.StatusOK, employee)
}

func DeleteEmployee(c *gin.Context) {
	var (
		appG       = app.Gin{C: c}
		employeeId = appG.C.Param("employeeId")
	)

	err := services.DeleteEmployee(employeeId)
	if err != nil {
		appG.Response(http.StatusInternalServerError, err)
		return
	}
	appG.Response(http.StatusOK, nil)
}
