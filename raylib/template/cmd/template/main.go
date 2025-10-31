package main

import (

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/mopore/go_playground/raylib/template/internal/resolution"
)

const (
	circleRadius = float32(20) // diameter = 40
	fontSize     = int32(40)
)

func main() {
	rl.SetConfigFlags(
		rl.FlagFullscreenMode |
			rl.FlagVsyncHint |
			rl.FlagWindowHighdpi |
			rl.FlagWindowTopmost,
	)

	res := resolution.ReadResolution()

	rl.HideCursor()
	rl.SetTargetFPS(60)

	rl.InitWindow(res.WindowWidth, res.WindowHeight, "Raylib test")
	defer rl.CloseWindow()

	drawLoop(res.DrawWidth, res.DrawHeight - res.DrawOffsetY)
}


func drawLoop(w int32, h int32) {
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
		rl.DrawText("here to party", w/2, h/2, fontSize, rl.Gray)

		rl.EndDrawing()
	}
}
