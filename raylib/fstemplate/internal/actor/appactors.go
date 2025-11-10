package actor

type AppActors struct {
	actors []Actor
}


func NewAppActors() *AppActors {
	a := &AppActors {
		actors: make([]Actor,0, 4),
	}
	return a
}

func (a *AppActors) Init(w int32, h int32) {
	for _, a := range(a.actors) {
		a.Init(w, h)
	}
}

func (a *AppActors) ReadInput() {
	for _, a := range(a.actors) {
		a.ReadInput()
	}
}

func (a *AppActors) UpdateState() {
	for _, a := range(a.actors) {
		a.UpdateState()
	}
}

func (a *AppActors) Render(w int32, h int32) {
	for _, a := range(a.actors) {
		a.Render(w, h)
	}
}

func (a *AppActors) Append(actor Actor){
	a.actors = append(a.actors, actor)
}
