package plugins

import (
	"fmt"

	"github.com/frysee/fthelper/shared/commandline/flags"
	"github.com/frysee/fthelper/shared/configs"
	"github.com/frysee/fthelper/shared/maps"
)

// SupportFSVar will add --fsvar "<name>=<value>" for assign data to fs.variables
// @deprecated SupportFSVar, use SupportVar instead
func SupportFSVar(p *PluginParameter) error {
	p.NewFlags(flags.Array{
		Name:    "fsvar",
		Default: []string{},
		Usage:   "add data to fs.variables config",
		Action: func(data []string) maps.Mapper {
			if len(data) > 0 {
				p.Logger.Warn("--fsvar is deprecated, please use --var instead")
			}

			var m = maps.New()
			for _, d := range data {
				var key, value, ok = configs.ParseOverride(d)
				if ok {
					m.Set(fmt.Sprintf("fs.variables.%s", key), value)
				}
			}
			return m
		},
	})
	return nil
}
