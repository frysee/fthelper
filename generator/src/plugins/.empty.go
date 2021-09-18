package plugins

import (
	"github.com/frysee/fthelper/generator/v4/src/clusters"
	"github.com/frysee/fthelper/shared/maps"
	"github.com/frysee/fthelper/shared/runners"
)

func Empty(data maps.Mapper, config maps.Mapper) *runners.Runner {
	return clusters.NewRunnerV2(data, config, func(p *clusters.ExecutorParameter) error {
		return nil
	}, &clusters.Settings{
		DefaultWithCluster: true,
	})
}
