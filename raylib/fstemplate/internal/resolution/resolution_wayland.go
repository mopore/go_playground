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

	dpiWidth := rl.GetWindowScaleDPI().X
	dpiHeight := rl.GetWindowScaleDPI().Y

	rl.CloseWindow()

	// When 200% -- Only before Gnome 49
	// Monitor width 3600, height 2252
	// Render width 7200, height 4576
	// We use 1800, 1126

	// When 100%
	// Monitor width 3600, height 2252
	// Render width 3600, height 2288 
	// Result we use 3600 and 2252

	// With Gnome 49+
	// ScaleDPI delivers 1.333 for 133% and so

	mtext := fmt.Sprintf("resoltion: monitor width %v, height %v", monWidth, monHeight)
	rtext := fmt.Sprintf("resolution: render width %v, height %v", renderWidth, renderHeight)
	dpitext := fmt.Sprintf("resolution: dpi x %v, dpi y %v", dpiWidth, dpiHeight)

	log.Println(mtext)
	log.Println(rtext)
	log.Println(dpitext)

	if renderWidth == 0 || monWidth == 0 {
		errMsg := fmt.Sprintf("resolution: could not get a valid reading. renderWidth is \"%d\", monWidth is \"%d\"", renderWidth, monWidth)
		panic(errMsg)
	}
	// TODO: Check if we can use ScaleDPI with Gnome < 49 instead of calculated scale
	// calcScale := float32(renderWidth) / float32(monWidth)
	calcScale := float32(renderWidth) / float32(monWidth)

	if calcScale == 0 {
		errMsg := fmt.Sprintf("resolution: scale is zero. renderWidth is \"%d\", monWidth is \"%d\"", renderWidth, monWidth)
		panic(errMsg)
	}

	drawWidth := int32(float32(monWidth) / calcScale)
	drawHeight := int32(float32(monHeight) / calcScale)

	// Before Gnome 49 Wayland gave a different monitor and render height resulting in the
	// need to provide a vertical offset for correct rendering.
	scaledRenderHeight := int32(float32(renderHeight) / calcScale)
	drawOffsetY := scaledRenderHeight - monHeight - 1
	if drawOffsetY == 1 {
		drawOffsetY = 0
	}

	res:= Resolution{
		WindowWidth: monWidth,
		WindowHeight: monHeight,
		DrawWidth:  drawWidth,
		DrawHeight: drawHeight,
		DrawOffsetY: drawOffsetY,
		Scale: calcScale,
	}
	log.Printf("resolution: calculated resolution: %v\n", res)

	return res
}
