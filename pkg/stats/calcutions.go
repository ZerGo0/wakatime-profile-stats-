package stats

import (
	"slices"

	"github.com/user/wakatime-profile-stats/pkg/wakatime"
)

const (
	maxTopItems = 5
)

func calculateTopProjects(stats *wakatime.WakaStats) []SortedObject {
	slices.SortFunc(stats.Data.Projects, func(a, b wakatime.Projects) int {
		return int(b.TotalSeconds - a.TotalSeconds)
	})

	topProjects := make([]SortedObject, 0)
	for i, project := range stats.Data.Projects {
		if i == maxTopItems {
			break
		}

		topProjects = append(topProjects, SortedObject{
			Name: project.Name,
			Text: project.Text,
		})
	}

	diff := maxTopItems - len(topProjects)
	if diff > 0 {
		for i := 0; i < diff; i++ {
			topProjects = append(topProjects, SortedObject{
				Name: "",
				Text: "",
			})
		}
	}

	return topProjects
}

func calculateTopLanguages(stats *wakatime.WakaStats) []SortedObject {
	slices.SortFunc(stats.Data.Languages, func(a, b wakatime.Languages) int {
		return int(b.TotalSeconds - a.TotalSeconds)
	})

	topLanguages := make([]SortedObject, 0)
	for i, language := range stats.Data.Languages {
		if i == maxTopItems {
			break
		}

		topLanguages = append(topLanguages, SortedObject{
			Name: language.Name,
			Text: language.Text,
		})
	}

	diff := maxTopItems - len(topLanguages)
	if diff > 0 {
		for i := 0; i < diff; i++ {
			topLanguages = append(topLanguages, SortedObject{
				Name: "",
				Text: "",
			})
		}
	}

	return topLanguages
}
