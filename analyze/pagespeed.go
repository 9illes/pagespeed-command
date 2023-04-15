package analyze

import (
	"context"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/pagespeedonline/v5"
)

const (
	DESKTOP = "DESKTOP"
	MOBILE  = "MOBILE"
)

// Result is a struct that contains the result of the analysis
type Result struct {
	URL      string    `json:"url"`
	Strategy string    `json:"strategy"`
	Score    float64   `json:"score"`
	Date     time.Time `json:"date"`
}

// PerfAnalyzer is a struct that contains the pagespeedonline service
type PerfAnalyzer struct {
	pagespeedonlineService *pagespeedonline.Service
	apiKey                 string
}

// NewPerfAnalyzer creates a new PerfAnalyzer
func NewPerfAnalyzer(apiKey string) *PerfAnalyzer {
	pagespeedonlineService, err := pagespeedonline.NewService(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		panic(err)
	}

	return &PerfAnalyzer{
		pagespeedonlineService: pagespeedonlineService,
		apiKey:                 apiKey,
	}
}

// Analyze returns the score of a given url and strategy
func (a *PerfAnalyzer) Analyze(url string, strategy string) (Result, error) {
	call := a.pagespeedonlineService.Pagespeedapi.Runpagespeed(url)
	call.Category("PERFORMANCE")
	call.Strategy(strategy)

	pagespeedApiPagespeedResponseV5, err := call.Do()
	if err != nil {
		panic(err)
	}

	score, _ := pagespeedApiPagespeedResponseV5.LighthouseResult.Categories.Performance.Score.(float64)

	return Result{
		URL:      url,
		Strategy: strategy,
		Score:    score,
		Date:     time.Now(),
	}, nil
}
