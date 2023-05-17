package main

import (
	"fmt"
	"hr/internal/models"
	"hr/internal/routers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
	models.Setup()
}

func main() {
	gin.SetMode(os.Getenv("GIN_MODE"))
	routersInit := routers.InitRouter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: routersInit,
	}

	server.ListenAndServe()

	// a := []string{"b", "c", "d"}
	// b := []string{"a", "c", "d"}
	// TestFunction(a, b)
}

func TestFunction(a []string, b []string) (*[]string, *[]string) {
	visited := make(map[string]bool, 0)
	var res1 = []string{}
	var res2 = []string{}

	for i := 0; i < len(a); i++ {
		if visited[a[i]] {
			res2 = append(res2[:i], res2[i+1:]...)
		} else {
			res1 = append(res1, a[i])
			res2 = append(res2, a[i])
			visited[a[i]] = true
		}
	}
	for i := 0; i < len(b); i++ {
		if visited[b[i]] {
			res2 = append(res2[:i], res2[i+1:]...)
		} else {
			res1 = append(res1, b[i])
			res2 = append(res2, b[i])
			visited[b[i]] = true
		}
	}

	fmt.Println(res1, res2)
	return &res1, &res2
}
