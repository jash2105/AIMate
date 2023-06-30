package validator

import (
	"strings"

	"github.com/jash2105/vault/errorlist"
)

type Validator interface {
	Validate(map[string]error)
}

func CheckNotEmpty(input, name string, errs errorlist.Errors) {
	if len(strings.TrimSpace(input)) == 0 {
		errs[name] = errorlist.NewError("cannot be blank")
	}
}
