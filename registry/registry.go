package registry

type Registry struct{}

func New() *Registry { return &Registry{} }

func (*Registry) handleGroup() {}
