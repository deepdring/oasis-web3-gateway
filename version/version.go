// Package version provides the version information for the Oasis Web3 Gateway.
package version

import (
	"runtime"
	"runtime/debug"
)

// Unknown represents an unknown version.
const Unknown = "unknown"

var (
	// Software represents the Emerald Web3 Gateway's version and
	// should be set by the linker.
	Software = "0.0.0-unset"

	// Toolchain is the version of the Go compiler/standard library.
	Toolchain = runtime.Version()
)

// GetOasisSDKVersion returns the version of the Oasis SDK dependency or
// unknown if it can't be obtained.
func GetOasisSDKVersion() string {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return Unknown
	}

	for _, dep := range bi.Deps {
		if dep.Path == "github.com/oasisprotocol/oasis-sdk/client-sdk/go" {
			return dep.Version
		}
	}
	return Unknown
}
