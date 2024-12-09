package utils

import (
	"errors"
	"gorm.io/gorm"
	"strings"
)

func ErrDBIsDuplicate(err error) bool {
	return strings.HasPrefix(err.Error(), "Error 1062 (23000): Duplicate entry")
}

func ErrDBIsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
