package main

import (
	"fmt"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// ---- Configure before InitWindow ----
	rl.SetConfigFlags(
		rl.FlagFullscreenMode |
			rl.FlagVsyncHint |
			rl.FlagWindowHighdpi |
			rl.FlagWindowTopmost,
	)

	w, h := readResolution()

	rl.InitWindow(w, h, "prescreen")
	rl.HideCursor()
	rl.SetTargetFPS(60)

	drawLoop(w, h-35)
}

func readResolution() (int32, int32) {
	rl.InitWindow(1, 1, "prescreen")

	monitor := rl.GetCurrentMonitor()
	monWidth := int32(rl.GetMonitorWidth(monitor))
	monHeight := int32(rl.GetMonitorHeight(monitor))
	mtext := fmt.Sprintf("Monitor width %v, height %v", monWidth, monHeight)

	rWidth := rl.GetRenderWidth()
	rHeight := rl.GetRenderHeight()
	rtext := fmt.Sprintf("Render width %v, height %v", rWidth, rHeight)

	sWidth := rl.GetScreenWidth()
	sHeight := rl.GetScreenHeight()
	stext := fmt.Sprintf("Screen width %v, height %v", sWidth, sHeight)

	log.Println(mtext)
	log.Println(rtext)
	log.Println(stext)

	rl.CloseWindow()

	return monWidth, monHeight
}


func drawLoop(w int32, h int32) {
	const (
		circleRadius = float32(20) // diameter = 40
		fontSize     = int32(40)
	)

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeyQ) {
			break
		}

		// Use framebuffer (render) size – real pixels, not logical 800×600

		rl.BeginDrawing()
		rl.ClearBackground(rl.Yellow)

		// rl.DrawRectangle(0, h-100, 100, 100, rl.Red)

		// Four red filled circles in corners
		rl.DrawCircle(int32(circleRadius), int32(circleRadius), circleRadius, rl.Red) // TL
		rl.DrawCircle(int32(w)-int32(circleRadius), int32(circleRadius), circleRadius, rl.Red)        // TR
		rl.DrawCircle(int32(circleRadius), int32(h)-int32(circleRadius), circleRadius, rl.Red)        // BL
		rl.DrawCircle(int32(w)-int32(circleRadius), int32(h)-int32(circleRadius), circleRadius, rl.Red)     // BR

		// debug text
		rl.DrawText("Hello!", w/2, h/2, fontSize, rl.Gray)
		// rl.DrawText(mtext, 100, 50, fontSize, rl.White)
		// rl.DrawText(rtext, 100, 100, fontSize, rl.White)
		// rl.DrawText(stext, 100, 150, fontSize, rl.White)

		rl.EndDrawing()
	}
}
