package registrytest

import "github.com/prashantv/govet-bug/registry"

type Registry struct {
	*registry.Registry
}

func New() *Registry {
	return &Registry{registry.New()}
}
