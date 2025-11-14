package validate

import (
	"github.com/go-playground/validator/v10"
)

var (
	fishing_spot_marker_validator = validator.New()
)

func Validate(a any) error {
	return fishing_spot_marker_validator.Struct(a)
}
