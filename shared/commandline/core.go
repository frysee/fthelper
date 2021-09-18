package commandline

import (
	"github.com/frysee/fthelper/shared/caches"
	"github.com/frysee/fthelper/shared/commandline/commands"
	"github.com/frysee/fthelper/shared/commandline/flags"
	"github.com/frysee/fthelper/shared/commandline/hooks"
	"github.com/frysee/fthelper/shared/commandline/models"
	"github.com/frysee/fthelper/shared/commandline/plugins"
	"github.com/frysee/fthelper/shared/loggers"
)

func New(cache *caches.Service, metadata *models.Metadata) *cli {
	return &cli{
		Metadata: metadata,
		flags:    flags.New(),
		commands: commands.New(),
		hooks:    hooks.New(),
		plugins:  plugins.New(),

		cache:  cache,
		logger: loggers.Get("commandline"),
	}
}
