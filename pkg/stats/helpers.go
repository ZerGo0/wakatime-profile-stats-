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
	stats.Data.Projects = lo.Filter(stats.Data.Projects, func(p wakatime.Projects, i int) bool {
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
	years := int(privateWorkProjectsTotal / yearInSeconds)
	months := int((privateWorkProjectsTotal % yearInSeconds) / monthInSeconds)
	days := int(((privateWorkProjectsTotal % yearInSeconds) % monthInSeconds) / dayInSeconds)
	hours := int((((privateWorkProjectsTotal % yearInSeconds) % monthInSeconds) % dayInSeconds) / hourInSeconds)
	minutes := int(((((privateWorkProjectsTotal % yearInSeconds) % monthInSeconds) % dayInSeconds) % hourInSeconds) / minuteInSeconds)

	if years > 0 {
		return fmt.Sprintf("%d years %d months %d days %d hrs %d mins", years, months, days, hours, minutes)
	} else if months > 0 {
		return fmt.Sprintf("%d months %d days %d hrs %d mins", months, days, hours, minutes)
	} else if days > 0 {
		return fmt.Sprintf("%d days %d hrs %d mins", days, hours, minutes)
	} else if hours > 0 {
		return fmt.Sprintf("%d hrs %d mins", hours, minutes)
	} else {
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
		} else {
			spaces := maxProjectNameLength - objectLen
			formattedObjects += fmt.Sprintf("%s%s%s\n", object.Name, getSpaces(spaces), object.Text)
			continue
		}
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
