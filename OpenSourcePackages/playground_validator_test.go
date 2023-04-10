package __

import (
	"fmt"
	"testing"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	val "gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

func TestV9(t *testing.T) {
	d := struct {
		Name  string `validate:"required"`
		Email string `validate:"required,email"`
	}{
		Name:  "",
		Email: "x",
	}
	err := val.New().Struct(d)
	for _, e := range err.(val.ValidationErrors) {
		fmt.Println(e)
	}
}

func TestBetterMessages(t *testing.T) {
	var err error

	d := struct {
		Name  string `validate:"required"`
		Email string `validate:"required,email"`
	}{
		Name:  "",
		Email: "x@y",
	}
	v := val.New()

	translator := en.New()
	uni := ut.New(translator, translator)

	// this is usually known or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, found := uni.GetTranslator("en")
	if !found {
		t.Fatal("translator not found")
	}
	if err = en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		t.Fatal(err)
	}

	err = v.Struct(d)
	for _, e := range err.(val.ValidationErrors) {
		fmt.Println(e.Field(), e.Tag(), ":", e.Translate(trans))
	}
}
