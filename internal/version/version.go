package version

var (
	Version   string
	GitCommit string
	BuildDate string
)

// VersionInfo contains information about app version.
type VersionInfo struct {
	Version   string
	GitCommit string
	BuildDate string
}

func GetVersionInfo() VersionInfo {
	return VersionInfo{
		Version:   Version,
		GitCommit: GitCommit,
		BuildDate: BuildDate,
	}
}
