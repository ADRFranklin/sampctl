package dependency

// Resource descibes a required object that can be cached, identified by version
// and ensured to a specific directory path.
type Resource interface {
	Version() (version string)                    // The resource version
	Cached(version string) (is bool, path string) // Get cached path is present
	Ensure(version, path string) (err error)      // Acquire if necessary
}
