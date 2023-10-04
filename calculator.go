package main

import (
	"encoding/json"
	"fmt"
	"os"
	"flag"
	"recipe-stats-calculator/stats"
)

func main() {
	fixturesFile := flag.String("file", "stats/test.json", "path to the fixtures file")
	searchRecipeNames := flag.String("recipes", "", "comma-separated recipe names to search by")
	searchPostcode := flag.String("postcode", "10120", "postcode to search for deliveries")
	searchTimeFrom := flag.String("from", "10AM", "start time for delivery search window")
	searchTimeTo := flag.String("to", "3PM", "end time for delivery search window")
	flag.Parse()
	
	// Calculate the recipe statistics
	stats := stats.CalculateRecipeStats(*fixturesFile, *searchRecipeNames, *searchPostcode, *searchTimeFrom, *searchTimeTo)
	
	// Generate JSON output
	output, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		os.Stderr.WriteString("Error generating JSON")
		os.Exit(1)
	}
	fmt.Println(string(output))
}





