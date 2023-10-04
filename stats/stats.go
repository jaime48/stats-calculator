package stats

import (
	"os"
	"sort"
	"strings"
	"time"
	"io/ioutil"
	"encoding/json"
)


type Recipe struct {
	Postcode  string `json:"postcode"`
	Recipe    string `json:"recipe"`
	Delivery  string `json:"delivery"`
}

type RecipeStats struct {
	UniqueRecipeCount       int                  `json:"unique_recipe_count"`
	CountPerRecipe          []RecipeCount        `json:"count_per_recipe"`
	BusiestPostcode         PostcodeDelivery     `json:"busiest_postcode"`
	CountPerPostcodeAndTime PostcodeTimeDelivery `json:"count_per_postcode_and_time"`
	MatchByName             []string             `json:"match_by_name"`
}

type RecipeCount struct {
	Recipe string `json:"recipe"`
	Count  int    `json:"count"`
}

type PostcodeDelivery struct {
	Postcode      string `json:"postcode"`
	DeliveryCount int    `json:"delivery_count"`
}

type PostcodeTimeDelivery struct {
	Postcode      string `json:"postcode"`
	From          string `json:"from"`
	To            string `json:"to"`
	DeliveryCount int    `json:"delivery_count"`
}

func CalculateRecipeStats(fixturesFile string, searchRecipeNames string, searchPostcode string, searchTimeFrom string, searchTimeTo string) RecipeStats {
	jsonData, err := ioutil.ReadFile(fixturesFile)
	if err != nil {
		os.Stderr.WriteString("Error reading file")
		os.Exit(1)
	}
	// Parse JSON data into Recipe slice
	var recipes []Recipe
	err = json.Unmarshal(jsonData, &recipes)
	if err != nil {
		os.Stderr.WriteString("Error parsing JSON")
		os.Exit(1)
	}
	
	uniqueRecipeCount, countPerRecipe := countRecipes(recipes)
	sort.Slice(countPerRecipe, func(i, j int) bool {
		return countPerRecipe[i].Recipe < countPerRecipe[j].Recipe
	})

	// Find the postcode with the most delivered recipes
	postcodeCounts := make(map[string]int)
	for _, recipe := range recipes {
		postcodeCounts[recipe.Postcode]++
	}
	busiestPostcode := getBusiestPostcode(postcodeCounts)

	// Count the number of deliveries to postcode 10120 that lie within the delivery time between 10AM and 3PM
	countPerPostcodeAndTime := countDeliveriesInRange(recipes, searchPostcode, searchTimeFrom, searchTimeTo)
	// List the recipe names (alphabetically ordered) that contain the specified words
	matchByNames := getRecipeNamesContaining(recipes, strings.Split(searchRecipeNames, ","))
	return RecipeStats{
		UniqueRecipeCount:       uniqueRecipeCount,
		CountPerRecipe:          countPerRecipe,
		BusiestPostcode:         busiestPostcode,
		CountPerPostcodeAndTime: countPerPostcodeAndTime,
		MatchByName:             matchByNames,
	}
}

func getBusiestPostcode(postcodeCounts map[string]int) PostcodeDelivery {
	maxCount := 0
	busiestPostcode := ""
	for postcode, count := range postcodeCounts {
		if count > maxCount {
			maxCount = count
			busiestPostcode = postcode
		}
	}
	return PostcodeDelivery{
		Postcode:      busiestPostcode,
		DeliveryCount: maxCount,
	}
}

func countDeliveriesInRange(recipes []Recipe, postcode, from, to string) PostcodeTimeDelivery {
	startTime, _ := time.Parse("3PM", from)
	endTime, _ := time.Parse("3PM", to)
	deliveryCount := 0
	for _, recipe := range recipes {
		if recipe.Postcode == postcode {
			deliveryStartTime, deliveryEndTime, err := parseDeliveryTime(recipe.Delivery)
			if err != nil {
				os.Stderr.WriteString("Error generating JSON")
				os.Exit(1)
			}
			if deliveryStartTime.After(startTime) && deliveryEndTime.Before(endTime) {
				deliveryCount++
			}
		}
	}
	return PostcodeTimeDelivery{
		Postcode:      postcode,
		From:          from,
		To:            to,
		DeliveryCount: deliveryCount,
	}
}

func parseDeliveryTime(delivery string) (time.Time, time.Time, error) {
	delivery = strings.ReplaceAll(delivery, " -", "")
	parts := strings.Split(delivery, " ")
	startTime, err := time.Parse("3PM", parts[1])
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	endTime, err := time.Parse("3PM", parts[2])
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	return startTime, endTime, nil
}

func getRecipeNamesContaining(recipes []Recipe, words []string) []string {
	matchingRecipes := make(map[string]bool)
	for _, recipe := range recipes {
		for _, word := range words {
			if strings.Contains(recipe.Recipe, word) {
				matchingRecipes[recipe.Recipe] = true
				break
			}
		}
	}

	result := make([]string, 0, len(matchingRecipes))
	for recipe := range matchingRecipes {
		result = append(result, recipe)
	}

	sort.Strings(result)
	return result
}

func countRecipes(recipes []Recipe) (int, []RecipeCount) {
	recipeCounts := make(map[string]int)

	for _, recipe := range recipes {
		recipeCounts[recipe.Recipe]++
	}

	countPerRecipe := make([]RecipeCount, 0, len(recipeCounts))
	for recipe, count := range recipeCounts {
		countPerRecipe = append(countPerRecipe, RecipeCount{Recipe: recipe, Count: count})
	}

	return len(countPerRecipe), countPerRecipe
}