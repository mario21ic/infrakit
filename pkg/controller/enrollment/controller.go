package enrollment

import (
	"github.com/docker/infrakit/pkg/controller"
	enrollment "github.com/docker/infrakit/pkg/controller/enrollment/types"
	"github.com/docker/infrakit/pkg/controller/internal"
	logutil "github.com/docker/infrakit/pkg/log"
	"github.com/docker/infrakit/pkg/run/scope"
	"github.com/docker/infrakit/pkg/spi/stack"
	"github.com/docker/infrakit/pkg/types"
)

var (
	log    = logutil.New("module", "controller/enrollment")
	debugV = logutil.V(200)
)

// NewController returns a controller implementation
func NewController(scope scope.Scope, leader func() stack.Leadership,
	options enrollment.Options) controller.Controller {
	return internal.NewController(
		leader,
		// the constructor
		func(spec types.Spec) (internal.Managed, error) {
			return newEnroller(scope, leader, options), nil
		},
		// the key function
		func(metadata types.Metadata) string {
			return metadata.Name
		},
	)
}

// NewTypedControllers return typed controllers
func NewTypedControllers(scope scope.Scope, leader func() stack.Leadership,
	options enrollment.Options) func() (map[string]controller.Controller, error) {

	return (internal.NewController(
		leader,
		// the constructor
		func(spec types.Spec) (internal.Managed, error) {
			log.Debug("Creating managed object", "spec", spec)
			return newEnroller(scope, leader, options), nil
		},
		// the key function
		func(metadata types.Metadata) string {
			return metadata.Name
		},
	)).ManagedObjects
}

func (l *enroller) started() bool {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.running
}
