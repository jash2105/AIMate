package form

import (
	"github.com/jash2105/vault/errorlist"
	"github.com/jash2105/vault/validator"
)

type QuestionForm struct {
	Question string `schema:"question"`
	Model    string `schema:model`
	UUID     string `schema:uuid`
	ApiKey   string `schema:apikey`
}

func (me *QuestionForm) Validate() errorlist.Errors {
	errs := errorlist.New()

	validator.CheckNotEmpty(me.Question, "question", errs)
	validator.CheckNotEmpty(me.Model, "model", errs)

	return errs
}

func (me *QuestionForm) String() string {
	return me.Question
}
