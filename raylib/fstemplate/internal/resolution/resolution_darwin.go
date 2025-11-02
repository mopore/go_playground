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
)

const  (
	regularOffsetY = int32(35)
)

func readPlatformResolution() Resolution {
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
		Scale : float32(scale),
	}
}
