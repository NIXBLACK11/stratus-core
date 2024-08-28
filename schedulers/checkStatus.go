package schedulers

import (
	"stratus-core/models"
	"stratus-core/ping"
	"time"

	"github.com/fatih/color"
)

func CheckStatus() {
	for {
		mu.Lock()
		if len(projects) > 0 {
			// Make a copy of projects or access its length outside the critical section
			localProjects := make([]models.Project, len(projects))
			copy(localProjects, projects)
			mu.Unlock() // Unlock as early as possible

			startTime := time.Now()

			ping.CheckProjectStatus(localProjects)

			duration := time.Since(startTime)
			if duration<time.Minute*15 {
				time.Sleep(time.Minute * 15 - duration)
			}
		} else {
			mu.Unlock() // Unlock before the sleep
			color.Red("Projects is empty, retrying after 3 minutes!!")
			time.Sleep(time.Minute * 3)
		}
	}
}