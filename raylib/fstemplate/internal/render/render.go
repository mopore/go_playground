package render

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/mopore/go_playground/raylib/fstemplate/internal/actor"
	"github.com/mopore/go_playground/raylib/fstemplate/internal/resolution"
)

const (
	circleRadius = float32(20) // diameter = 40
	fontSize     = int32(40)
)

func RenderLoop(res resolution.Resolution, actor actor.Actor) {
	w := res.DrawWidth
	h := res.DrawHeight - res.DrawOffsetY

	for !rl.WindowShouldClose() {

		// TODO: Externalize the key input
		if rl.IsKeyPressed(rl.KeyQ) {
			break
		}
		actor.ReadInput()
		actor.UpdateState()

		rl.BeginDrawing()
		render(w,h)
		actor.Render()
		rl.EndDrawing()
	}
}

func render(w int32, h int32) {
	rl.ClearBackground(rl.Yellow)

	rl.DrawCircle(int32(circleRadius), int32(circleRadius), circleRadius, rl.Red) // TL
	rl.DrawCircle(int32(w)-int32(circleRadius), int32(circleRadius), circleRadius, rl.Red)        // TR
	rl.DrawCircle(int32(circleRadius), int32(h)-int32(circleRadius), circleRadius, rl.Red)        // BL
	rl.DrawCircle(int32(w)-int32(circleRadius), int32(h)-int32(circleRadius), circleRadius, rl.Red)     // BR

	// debug text
	rl.DrawText("here to shine", w/2, h/2, fontSize, rl.Gray)
}

