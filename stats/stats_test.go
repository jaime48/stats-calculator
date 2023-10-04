package stats

import (
	"testing"
)

func TestCalculateRecipeStats(t *testing.T) {
	stats := CalculateRecipeStats("test.json", "Chicken,Steak", "10224", "10AM", "3PM")

	expectedUniqueCount := 3
	if stats.UniqueRecipeCount != expectedUniqueCount {
		t.Errorf("Unexpected unique recipe count. Expected: %d, Got: %d", expectedUniqueCount, stats.UniqueRecipeCount)
	}

	exptectedBusiestPostcode := "10208"
	if stats.BusiestPostcode.Postcode != exptectedBusiestPostcode {
		t.Errorf("Unexpected busiest postcode. Expected: %s, Got: %s", exptectedBusiestPostcode, stats.BusiestPostcode.Postcode)
	}
	exptectedBusiestPostcodeCount := 2
	if stats.BusiestPostcode.DeliveryCount != exptectedBusiestPostcodeCount {
		t.Errorf("Unexpected busiest postcode count. Expected:  %d, Got:  %d", exptectedBusiestPostcodeCount, stats.BusiestPostcode.DeliveryCount)
	}
}