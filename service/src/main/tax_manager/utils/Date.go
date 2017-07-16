package utils

import (
	"main/tax_manager"
	"time"
)

func Parse(date string) (time.Time) {
	parsedTime, _ := time.Parse(tax_manager.DEFAULT_DATE_FORMAT, date)
	return parsedTime
}
