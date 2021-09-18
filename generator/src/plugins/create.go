package plugins

import (
	"github.com/frysee/fthelper/generator/v4/src/clusters"
	"github.com/frysee/fthelper/shared/errors"
	"github.com/frysee/fthelper/shared/fs"
	"github.com/frysee/fthelper/shared/maps"
	"github.com/frysee/fthelper/shared/runners"
)

func Create(data maps.Mapper, config maps.Mapper) *runners.Runner {
	return clusters.NewRunnerV2(data, config, func(p *clusters.ExecutorParameter) error {
		var output, err = fs.Build(fs.ToObject(p.Data.Zi("output"), p.Config), p.VarConfig)
		if err != nil {
			p.Logger.Error("cannot get output information")
			return err
		}

		var errs = errors.New()
		for _, f := range output.All() {
			errs.And(f.Build())
		}

		return errs.Error()
	}, &clusters.Settings{
		DefaultWithCluster: true,
	})
}
