package xtemplates_test

import (
	"testing"

	"github.com/frysee/fthelper/shared/maps"
	"github.com/frysee/fthelper/shared/xtemplates"
	"github.com/frysee/fthelper/shared/xtests"
)

func TestXtemplate(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("invalid template").
		WithExpected("function \"invalid\" not defined").
		WithActualAndError(xtemplates.Text("{{ invalid \"function\" }}", maps.New())).
		MustContainError()
}
