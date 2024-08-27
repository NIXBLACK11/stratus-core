package ping

import (
	"stratus-core/models"
	"stratus-core/utils"
	"time"

	"github.com/fatih/color"
)

const maxRetries = 3

func CheckProjectStatus(projects []models.Project) {
	for _, project := range projects {
		alertTriggers := project.AlertTriggers

		for _, alertTrigger := range alertTriggers {
			triggered := CheckStatus(alertTrigger.SiteUrl, alertTrigger.AlertType)

			if len(triggered) > 0 {
				successChan := make(chan bool, len(triggered))

				for _, trigger := range triggered {
					go func(trigger string) {
						successChan <- sendMailWithRetries(project, alertTrigger, trigger)
					}(trigger)
				}

				failedTriggers := []string{}
				for i := 0; i < len(triggered); i++ {
					if !<-successChan {
						failedTriggers = append(failedTriggers, triggered[i])
					}
				}

				if len(failedTriggers) > 0 {
					color.Red("Some emails failed after multiple attempts!!!")
				}
			}
		}
	}
}

func sendMailWithRetries(project models.Project, alertTrigger models.AlertTrigger, trigger string) bool {
	for attempt := 1; attempt <= maxRetries; attempt++ {
		success := utils.SendMail(
			project.UserName,
			project.ProjectName,
			alertTrigger.SiteName,
			alertTrigger.SiteUrl,
			trigger,
		)
		if success {
			return true
		}
		if attempt < maxRetries {
			time.Sleep(time.Second * 2)
		}
	}
	return false
}
