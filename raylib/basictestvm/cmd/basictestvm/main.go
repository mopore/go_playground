package main

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	opengl21Tag = "-DGRAPHICS_API_OPENGL_21"
	regularOffsetY = int32(35)
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
	if value, ok := os.LookupEnv("CGO_CFLAGS"); ok {
		if value == opengl21Tag{
			text = "At Arch VM on M2 Max"
		} else {
			text = value
		}
	}

	drawLoop(res.drawWidth, res.drawHeight - res.drawOffsetY, text)
}

func readResolution() resolution {
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

	scale := rWidth / monWidth

	resWidth := monWidth / scale
	resHeight := monHeight / scale
	offsetY := regularOffsetY / scale

	return resolution{
		windowWidth: monWidth,
		windowHeight: monHeight,
		drawWidth:  resWidth,
		drawHeight: resHeight,
		drawOffsetY: offsetY,
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
