package enums

import "errors"

type InterestType string

const (
	InterestFromAge InterestType = "fromAge"
	InterestToAge   InterestType = "toAge"
	InterestGender  InterestType = "gender"
)

func (t *InterestType) label() (string, error) {
	switch *t {
	case InterestFromAge:
		return "From Age", nil
	case InterestToAge:
		return "To Age", nil
	case InterestGender:
		return "Gender", nil
	default:
		return "", errors.New("invalid interest type")
	}
}
