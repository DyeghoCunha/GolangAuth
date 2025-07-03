package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dyeghocunha/golang-auth/db"
	"github.com/dyeghocunha/golang-auth/routes"
)

func main() {

	err := db.Connect()
	if err != nil {
		log.Fatal("❌ Erro ao conectar no banco:", err)
	}
	routes.HealthCheckHandler()
	routes.SetupRoutes()
	routes.UserRegistryHandler()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not set
	}
	log.Println("🚀 Servidor rodando na porta:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
	log.Println("JWT_SECRET carregado:", os.Getenv("JWT_SECRET"))

}
