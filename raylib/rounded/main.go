package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.SetConfigFlags(rl.FlagWindowUndecorated | rl.FlagWindowTransparent)
	rl.InitWindow(800, 450, "Fake Rounded Window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		// Clear fully transparent
		rl.ClearBackground(rl.Blank)

		// Draw rounded rectangle (the visible “window”)
		rl.DrawRectangleRounded(
			rl.NewRectangle(0, 0, 800, 450),
			0.05, // corner radius (5%)
			32,
			rl.Color{R: 245, G: 245, B: 245, A: 255},
		)

		rl.DrawText("Simulated rounded window", 220, 210, 20, rl.DarkGray)

		rl.EndDrawing()
	}
}

