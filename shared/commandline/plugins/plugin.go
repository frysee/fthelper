package plugins

import (
	"github.com/frysee/fthelper/shared/commandline/commands"
	"github.com/frysee/fthelper/shared/commandline/flags"
	"github.com/frysee/fthelper/shared/commandline/hooks"
	"github.com/frysee/fthelper/shared/commandline/models"
	"github.com/frysee/fthelper/shared/loggers"
	"github.com/frysee/fthelper/shared/maps"
)

type PluginParameter struct {
	Metadata models.Metadata

	NewCommand commands.Creator
	NewFlags   flags.Creator
	NewHook    hooks.Creator

	Config maps.Mapper
	Logger *loggers.Logger
}

type Plugin func(p *PluginParameter) error
