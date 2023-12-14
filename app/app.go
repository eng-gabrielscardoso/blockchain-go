package app

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

// This function generates a Urfave CLI application
func Generate() (app *cli.App) {
	app = cli.NewApp()

	app.Name = os.Getenv("APP_NAME")

	app.Author = os.Getenv("APP_AUTHOR")

	app.Email = os.Getenv("APP_EMAIL")

	app.Usage = os.Getenv("APP_USAGE")

	app.Description = os.Getenv("APP_DESCRIPTION")

	app.Commands = []cli.Command{
		CreateBlockchainCommand(),
	}

	return app
}

// This function bootstraps the application
func Bootstrap() {
	if error := godotenv.Load(); error != nil {
		log.Fatal(error)
	}

	if error := Generate().Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
