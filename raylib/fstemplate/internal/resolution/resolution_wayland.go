//go:build linux

package resolution

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const  (
	regularOffsetY = int32(35)
)

func readPlatformResolution() Resolution {
	rl.InitWindow(1, 1, "prescreen")

	monitor := rl.GetCurrentMonitor()
	monWidth := int32(rl.GetMonitorWidth(monitor))
	monHeight := int32(rl.GetMonitorHeight(monitor))

	rWidth := int32(rl.GetRenderWidth())
	// rHeight := int32(rl.GetRenderHeight())

	rl.CloseWindow()

	// When 200%
	// Monitor width 3600, height 2252
	// Render width 7200, height 4576
	// We use 1800, 1126

	// When 100%
	// Monitor width 3600, height 2252
	// Render width 3600, height 2288 
	// Result we use 3600 and 2252

	// mtext := fmt.Sprintf("Monitor width %v, height %v", monWidth, monHeight)
	// rtext := fmt.Sprintf("Render width %v, height %v", rWidth, rHeight)
	//
	// log.Println(mtext)
	// log.Println(rtext)

	if rWidth == 0 || monWidth == 0 {
		errMsg := fmt.Sprintf("Main: could not get a valid reading. renderWidth is \"%d\", monWidth is \"%d\"", rWidth, monWidth)
		panic(errMsg)
	}
	scale := rWidth / monWidth

	if scale == 0 {
		errMsg := fmt.Sprintf("Main: scale is zero. renderWidth is \"%d\", monWidth is \"%d\"", rWidth, monWidth)
		panic(errMsg)
	}

	resWidth := monWidth / scale
	resHeight := monHeight / scale
	offsetY := regularOffsetY / scale

	return Resolution{
		WindowWidth: monWidth,
		WindowHeight: monHeight,
		DrawWidth:  resWidth,
		DrawHeight: resHeight,
		DrawOffsetY: offsetY,
	}
}
