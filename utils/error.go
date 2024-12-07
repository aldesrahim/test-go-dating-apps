package utils

import "strings"

func ErrorDBIsDuplicate(err error) bool {
	return strings.HasPrefix(err.Error(), "Error 1062 (23000): Duplicate entry")
}
