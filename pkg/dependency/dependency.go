package dependency

import (
	"github.com/Southclaws/sampctl/print"
)

// Resource descibes a required object that can be cached, identified by version
// and ensured to a specific directory path.
type Resource interface {
	Version() (version string)            // The resource version
	Cached() (is bool, path string)       // Get cached path is present
	Ensure() (err error)                  // Acquire if necessary
	Dependencies() (resources []Resource) // Get dependencies
}

// Node represents a node in the dependency tree. The dependency tree is a
// directed acyclic graph of resources.
type Node struct {
	Element  Resource // the Resource this node contains
	Children []*Node
}

// Root represents a root node of a dependency tree
type Root Node

// Build recursively discovers all dependencies and builds a tree of concrete
// typed and ensured Resource leaves.
func (n *Root) Build() {
	var f func(e Resource)
	f = func(e Resource) {
		err := e.Ensure()
		if err != nil {
			print.Erro(err)
		}

		for _, child := range e.Dependencies() {
			f(child)
		}
	}
	f(n.Element)
	return
}
