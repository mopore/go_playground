package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Note: When running in Arch VM under MAC remember to set:
// export CGO_CFLAGS=\"-DGRAPHICS_API_OPENGL_21\"
//
// You also want to set your scaling to 100%


func main() {

	const (
		screenWidth  = int32(800)
		screenHeight = int32(600)
		circleRadius = float32(20) // diameter = 40
	)

	rl.InitWindow(screenWidth, screenHeight, "Raylib Go Example")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		// Check if 'q' is pressed
		if rl.IsKeyPressed(rl.KeyQ) {
			break
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// Draw four red filled circles at the corners
		rl.DrawCircle(int32(circleRadius), int32(circleRadius), circleRadius, rl.Red)                               // top-left
		rl.DrawCircle(screenWidth-int32(circleRadius), int32(circleRadius), circleRadius, rl.Red)                    // top-right
		rl.DrawCircle(int32(circleRadius), screenHeight-int32(circleRadius), circleRadius, rl.Red)                   // bottom-left
		rl.DrawCircle(screenWidth-int32(circleRadius), screenHeight-int32(circleRadius), circleRadius, rl.Red)       // bottom-right

		// Draw centered text
		text := "press 'q' to quit"
		fontSize := int32(20)
		textWidth := rl.MeasureText(text, fontSize)
		rl.DrawText(
			text,
			(screenWidth-textWidth)/2,
			(screenHeight-fontSize)/2,
			fontSize,
			rl.White,
		)

		rl.EndDrawing()
	}
}

