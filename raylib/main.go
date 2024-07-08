package main

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	log.Printf("Testing raylib-go")
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("This is a test for raylib", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
}
