package trending

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

type Repository struct {
	author      string
	name        string
	href        string
	description string
	language    string
	stars       int
	forks       int
}

func All(period string) []Repository {
	return scrape("https://github.com/trending?since=" + period)
}

func Language(language, period string) []Repository {
	language = strings.ToLower(language)
	language = strings.Replace(language, " ", "-", -1)

	return scrape("https://github.com/trending/" + language + "?since=" + period)
}

func scrape(url string) []Repository {
	document, err := goquery.NewDocument(url)

	if err != nil {
		panic(err)
	}

	repositories := []Repository{}

	document.
		Find("ol.repo-list li").
		Each(func(i int, s *goquery.Selection) {
			titleText := s.Find("h3").Text()
			titleText = strings.TrimSpace(titleText)
			titleTextPieces := strings.Split(titleText, " / ")

			repo := titleTextPieces[0] + "/" + titleTextPieces[1]

			description := s.Find(".py-1 p").Text()
			description = strings.TrimSpace(description)

			language := s.Find("[itemprop=programmingLanguage]").Text()
			language = strings.TrimSpace(language)

			starLink := "/" + repo + "/stargazers"
			forkLink := "/" + repo + "/network"

			stars := s.Find("[href='" + starLink + "']").Text()
			stars = strings.TrimSpace(stars)
			stars = strings.Replace(stars, ",", "", -1)
			starsInt, _ := strconv.Atoi(stars)

			forks := s.Find("[href='" + forkLink + "']").Text()
			forks = strings.TrimSpace(forks)
			forks = strings.Replace(forks, ",", "", -1)
			forksInt, _ := strconv.Atoi(forks)

			repository := Repository{
				author:      titleTextPieces[0],
				name:        titleTextPieces[1],
				href:        "https://github.com/" + repo,
				description: description,
				language:    language,
				stars:       starsInt,
				forks:       forksInt,
			}

			repositories = append(repositories, repository)
		})

	return repositories
}
