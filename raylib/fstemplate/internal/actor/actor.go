package actor


type Actor interface {
	Init(w int32, h int32)
	ReadInput()
	UpdateState()
	Render(w int32, h int32)
}
