package form

import "github.com/jash2105/vault/errorlist"

type Form interface {
	Validate() errorlist.Errors
	String() string
}
