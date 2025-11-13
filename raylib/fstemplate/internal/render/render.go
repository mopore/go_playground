package render

import (
	"fmt"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/mopore/go_playground/raylib/fstemplate/internal/actor"
	"github.com/mopore/go_playground/raylib/fstemplate/internal/resolution"
)

func RenderLoop(res resolution.Resolution, a actor.Actor) {
	w := res.DrawWidth
	h := res.DrawHeight - res.DrawOffsetY

	a.Init(w, h)

	for !rl.WindowShouldClose() {

		// TODO: Externalize the key input as an actor
		if rl.IsKeyPressed(rl.KeyQ) {
			break
		}
		a.ReadInput()

		quitApp := false
		if reqs := a.UpdateState(); reqs != nil {
			for _, r := range reqs {
				switch t := r.(type) {
				case actor.QuitActorRequest:
					log.Println("RenderLoop: quit request, reason:", r.Reason())
					quitApp = true
				default:
					err := fmt.Errorf("RenderLoop: unknown AppRequest \"%v\"", t)
					panic(err)
				}
			}
		}
		if quitApp {
			break
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		a.Render(w, h)
		rl.EndDrawing()
	}
}
