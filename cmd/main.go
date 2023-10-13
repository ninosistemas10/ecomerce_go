package main

import (
	"log"
	"os"

	"github.com/ninosistemas10/ecommerce/infrastructure/handler"
	"github.com/ninosistemas10/ecommerce/infrastructure/handler/response"
)

func main() {
	err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	err = validateEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	app := newHTTP(response.HTTPErrorHandler)


	dbPool, err := newDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	handler.InitRoutes(app, dbPool)

	err = app.Listen(":" + os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
