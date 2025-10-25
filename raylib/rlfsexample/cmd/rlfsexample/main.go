package main

import (
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	message     = "press 'q' to quit"
	fontSize    = 40
	cycleTime   = 2.0 // seconds for full red→blue→red cycle
)

func main() {
	screenWidth := int32(rl.GetMonitorWidth(0))
	screenHeight := int32(rl.GetMonitorHeight(0))

	rl.InitWindow(screenWidth, screenHeight, "Fullscreen Example")
	rl.ToggleFullscreen()
	rl.HideCursor()

	defer func() {
		rl.CloseWindow()
		rl.ShowCursor()
	}()

	startTime := time.Now()

	renderW := rl.GetRenderWidth()
	renderH := rl.GetRenderHeight()

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		if shouldQuit() {
			break
		}

		color := messageColor(startTime)
		drawFrame(int32(renderW), int32(renderH), color)
	}

	exitFullscreen()
}

// shouldQuit handles input logic cleanly
func shouldQuit() bool {
	return rl.IsKeyPressed(rl.KeyQ)
}

// messageColor returns the interpolated color between red and blue
func messageColor(start time.Time) rl.Color {
	elapsed := time.Since(start).Seconds()
	t := math.Sin((elapsed/cycleTime)*math.Pi*2)*0.5 + 0.5 // 0..1
	r := uint8(255 * (1 - t))
	b := uint8(255 * t)
	return rl.NewColor(r, 0, b, 255)
}

// drawFrame is responsible only for drawing
func drawFrame(renderW, renderH int32, color rl.Color) {

	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.Black)

	textWidth := rl.MeasureText(message, fontSize)
	x := renderW/2 - textWidth/2
	y := renderH/2 - fontSize/2

	rl.DrawText(message, x, y, fontSize, color)
}

// exitFullscreen ensures graceful teardown
func exitFullscreen() {
	if rl.IsWindowFullscreen() {
		rl.ToggleFullscreen()
	}
}
