package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	const (
		circleRadius = float32(20) // diameter = 40
		fontSize     = int32(20)
	)

	// ---- Configure before InitWindow ----
	rl.SetConfigFlags(
		rl.FlagFullscreenMode |
			rl.FlagVsyncHint |
			rl.FlagWindowHighdpi |
			rl.FlagWindowTopmost,
	)

	w := int32(3600)
	h := int32(2252)

	// Size ignored in fullscreen
	rl.InitWindow(w, h, "raylib-go fullscreen")
	defer rl.CloseWindow()

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

	rl.HideCursor()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeyQ) {
			break
		}

		// Use framebuffer (render) size – real pixels, not logical 800×600

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.DrawRectangle(0, h-100, 100, 100, rl.White)

		// Four red filled circles in corners
		rl.DrawCircle(int32(circleRadius), int32(circleRadius), circleRadius, rl.Red) // TL
		rl.DrawCircle(int32(w)-int32(circleRadius), int32(circleRadius), circleRadius, rl.Red)        // TR
		rl.DrawCircle(int32(circleRadius), int32(h)-int32(circleRadius), circleRadius, rl.Red)        // BL
		rl.DrawCircle(int32(w)-int32(circleRadius), int32(h)-int32(circleRadius), circleRadius, rl.Red)     // BR

		// debug text
		rl.DrawText(mtext, 100, 50, fontSize, rl.White)
		rl.DrawText(rtext, 100, 100, fontSize, rl.White)
		rl.DrawText(stext, 100, 150, fontSize, rl.White)

		rl.EndDrawing()
	}
}
