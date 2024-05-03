package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var Validator = validator.New()
var PositiveIntRegex, _ = regexp.Compile(`^[1-9]\d*$`)
