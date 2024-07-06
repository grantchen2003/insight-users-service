package main

import (
	"fmt"
	"log"
	"os"

	"github.com/grantchen2003/insight/users/internal/config"
	databasePackage "github.com/grantchen2003/insight/users/internal/database"
	serverPackage "github.com/grantchen2003/insight/users/internal/server"
)

func main() {
	env := os.Getenv("ENV")
	log.Printf("ENV=%s", env)
	if err := config.LoadEnvVars(env); err != nil {
		log.Fatalf("failed to load env vars")
	}

	database := databasePackage.GetSingletonInstance()
	if err := database.Connect(); err != nil {
		log.Fatal("failed to connect to database")
	}
	defer database.Close()

	address := fmt.Sprintf("%s:%s", os.Getenv("DOMAIN"), os.Getenv("PORT"))
	server := serverPackage.NewServer()
	if err := server.Start(address); err != nil {
		log.Fatalf("failed to start server")
	}
}
