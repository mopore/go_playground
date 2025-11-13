package actor

type CollectionActor struct {
	actors []Actor
}


func NewCollectionActor() *CollectionActor {
	a := &CollectionActor {
		actors: make([]Actor,0, 4),
	}
	return a
}

func (a *CollectionActor) Init(w int32, h int32) {
	for _, a := range(a.actors) {
		a.Init(w, h)
	}
}

func (a *CollectionActor) ReadInput() {
	for _, a := range(a.actors) {
		a.ReadInput()
	}
}

func (a *CollectionActor) UpdateState() []ActorRequest {
	var allReqs []ActorRequest
	for _, a := range(a.actors) {
		reqs := a.UpdateState()
		allReqs = append(allReqs, reqs...)
	}
	return allReqs
}

func (a *CollectionActor) Render(w int32, h int32) {
	for _, a := range(a.actors) {
		a.Render(w, h)
	}
}

func (a *CollectionActor) Append(actor Actor){
	a.actors = append(a.actors, actor)
}
