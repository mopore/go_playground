//go:build linux

package resolution

import (
	"fmt"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func readPlatformResolution() Resolution {
	rl.InitWindow(1, 1, "prescreen")

	monitor := rl.GetCurrentMonitor()
	monWidth := int32(rl.GetMonitorWidth(monitor))
	monHeight := int32(rl.GetMonitorHeight(monitor))

	renderWidth := int32(rl.GetRenderWidth())
	renderHeight := int32(rl.GetRenderHeight())

	rl.CloseWindow()

	// When 200%
	// Monitor width 3600, height 2252
	// Render width 7200, height 4576
	// We use 1800, 1126

	// When 100%
	// Monitor width 3600, height 2252
	// Render width 3600, height 2288 
	// Result we use 3600 and 2252

	mtext := fmt.Sprintf("resoltion: monitor width %v, height %v", monWidth, monHeight)
	rtext := fmt.Sprintf("resolution: render width %v, height %v", renderWidth, renderHeight)

	log.Println(mtext)
	log.Println(rtext)

	if renderWidth == 0 || monWidth == 0 {
		errMsg := fmt.Sprintf("resolution: could not get a valid reading. renderWidth is \"%d\", monWidth is \"%d\"", renderWidth, monWidth)
		panic(errMsg)
	}
	scale := float32(renderWidth) / float32(monWidth)

	if scale == 0 {
		errMsg := fmt.Sprintf("resolution: scale is zero. renderWidth is \"%d\", monWidth is \"%d\"", renderWidth, monWidth)
		panic(errMsg)
	}

	resWidth := int32(float32(monWidth) / scale)
	resHeight := int32(float32(monHeight) / scale)

	// Before Gnome 49 Wayland gave a different monitor and render height resulting in the
	// need to provide a vertical offset for correct rendering.
	scaledRenderHeight := int32(float32(renderHeight) / scale)
	offsetY := scaledRenderHeight - monHeight - 1
	if offsetY == 1 {
		offsetY = 0
	}

	log.Println("resolution: calculated offset:", offsetY)

	return Resolution{
		WindowWidth: monWidth,
		WindowHeight: monHeight,
		DrawWidth:  resWidth,
		DrawHeight: resHeight,
		DrawOffsetY: offsetY,
		Scale: float32(scale),
	}
}
