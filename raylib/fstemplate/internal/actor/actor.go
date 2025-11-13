package actor

type ActorRequest interface {
	Reason() string
}


type QuitActorRequest struct{
	reason string
}

func (q QuitActorRequest) Reason() string {
	return q.reason
}

func NewQuitActorRequest(reason string) QuitActorRequest{
	return QuitActorRequest{reason: reason}
}

type Actor interface {
	Init(w int32, h int32)
	ReadInput()
	UpdateState() []ActorRequest
	Render(w int32, h int32)
}
