package schedulers

import (
	"stratus-core/database"
	"stratus-core/models"
	"sync"
	"time"

	"github.com/fatih/color"
)

var projects []models.Project
var mu sync.Mutex

func UpdateProjects() {
	for {
		color.Green("Reading DB...")
		success := false
		for success==false {
			newProjects, err := database.GetProjects()
			if err!=nil {
				color.Magenta("Unable to read DB, retrying in 3 minutes!!")
				time.Sleep(time.Minute * 3)
			} else {
				mu.Lock()
				projects = newProjects
				mu.Unlock()
				success = true
			}
		}
		time.Sleep(time.Hour * 24)
	}
}