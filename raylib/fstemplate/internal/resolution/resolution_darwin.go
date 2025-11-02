//go:build darwin

package resolution

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
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func readPlatformResolution() Resolution {
	var w, h C.int
	var scale C.double
	C.getMainDisplayPixels(&w, &h, &scale)
	fmt.Printf("Main display: %dx%d px (backing scale %.2f)\n", int(w), int(h), float64(scale))
	scaledW := int32(float32(w)/float32(scale))
	scaledH := int32(float32(h)/float32(scale))

	rl.InitWindow(scaledW, scaledH, "prescreen")

	monitor := rl.GetCurrentMonitor()
	monWidth := int32(rl.GetMonitorWidth(monitor))
	monHeight := int32(rl.GetMonitorHeight(monitor))

	renderWidth := int32(rl.GetRenderWidth())
	renderHeight := int32(rl.GetRenderHeight())

	rl.CloseWindow()

	mtext := fmt.Sprintf("resolution: monitor width %v, height %v", monWidth, monHeight)
	rtext := fmt.Sprintf("resolution: render width %v, height %v", renderWidth, renderHeight)

	log.Println(mtext)  // resolution: monitor width 1920, height 1200
	log.Println(rtext)  // resolution: render width 1800, height 1169

	if renderWidth == 0 || monWidth == 0 {
		errMsg := fmt.Sprintf("resolution: could not get a valid reading. renderWidth is \"%d\", monWidth is \"%d\"", renderWidth, monWidth)
		panic(errMsg)
	}
	cscale := float32(renderWidth) / float32(monWidth)

	return Resolution{
		WindowWidth: monWidth,
		WindowHeight: monHeight,
		DrawWidth:  monWidth,
		DrawHeight: monHeight,
		DrawOffsetY: 0,
		Scale : float32(cscale),
	}
}
