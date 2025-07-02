package routes

import (
	"fmt"
	"net/http"

	"github.com/dyeghocunha/golang-auth/controller"
	"github.com/dyeghocunha/golang-auth/middleware"
)

func SetupRoutes() {

	//	http.HandleFunc("/enable-2fa", controller.Enable2FAHandler)
	//http.HandleFunc("/verify-2fa", controller.Verify2FAHandler)

	http.HandleFunc("/generate-qr", controller.GenerateQRCodeHandler)

	http.Handle("/minha-rota", middleware.JWTAuthMiddleware(http.HandlerFunc(controller.ProfileHandler)))
	http.Handle("/rota-segura", middleware.JWTAuthMiddlewareWith2FA(http.HandlerFunc(controller.SensitiveHandler)))

	http.HandleFunc("/register", controller.RegisterHandler)
	http.Handle("/profile", middleware.JWTAuthMiddleware(http.HandlerFunc(controller.ProfileHandler)))

	http.Handle("/sensitive-data", middleware.JWTAuthMiddlewareWith2FA(http.HandlerFunc(controller.SensitiveHandler)))

	http.Handle("/enable-2fa", middleware.JWTAuthMiddleware(http.HandlerFunc(controller.Enable2FAHandler)))
	http.Handle("/verify-2fa", middleware.JWTAuthMiddleware(http.HandlerFunc(controller.Verify2FAHandler)))

}

func UserRegistryHandler() {
	http.HandleFunc("/user", controller.GetUserHandler)
}
func HealthCheckHandler() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Auth service está online!")
	})

}
