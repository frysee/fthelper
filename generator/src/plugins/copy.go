package plugins

import (
	"github.com/frysee/fthelper/generator/v4/src/clusters"
	"github.com/frysee/fthelper/shared/fs"
	"github.com/frysee/fthelper/shared/maps"
	"github.com/frysee/fthelper/shared/runners"
)

func Copy(data maps.Mapper, config maps.Mapper) *runners.Runner {
	return clusters.NewRunnerV2(data, config, func(p *clusters.ExecutorParameter) error {
		input, err := fs.Build(fs.ToObject(p.Data.Zi("input"), p.Config), p.VarConfig)
		if err != nil {
			p.Logger.Error("cannot get output information")
			return err
		}

		output, err := fs.Build(fs.ToObject(p.Data.Zi("output"), p.Config), p.VarConfig)
		if err != nil {
			p.Logger.Error("cannot get output information")
			return err
		}

		return fs.Copy(input.Single(), output.Single())
	}, &clusters.Settings{
		DefaultWithCluster: false,
	})
}
