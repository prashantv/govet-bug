package registry_test

import (
	"testing"

	"github.com/prashantv/govet-bug/registry/registrytest"
)

func TestFoo(t *testing.T) {
	r := registrytest.New()
	r.HandleGroup()
}
