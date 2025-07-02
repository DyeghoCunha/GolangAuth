package routes

import (
	"fmt"
	"net/http"

	"github.com/dyeghocunha/golang-auth/controller"
)

func SetupRoutes() {
	http.HandleFunc("/register", controller.RegisterHandler)
	http.HandleFunc("/enable-2fa", controller.Enable2FAHandler)
	http.HandleFunc("/generate-qr", controller.GenerateQRCodeHandler)
	http.HandleFunc("/verify-2fa", controller.Verify2FAHandler)
}

func UserRegistryHandler() {
	http.HandleFunc("/user", controller.GetUserHandler)
}
func HealthCheckHandler() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Auth service está online!")
	})

}
