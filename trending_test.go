package trending

import (
	"testing"
)

var (
	// Get trending repos from all languages today
	allDailyResponse   = All("daily")
	allWeeklyResponse  = All("weekly")
	allMonthlyResponse = All("monthly")

	// Get trending repos from Jupyter Notebook
	jupyterDailyResponse = Language("Jupyter Notebook", "daily")
)

func TestReturnSliceWith25Repositories(t *testing.T) {
	if len(allDailyResponse) != 25 {
		t.Error("Length of returned slice should be 25")
	}
}

func TestReturnsDifferentLanguages(t *testing.T) {
	languages := map[string]int{}

	for _, repo := range allDailyResponse {
		if _, ok := languages[repo.language]; ok {
			languages[repo.language]++
		} else {
			languages[repo.language] = 1
		}
	}

	if len(languages) == 1 {
		t.Error("Only found repositories from a single language")
	}
}

func TestPeriodShouldReturnDifferentRepos(t *testing.T) {
	dailyAndWeeklyMatches := true
	dailyAndMonthlyMatches := true
	weeklyAndMonthlyMatches := true

	for i := 0; i < 25; i++ {
		if allDailyResponse[i].name != allWeeklyResponse[i].name {
			dailyAndWeeklyMatches = false
		}

		if allDailyResponse[i].name != allMonthlyResponse[i].name {
			dailyAndMonthlyMatches = false
		}

		if allWeeklyResponse[i].name != allMonthlyResponse[i].name {
			weeklyAndMonthlyMatches = false
		}
	}

	if dailyAndWeeklyMatches {
		t.Error("Daily and weekly repos should not match")
	}

	if dailyAndMonthlyMatches {
		t.Error("Daily and monthly repos should not match")
	}

	if weeklyAndMonthlyMatches {
		t.Error("Weekly and monthly repos should not match")
	}
}

func TestReturnsOnlyTheSpecifiedLanguage(t *testing.T) {
	languages := map[string]int{}

	for _, repo := range jupyterDailyResponse {
		if _, ok := languages[repo.language]; ok {
			languages[repo.language]++
		} else {
			languages[repo.language] = 1
		}
	}

	if len(languages) != 1 {
		t.Error("Found repositories from a multiple languages")
	}
}
