package main

import (
	"runtime"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/mopore/go_playground/raylib/fstemplate/internal/render"
	"github.com/mopore/go_playground/raylib/fstemplate/internal/resolution"
)

const (
	appName = "Raylib Fullscreen Template"
)

func main() {
	runtime.LockOSThread()

	rl.SetConfigFlags(
		rl.FlagFullscreenMode |
			rl.FlagVsyncHint |
			rl.FlagWindowHighdpi |
			rl.FlagWindowTopmost |
			rl.FlagMsaa4xHint, // Smoother rendering
	)

	res := resolution.ReadResolution()

	rl.InitWindow(res.WindowWidth, res.WindowHeight, appName)
	defer rl.CloseWindow()

	rl.HideCursor()
	rl.SetTargetFPS(60)

	render.RenderLoop(res)
}
