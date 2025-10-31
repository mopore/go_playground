package main

/*
#cgo darwin CFLAGS: -x objective-c
#cgo darwin LDFLAGS: -framework CoreGraphics -framework AppKit
#ifdef __APPLE__
#import <CoreGraphics/CoreGraphics.h>
#import <AppKit/NSScreen.h>

static void getMainDisplayPixels(int *w, int *h, double *scale) {
    CGDirectDisplayID did = CGMainDisplayID();
    CGDisplayModeRef mode = CGDisplayCopyDisplayMode(did);
    if (!mode) {
        *w = 0; *h = 0; *scale = 1.0;
        return;
    }

    size_t pw = CGDisplayModeGetPixelWidth(mode);
    size_t ph = CGDisplayModeGetPixelHeight(mode);
    *w = (int)pw;
    *h = (int)ph;

    NSScreen *s = [NSScreen mainScreen];
    *scale = s ? [s backingScaleFactor] : 1.0;

    CGDisplayModeRelease(mode);
}
#endif
*/
import "C"

import (
	"fmt"
	"runtime"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	circleRadius = float32(20) // diameter = 40
	fontSize     = int32(40)
)

type resolution struct {
	windowWidth int32
	windowHeight int32
	drawWidth  int32
	drawHeight int32
	drawOffsetY int32
}


func main() {

	if runtime.GOOS != "darwin" {
		fmt.Println("This only runs with MacOS")
		return
	}

	// ---- Configure before InitWindow ----
	rl.SetConfigFlags(
		rl.FlagFullscreenMode |
			rl.FlagVsyncHint |
			rl.FlagWindowHighdpi |
			rl.FlagWindowTopmost,
	)

	res := readResolution()

	rl.InitWindow(res.windowWidth, res.windowHeight, "Raylib test")
	defer rl.CloseWindow()

	rl.HideCursor()
	rl.SetTargetFPS(60)

	text := "Regular Environment"

	drawLoop(res.drawWidth, res.drawHeight - res.drawOffsetY, text)
}

func readResolution() resolution {
	var w, h C.int
	var scale C.double
	C.getMainDisplayPixels(&w, &h, &scale)
	fmt.Printf("Main display: %dx%d px (backing scale %.2f)\n", int(w), int(h), float64(scale))
	scaledW := int32(float32(w)/float32(scale))
	scaledH := int32(float32(h)/float32(scale))

	return resolution{
		windowWidth: scaledW,
		windowHeight: scaledH,
		drawWidth:  scaledW,
		drawHeight: scaledH,
		drawOffsetY: 0,
	}
}


func drawLoop(w int32, h int32, text string) {
	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeyQ) {
			break
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Yellow)

		rl.DrawCircle(int32(circleRadius), int32(circleRadius), circleRadius, rl.Red) // TL
		rl.DrawCircle(int32(w)-int32(circleRadius), int32(circleRadius), circleRadius, rl.Red)        // TR
		rl.DrawCircle(int32(circleRadius), int32(h)-int32(circleRadius), circleRadius, rl.Red)        // BL
		rl.DrawCircle(int32(w)-int32(circleRadius), int32(h)-int32(circleRadius), circleRadius, rl.Red)     // BR

		// debug text
		rl.DrawText(text, w/2, h/2, fontSize, rl.Gray)

		rl.EndDrawing()
	}
}
