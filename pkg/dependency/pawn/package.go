package pawn

import (
	"github.com/Southclaws/sampctl/pkg/dependency"
	"github.com/Southclaws/sampctl/versioning"
)

// Package represents a definition for a Pawn package and can either be used to define a build or
// as a description of a package in a repository. This is akin to npm's package.json and combines
// a project's dependencies with a description of that project.
//
// For example, a gamemode that includes a library does not need to define the User, Repo, Version,
// Contributors and Include fields at all, it can just define the Dependencies list in order to
// build correctly.
//
// On the flip side, a library written in pure Pawn should define some contributors and a web URL
// but, being written in pure Pawn, has no dependencies.
//
// Finally, if a repository stores its package source files in a subdirectory, that directory should
// be specified in the Include field. This is common practice for plugins that store the plugin
// source code in the root and the Pawn source in a subdirectory called 'include'.
// nolint:lll
type Package struct {
	// Parent indicates that this package is a "working" package that the user has explicitly
	// created and is developing. The opposite of this would be packages that exist in the
	// `dependencies` directory that have been downloaded as a result of an Ensure.
	Parent bool `json:"-" yaml:"-"`
	// LocalPath indicates the Package object represents a local copy which is a directory
	// containing a `samp.json`/`samp.yaml` file and a set of Pawn source code files.
	// If this field is not set, then the Package is just an in-memory pointer to a remote package.
	LocalPath string `json:"-" yaml:"-"`
	// The vendor directory - for simple packages with no sub-dependencies, this is simply
	// `<local>/dependencies` but for nested dependencies, this needs to be set.
	Vendor string `json:"-" yaml:"-"`
	// format stores the original format of the package definition file, either `json` or `yaml`
	Format string `json:"-" yaml:"-"`

	// Inferred metadata, not always explicitly set via JSON/YAML but inferred from the dependency path
	versioning.DependencyMeta

	// Metadata, set by the package author to describe the package
	Contributors []string `json:"contributors,omitempty" yaml:"contributors,omitempty"` // list of contributors
	Website      string   `json:"website,omitempty" yaml:"website,omitempty"`           // website or forum topic associated with the package

	// Functional, set by the package author to declare relevant files and dependencies
	Entry       string   `json:"entry,omitempty"`                                              // entry point script to compile the project
	Output      string   `json:"output,omitempty"`                                             // output amx file
	Deps        []string `json:"dependencies,omitempty" yaml:"dependencies,omitempty"`         // list of packages that the package depends on
	Development []string `json:"dev_dependencies,omitempty" yaml:"dev_dependencies,omitempty"` // list of packages that only the package builds depend on
	Local       bool     `json:"local,omitempty" yaml:"local,omitempty"`                       // run package in local dir instead of in a temporary runtime
	// Build        *BuildConfig                  `json:"build,omitempty" yaml:"build,omitempty"`                       // build configuration
	// Builds       []*BuildConfig                `json:"builds,omitempty" yaml:"builds,omitempty"`                     // multiple build configurations
	// Runtime      *Runtime                      `json:"runtime,omitempty" yaml:"runtime,omitempty"`                   // runtime configuration
	// Runtimes     []*Runtime                    `json:"runtimes,omitempty" yaml:"runtimes,omitempty"`                 // multiple runtime configurations
	// IncludePath  string                        `json:"include_path,omitempty" yaml:"include_path,omitempty"`         // include path within the repository, so users don't need to specify the path explicitly
	// Resources    []Resource                    `json:"resources,omitempty" yaml:"resources,omitempty"`               // list of additional resources associated with the package
}

// Version - Implements dependency.Resource
func (p *Package) Version() (version string) { return p.DependencyMeta.String() }

// Cached - Implements dependency.Resource
func (p *Package) Cached() (is bool, path string) { return }

// Ensure - Implements dependency.Resource
func (p *Package) Ensure() (err error) { return }

// Dependencies - Implements dependency.Resource
func (p *Package) Dependencies() (resources []dependency.Resource) { return }
