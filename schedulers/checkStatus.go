package schedulers

import (
	"stratus-core/ping"
	"time"

	"github.com/fatih/color"
)

func CheckStatus() {
	for {
		if len(projects) > 0 {
			for {
				startTime := time.Now()

				ping.CheckProjectStatus(projects)

				duration := time.Since(startTime)
				if duration<time.Minute*15 {
					time.Sleep(time.Minute * 15 - duration)
				}
			}
		} else {
			color.Red("Projects is empty, retrying after 3 minutes!!")
			time.Sleep(time.Minute * 3)
		}
	}
}
