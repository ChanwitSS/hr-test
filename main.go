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
	// a := []int{1, 2, 3}
	// b := []int{2, 3, 4}
	// TestFunction(a, b)
	gin.SetMode(os.Getenv("GIN_MODE"))
	routersInit := routers.InitRouter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: routersInit,
	}

	server.ListenAndServe()
}

func TestFunction(a []int, b []int) (*[]int, *[]int) {
	visited := make(map[int]bool, 0)
	var res1 = []int{}
	var res2 = []int{}

	for i := 0; i < len(a); i++ {
		if visited[a[i]] == true {
			res2 = append(res2[:i], res2[i+1:]...)
		} else {
			res1 = append(res1, a[i])
			res2 = append(res2, a[i])
			visited[a[i]] = true
		}
	}
	for i := 0; i < len(b); i++ {
		if visited[b[i]] == true {
			res2 = append(res2[:i], res2[i+1:]...)
		} else {
			res1 = append(res1, b[i])
			res2 = append(res2, b[i])
			visited[b[i]] = true
		}
	}

	fmt.Println(res1, res2)
	return &a, &b
}
