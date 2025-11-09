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
			0.15, // corner radius (5%)
			32,
			rl.Color{R: 145, G: 145, B: 145, A: 255},
		)

		rl.DrawText("Simulated rounded window", 180, 210, 40, rl.DarkPurple)

		rl.EndDrawing()
	}
}

