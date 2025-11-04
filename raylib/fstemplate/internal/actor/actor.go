package actor


type Actor interface {
	ReadInput()
	UpdateState()
	Render()
}
