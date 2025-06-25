package routes

import (
	"fmt"
	"net/http"

	"github.com/dyeghocunha/golang-auth/controller"
)

func SetupRoutes() {
	http.HandleFunc("/register", controller.RegisterHandler)
}

func HealthCheckHandler() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Auth service está online!")
	})

}
