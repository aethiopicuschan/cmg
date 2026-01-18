package variables

import "runtime/debug"

func Version() (version string) {
	bi, ok := debug.ReadBuildInfo()
	if ok && bi.Main.Version != "" {
		version = bi.Main.Version
	} else {
		version = "unknown"
	}
	return
}
