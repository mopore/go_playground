package render

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/mopore/go_playground/raylib/fstemplate/internal/actor"
	"github.com/mopore/go_playground/raylib/fstemplate/internal/resolution"
)

const (
	circleRadius = float32(20) // diameter = 40
	fontSize     = int32(40)
	renderedText = "This text has a size of 40"
)

func RenderLoop(res resolution.Resolution, actors []actor.Actor) {
	w := res.DrawWidth
	h := res.DrawHeight - res.DrawOffsetY

	textWidth := rl.MeasureText(renderedText, fontSize)
	textX := (w/2) - (textWidth/2)
	textY := (h/2) - (fontSize/2)

	for !rl.WindowShouldClose() {

		// TODO: Externalize the key input as an actor
		if rl.IsKeyPressed(rl.KeyQ) {
			break
		}
		for _, a := range(actors) {
			a.ReadInput()
		}
		for _, a := range(actors) {
			a.UpdateState()
		}

		rl.BeginDrawing()
		render(w, h, textX, textY)
		for _, a := range(actors) {
			a.Render()
		}
		rl.EndDrawing()
	}
}

func render(w int32, h int32, textX int32, textY int32) {
	rl.ClearBackground(rl.Yellow)

	rl.DrawCircle(int32(circleRadius), int32(circleRadius), circleRadius, rl.Red) // TL
	rl.DrawCircle(int32(w)-int32(circleRadius), int32(circleRadius), circleRadius, rl.Red)        // TR
	rl.DrawCircle(int32(circleRadius), int32(h)-int32(circleRadius), circleRadius, rl.Red)        // BL
	rl.DrawCircle(int32(w)-int32(circleRadius), int32(h)-int32(circleRadius), circleRadius, rl.Red)     // BR

	// debug text
	rl.DrawText(renderedText, textX, textY, fontSize, rl.Gray)
}

