package operands

import (
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"kubevirt.io/ssp-operator/internal/common"
)

type Operand interface {
	// AddWatchTypesToScheme adds any additional types to the scheme.
	// The default scheme already contains types from k8s.io/api.
	AddWatchTypesToScheme(*runtime.Scheme) error

	// WatchTypes returns a slice of namespaced resources, that the operator should watch.
	WatchTypes() []client.Object

	// WatchClusterTypes returns a slice of cluster resources, that the operator should watch.
	WatchClusterTypes() []client.Object

	// Reconcile creates and updates resources.
	Reconcile(*common.Request) ([]common.ResourceStatus, error)

	// Cleanup removes any created cluster resources.
	// They don't use owner references, so the garbage collector will not remove them.
	Cleanup(*common.Request) error

	// Name returns the name of the operand
	Name() string
}
