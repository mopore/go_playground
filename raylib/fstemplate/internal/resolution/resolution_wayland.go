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

	scaleX := rl.GetWindowScaleDPI().X
	scaleY := rl.GetWindowScaleDPI().Y

	rl.CloseWindow()

	mtext := fmt.Sprintf("resoltion: monitor width %v, height %v", monWidth, monHeight)
	rtext := fmt.Sprintf("resolution: render width %v, height %v", renderWidth, renderHeight)
	dpitext := fmt.Sprintf("resolution: dpi x %v, dpi y %v", scaleX, scaleY)
	//
	// When 200% -- Only before Gnome 49 there could be uneven values
	// Monitor width 3600, height 2252
	// Render width 7200, height 4576 (4572/2 == 2288 -> Offset of 36)
	//
	// When 100%
	// Monitor width 3600, height 2252
	// Render width 3600, height 2288 
	//
	// ScaleDPI delivers 1.333 for 133% and so...

	log.Println(mtext)
	log.Println(rtext)
	log.Println(dpitext)

	if renderWidth == 0 || monWidth == 0 {
		errMsg := fmt.Sprintf("resolution: could not get a valid reading. renderWidth is \"%d\", monWidth is \"%d\"", renderWidth, monWidth)
		panic(errMsg)
	}

	winWidth := int32(float32(monWidth) / scaleX)
	winHeight := int32(float32(monHeight) / scaleY)

	// Before Gnome 49 Wayland gave a different monitor and render height resulting in the
	// need to provide a vertical offset for correct rendering.
	scaledMonHeight := int32(float32(monHeight) / scaleY)		
	scaledRenderHeight := int32(float32(renderHeight) / scaleY)
	drawOffsetY := scaledRenderHeight - scaledMonHeight
	if drawOffsetY > 0 {
		drawOffsetY = drawOffsetY - 1
	}

	drawWidth := int32(float32(monWidth) / scaleX)
	drawHeight := int32(float32(monHeight) / scaleY)

	res:= Resolution{
		WindowWidth: winWidth,
		WindowHeight: winHeight,
		DrawWidth:  drawWidth,
		DrawHeight: drawHeight,
		DrawOffsetY: drawOffsetY,
		Scale: scaleX,
	}
	log.Printf("resolution: calculated resolution: %v\n", res)

	return res
}
