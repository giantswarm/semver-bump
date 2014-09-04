package storage

import "regexp"

func extractVersionNumber(s string) string {
	re := regexp.MustCompile("[0-9]{1,4}\\.[0-9]{1,4}\\.[0-9]{1,4}")
	return re.FindString(s)
}
