package ping

import (
	"stratus-core/models"
	"stratus-core/utils"
)

const maxRetries = 3

func CheckProjectStatus(projects []models.Project) {
	for i := len(projects) - 1; i >= 0; i-- {
		project := &projects[i]
		alertTriggers := project.AlertTriggers

		Triggered := false
		for _, alertTrigger := range alertTriggers {
			triggered := CheckStatus(alertTrigger.SiteUrl, alertTrigger.AlertType)

			if len(triggered) > 0 {
				Triggered = true
			}

			for _, trigger := range triggered {
				go utils.SendMail(
					project.UserName,
					project.ProjectName,
					alertTrigger.SiteName,
					alertTrigger.SiteUrl,
					trigger,
				)
			}
		}
		if Triggered==true {
			project.Tries--;
		}
		if project.Tries < 0 {
			projects = append(projects[:i], projects[i+1:]...)
		}
	}
}