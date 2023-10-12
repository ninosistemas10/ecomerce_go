package main

import (
	"log"
	"os"

	
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

	app := newHTTP()

	
	
	err = app.Listen(":" + os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
