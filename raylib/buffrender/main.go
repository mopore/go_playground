package main

import (
	"runtime"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// Raylib requires the main thread for OpenGL context, esp. on macOS
	runtime.LockOSThread()

	const (
		screenWidth  = 800
		screenHeight = 450
	)

	rl.InitWindow(screenWidth, screenHeight, "Buffered Rendering Demo (Go)")
	defer rl.CloseWindow()

	// Create the offscreen render target
	target := rl.LoadRenderTexture(screenWidth, screenHeight)
	defer rl.UnloadRenderTexture(target)

	// Optional: enable bilinear filtering for smoother scaling
	rl.SetTextureFilter(target.Texture, rl.FilterBilinear)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		// 1️⃣ Draw everything to offscreen buffer
		rl.BeginTextureMode(target)
		{
			rl.ClearBackground(rl.RayWhite)
			rl.DrawText("hello world", 250, 200, 40, rl.Black)
			rl.DrawRectangleLinesEx(rl.NewRectangle(0, 0, float32(screenWidth), float32(screenHeight)), 10, rl.DarkGray)
		}
		rl.EndTextureMode()

		// Draw the buffered texture to the actual screen
		// 
		// This would also leave room for post processing
		//
		rl.BeginDrawing()
		{
			rl.ClearBackground(rl.RayWhite)

			// NOTE: When drawing textures from RenderTexture2D, Y axis is flipped (OpenGL thing)
			source := rl.NewRectangle(0, 0, float32(target.Texture.Width), -float32(target.Texture.Height))
			dest := rl.NewRectangle(0, 0, float32(screenWidth), float32(screenHeight))
			origin := rl.NewVector2(0, 0)

			rl.DrawTexturePro(target.Texture, source, dest, origin, 0.0, rl.White)
		}
		rl.EndDrawing()
	}
}
