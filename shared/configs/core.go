package configs

import (
	"github.com/frysee/fthelper/shared/loggers"
	"github.com/frysee/fthelper/shared/maps"
)

func New(name string, config maps.Mapper) *Builder {
	return &Builder{
		name:     name,
		env:      "",
		config:   config,
		override: maps.New(),
		strategy: maps.New(),

		logger: loggers.Get("config", "builder"),
	}
}
