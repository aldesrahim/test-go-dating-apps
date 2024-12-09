package enums

import "errors"

type ProductType string

const (
	ProductTypeInfiniteSwipe   ProductType = "infSwipe"
	ProductTypeVerifiedProfile ProductType = "verifiedProfile"
	ProductTypeCombo1          ProductType = "combo1" // infSwipe + verifiedProfile
)

func (t *ProductType) label() (string, error) {
	switch *t {
	case ProductTypeInfiniteSwipe:
		return "Infinite Swipe", nil
	case ProductTypeVerifiedProfile:
		return "Verified Profile", nil
	case ProductTypeCombo1:
		return "Combo (Infinite Swipe + Verified Profile)", nil
	default:
		return "", errors.New("invalid product type")
	}
}
