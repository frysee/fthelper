package commands

import (
	"github.com/frysee/fthelper/shared/caches"
	"github.com/frysee/fthelper/shared/commandline/models"
	"github.com/frysee/fthelper/shared/loggers"
	"github.com/frysee/fthelper/shared/maps"
)

type ExecutorParameter struct {
	Name   string
	Meta   *models.Metadata
	Config maps.Mapper
	Cache  *caches.Service
	Logger *loggers.Logger
	Args   []string
}

type Executor func(p *ExecutorParameter) error
