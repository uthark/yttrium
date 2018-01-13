package version

var (
	// Version is an app version. Set during a build.
	Version string
	// GitCommit is a git commit id. Set during a build.
	GitCommit string
	// BuildDate is a date when app was built. Set during a build. Follows ISO format.
	BuildDate string
)

// Info contains information about app version.
type Info struct {
	Version   string
	GitCommit string
	BuildDate string
}

// GetVersionInfo return current version information.
func GetVersionInfo() Info {
	return Info{
		Version:   Version,
		GitCommit: GitCommit,
		BuildDate: BuildDate,
	}
}
