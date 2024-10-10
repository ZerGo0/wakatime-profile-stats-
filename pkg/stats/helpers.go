package stats

import (
	"fmt"

	"github.com/google/go-github/v65/github"
	"github.com/samber/lo"
	"github.com/user/wakatime-profile-stats/pkg/wakatime"
)

const (
	yearInSeconds   = 31536000
	monthInSeconds  = 2592000
	dayInSeconds    = 86400
	hourInSeconds   = 3600
	minuteInSeconds = 60

	maxProjectNameLength = 25
)

func calculateWorkTime(stats *wakatime.WakaStats, githubRepos []*github.Repository) {
	privateProjects := make([]string, 0)

	privateWorkProjectsTotal := 0
projectLoop:
	for _, project := range stats.Data.Projects {
		for _, repo := range githubRepos {
			if project.Name == *repo.Name {
				continue projectLoop
			}
		}

		privateWorkProjectsTotal += int(project.TotalSeconds)
		privateProjects = append(privateProjects, project.Name)
	}

	// Remove private projects from stats
	stats.Data.Projects = lo.Filter(stats.Data.Projects, func(p wakatime.Projects, _ int) bool {
		_, found := lo.Find(privateProjects, func(project string) bool {
			return project == p.Name
		})

		return !found
	})

	stats.Data.Projects = append(stats.Data.Projects, wakatime.Projects{
		Name:         "Work/Private Projects",
		Text:         secondsToHumanReadable(privateWorkProjectsTotal),
		TotalSeconds: float64(privateWorkProjectsTotal),
	})
}

func secondsToHumanReadable(privateWorkProjectsTotal int) string {
	years := privateWorkProjectsTotal / yearInSeconds
	months := (privateWorkProjectsTotal % yearInSeconds) / monthInSeconds
	days := ((privateWorkProjectsTotal % yearInSeconds) % monthInSeconds) / dayInSeconds
	hours := (((privateWorkProjectsTotal % yearInSeconds) % monthInSeconds) % dayInSeconds) / hourInSeconds
	minutes := ((((privateWorkProjectsTotal % yearInSeconds) % monthInSeconds) % dayInSeconds) % hourInSeconds) / minuteInSeconds

	switch {
	case years > 0:
		return fmt.Sprintf("%d years %d months %d days %d hrs %d mins", years, months, days, hours, minutes)
	case months > 0:
		return fmt.Sprintf("%d months %d days %d hrs %d mins", months, days, hours, minutes)
	case days > 0:
		return fmt.Sprintf("%d days %d hrs %d mins", days, hours, minutes)
	case hours > 0:
		return fmt.Sprintf("%d hrs %d mins", hours, minutes)
	default:
		return fmt.Sprintf("%d mins", minutes)
	}
}

func formatObjects(objects []SortedObject) string {
	var formattedObjects string
	for _, object := range objects {
		objectLen := len(object.Name)

		if objectLen >= maxProjectNameLength {
			truncatedObject := object.Name[:maxProjectNameLength-len("...")] + "..."
			formattedObjects += fmt.Sprintf("%s %s\n", truncatedObject, object.Text)
			continue
		}

		spaces := maxProjectNameLength - objectLen
		formattedObjects += fmt.Sprintf("%s%s%s\n", object.Name, getSpaces(spaces), object.Text)
		continue
	}

	return formattedObjects
}

func getSpaces(spaces int) string {
	var s string
	for i := 0; i < spaces; i++ {
		s += " "
	}

	return s
}

// Note: idk how this event happens, but it actually does...
// (https://github.com/ZerGo0/ZerGo0/commit/d3a9a9c5f4e242bf1997fb56921d3a8483f05bad)
func calculateMaxCodingTime(targetStats *wakatime.WakaStats, prevStats *wakatime.WakaStats) string {
	targetCodingTime := targetStats.Data.HumanReadableTotalIncludingOtherLanguage
	if targetStats.Data.TotalSecondsIncludingOtherLanguage <
		prevStats.Data.TotalSecondsIncludingOtherLanguage {
		targetCodingTime = prevStats.Data.HumanReadableTotalIncludingOtherLanguage
	}

	return targetCodingTime
}
