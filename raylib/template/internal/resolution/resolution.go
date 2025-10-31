package resolution

import (
	"fmt"
	"runtime"
)

type Resolution struct {
	WindowWidth int32
	WindowHeight int32
	DrawWidth  int32
	DrawHeight int32
	DrawOffsetY int32
}

func ReadResolution() (Resolution) {
	macos := runtime.GOOS == "darwin"
	linux := runtime.GOOS == "linux"

	if !linux && !macos {
		err := fmt.Errorf("resolution: whether darwin nor linux detected: unsupported platform")
		panic(err)
	}

	return readPlatformResolution()
}

