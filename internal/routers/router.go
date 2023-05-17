package routers

import (
	// _ "post/docs"

	"hr/internal/middleware/jwt"
	"hr/internal/routers/api"
	v1 "hr/internal/routers/api/v1"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := r.Group("/auth")
	{
		auth.POST("register", api.Register)
		auth.POST("login", api.Login)
	}

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		employee := apiv1.Group("/employee")
		{
			employee.GET("", v1.GetEmployees)
			employee.GET("/:employeeId", v1.GetEmployee)
			employee.POST("", v1.CreateEmployee)
			employee.PATCH("/:employeeId", v1.UpdateEmployee)
			employee.DELETE("/:employeeId", v1.DeleteEmployee)
		}
	}

	return r
}
