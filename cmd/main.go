package main

import (
	"log"
	"stratus-core/database"
	"stratus-core/schedulers"

	"github.com/fatih/color"
)

func main() {
	err := database.InitMongoDB()
	if err != nil {
		color.Red("Error in mongo connection")
		log.Fatal(err)
	}

	color.Green("Starting stratus health checker...")

	schedulers.UpdateProjects()
	schedulers.CheckStatus()
}