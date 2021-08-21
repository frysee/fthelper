package plugins

import (
	"github.com/kamontat/fthelper/generator/v4/src/clusters"
	"github.com/kamontat/fthelper/shared/fs"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/runners"
)

func Copy(data maps.Mapper, config maps.Mapper) runners.Runner {
	return clusters.NewRunner(data, config, func(p *clusters.ExecutorParameter) error {
		input, err := fs.Build(p.Data.Mi("input"), p.VarConfig)
		if err != nil {
			p.Logger.Error("cannot get output information")
			return err
		}

		output, err := fs.Build(p.Data.Mi("output"), p.VarConfig)
		if err != nil {
			p.Logger.Error("cannot get output information")
			return err
		}

		return fs.Copy(input.Single(), output.Single())
	}, &clusters.Settings{
		DefaultWithCluster: false,
	})
}
