package flags

import (
	"flag"

	"github.com/frysee/fthelper/shared/maps"
)

type Flag interface {
	FlagName() string
	Parse(flag *flag.FlagSet) interface{}
	Build(value interface{}) maps.Mapper
}
