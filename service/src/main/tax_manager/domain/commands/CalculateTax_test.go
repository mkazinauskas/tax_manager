package commands

import (
	"testing"
	"time"
)

func TestCalculateTax(testContext *testing.T) {
	currentTime, _ := time.Parse("yyyy-MM-dd", "2016-01-01")
	if Calculate("Vilnius", currentTime) != 0.1 {
		testContext.Error("Vilnius tax for 2016-01-01 was not 0.1")
	}
}
