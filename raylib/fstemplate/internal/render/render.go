package render

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/mopore/go_playground/raylib/fstemplate/internal/actor"
	"github.com/mopore/go_playground/raylib/fstemplate/internal/resolution"
)

func RenderLoop(res resolution.Resolution, actor actor.Actor) {
	w := res.DrawWidth
	h := res.DrawHeight - res.DrawOffsetY

	actor.Init(w, h)

	for !rl.WindowShouldClose() {

		// TODO: Externalize the key input as an actor
		if rl.IsKeyPressed(rl.KeyQ) {
			break
		}
		actor.ReadInput()

		actor.UpdateState()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		actor.Render(w, h)
		rl.EndDrawing()
	}
}
