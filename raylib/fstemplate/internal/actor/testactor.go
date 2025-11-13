package actor

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	circleRadius = float32(20) // diameter = 40
	fontSize     = int32(40)
	renderedText = "press 't' to toggle test rendering"
)

type TestActor struct {
	active bool
	textX int32
	textY int32
}

func NewTestActor() *TestActor {
	a := &TestActor {
		active: true,
	}
	return a
}

func (a *TestActor) Init(w int32, h int32) {
	textWidth := rl.MeasureText(renderedText, fontSize)
	a.textX = (w/2) - (textWidth/2)
	a.textY = (h/2) - (fontSize/2)
}

func (a *TestActor) ReadInput() {
	if rl.IsKeyPressed(rl.KeyT) {
		a.active = !a.active
	}
}

func (a *TestActor) UpdateState() []ActorRequest {
	return nil
}


func (a *TestActor) Render(w int32, h int32) {
	if a.active {
		rl.ClearBackground(rl.Yellow)

		rl.DrawCircle(int32(circleRadius), int32(circleRadius), circleRadius, rl.Red) // TL
		rl.DrawCircle(int32(w)-int32(circleRadius), int32(circleRadius), circleRadius, rl.Red)        // TR
		rl.DrawCircle(int32(circleRadius), int32(h)-int32(circleRadius), circleRadius, rl.Red)        // BL
		rl.DrawCircle(int32(w)-int32(circleRadius), int32(h)-int32(circleRadius), circleRadius, rl.Red)     // BR

		// debug text
		rl.DrawText(renderedText, a.textX, a.textY, fontSize, rl.Gray)
	}
}

