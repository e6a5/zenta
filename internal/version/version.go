// Package version provides build and version information for zenta.
// It includes details about the version, git commit, build date, and runtime environment.
package version

import (
	"fmt"
	"runtime"
)

// Build information. Populated at build-time.
var (
	Version   = "dev"
	GitCommit = "unknown"
	BuildDate = "unknown"
	GoVersion = runtime.Version()
)

// Info represents version information
type Info struct {
	Version   string `json:"version"`
	GitCommit string `json:"git_commit"`
	BuildDate string `json:"build_date"`
	GoVersion string `json:"go_version"`
	Platform  string `json:"platform"`
}

// Get returns version information
func Get() Info {
	return Info{
		Version:   Version,
		GitCommit: GitCommit,
		BuildDate: BuildDate,
		GoVersion: GoVersion,
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

// String returns version as a string
func (i Info) String() string {
	return fmt.Sprintf("zenta %s (%s) built with %s on %s for %s",
		i.Version, i.GitCommit, i.GoVersion, i.BuildDate, i.Platform)
}
