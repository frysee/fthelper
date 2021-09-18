package hooks

import "github.com/frysee/fthelper/shared/maps"

// Hook action
type Hook func(config maps.Mapper) error
