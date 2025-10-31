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
	"runtime"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Resolution struct {
	WindowWidth int32
	WindowHeight int32
	DrawWidth  int32
	DrawHeight int32
	DrawOffsetY int32
}


const  (
	regularOffsetY = int32(35)
)

func ReadResolution() (Resolution) {
	macos := runtime.GOOS == "darwin"
	linux := runtime.GOOS == "linux"
	if macos {
		res := readMacResolution()
		return res
	} else if linux {
		res := readWaylandResolution()
		return res
	} else {
		err := fmt.Errorf("resolution: whether darwin nor linux detected: unsupported platform")
		panic(err)
	}
}

func readMacResolution() Resolution {
	var w, h C.int
	var scale C.double
	C.getMainDisplayPixels(&w, &h, &scale)
	fmt.Printf("Main display: %dx%d px (backing scale %.2f)\n", int(w), int(h), float64(scale))
	scaledW := int32(float32(w)/float32(scale))
	scaledH := int32(float32(h)/float32(scale))

	return Resolution{
		WindowWidth: scaledW,
		WindowHeight: scaledH,
		DrawWidth:  scaledW,
		DrawHeight: scaledH,
		DrawOffsetY: 0,
	}
}

func readWaylandResolution() Resolution {
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
