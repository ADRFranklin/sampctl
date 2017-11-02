package rook

import (
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/Southclaws/sampctl/compiler"
	"github.com/Southclaws/sampctl/download"
)

// Build compiles a package, dependencies are ensured and a list of paths are sent to the compiler.
func (pkg Package) Build(version string) (output string, err error) {
	err = pkg.EnsureDependencies()
	if err != nil {
		return
	}

	includes := make([]string, len(pkg.Dependencies))

	for _, depStr := range pkg.Dependencies {
		dep, err := PackageFromDep(depStr)
		if err != nil {
			return "", errors.Errorf("package dependency '%s' is invalid: %v", depStr, err)
		}

		includes = append(includes, filepath.Join(pkg.local, "dependencies", dep.repo))
	}

	cacheDir, err := download.GetCacheDir()
	if err != nil {
		return
	}

	err = compiler.CompileSource(
		filepath.Join(pkg.local, pkg.Entry),
		filepath.Join(pkg.local, pkg.Output),
		includes,
		cacheDir,
		version,
	)
	if err != nil {
		return
	}

	output = pkg.Output

	return
}