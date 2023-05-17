package api

import (
	"hr/internal/middleware/jwt"
	"hr/internal/models"
	"hr/internal/services"
	"hr/pkg/app"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	Username string `valid:"Required; MaxSize(50)" json:"username"`
	Password string `valid:"Required; MaxSize(50)" json:"password"`
	Email    string `valid:"Required; MaxSize(50)" json:"email"`
}

type AuthResponse struct {
	Token *string
}

func Register(c *gin.Context) {
	var (
		appG       = app.Gin{C: c}
		createUser = models.User{}
	)

	if err := c.ShouldBindJSON(&createUser); err != nil {
		appG.Response(http.StatusBadRequest, strings.Split(err.Error(), "\n"))
		return
	}

	user, err := services.Register(createUser)
	if err != nil {
		appG.Response(http.StatusInternalServerError, err)
		return
	}

	appG.C.JSON(http.StatusOK, user)
}

func Login(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		auth = Auth{}
	)

	if err := c.ShouldBindJSON(&auth); err != nil {
		appG.Response(http.StatusBadRequest, strings.Split(err.Error(), "\n"))
		return
	}

	_, err := services.Login(auth.Username, auth.Password, auth.Email)
	if err != nil {
		appG.Response(http.StatusBadRequest, err)
		return
	}

	token, err := jwt.GenerateToken()
	if err != nil {
		appG.Response(http.StatusInternalServerError, err)
		return
	}

	appG.Response(http.StatusOK, &AuthResponse{
		Token: token,
	})
}

// func GetAuth(c *gin.Context) {
// 	appG := app.Gin{C: c}
// 	valid := validation.Validation{}

// 	username := c.PostForm("username")
// 	password := c.PostForm("password")

// 	a := auth{Username: username, Password: password}
// 	ok, _ := valid.Valid(&a)

// 	if !ok {
// 		app.MarkErrors(valid.Errors)
// 		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
// 		return
// 	}

// 	authService := auth_service.Auth{Username: username, Password: password}
// 	isExist, err := authService.Check()
// 	if err != nil {
// 		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
// 		return
// 	}

// 	if !isExist {
// 		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
// 		return
// 	}

// 	token, err := util.GenerateToken(username, password)
// 	if err != nil {
// 		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
// 		return
// 	}

// 	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
// 		"token": token,
// 	})
// }
