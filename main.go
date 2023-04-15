package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/9illes/pagespeed-command/analyze"
	"github.com/joho/godotenv"
)

// URL is the URL to analyze
var URL string

// strategy is the strategy to use (MOBILE|DESKTOP)
var strategy string

func init() {
	flag.StringVar(&URL, "url", "", "Specify URL (e.g. https://www.example.com/)")
	flag.StringVar(&strategy, "strategy", "MOBILE", "Specify strategy (MOBILE|DESKTOP). Default is MOBILE")
	flag.Parse()

	if len(URL) == 0 {
		fmt.Println("Please specify a URL (e.g. -url=https://www.example.com/)")
		os.Exit(1)
	}

	if len(strategy) > 0 && strategy != analyze.DESKTOP && strategy != analyze.MOBILE {
		fmt.Println("Please specify a valid strategy (MOBILE|DESKTOP)")
		os.Exit(1)
	}

	godotenv.Load(".env.local")
	godotenv.Load(".env")
}

func main() {
	googleAPIKey := os.Getenv("GOOGLE_API_KEY")

	if len(googleAPIKey) == 0 {
		fmt.Println("Please set the GOOGLE_API_KEY environment variable")
		os.Exit(1)
	}

	analyzer := analyze.NewPerfAnalyzer(googleAPIKey)
	result, err := analyzer.Analyze(URL, "DESKTOP")

	if err != nil {
		panic(err)
	}

	bytes, _ := json.Marshal(result)
	fmt.Println(string(bytes))
}
