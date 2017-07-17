package utils

import (
	"main/tax_manager"
	"time"
)

func ParseDate(date string) (time.Time) {
	parsedTime, parsingError := time.Parse(tax_manager.DEFAULT_DATE_FORMAT, date)
	CheckError(parsingError, "Failed to parse date from `%s` as `%s`", date, tax_manager.DEFAULT_DATE_FORMAT)
	return parsedTime
}
