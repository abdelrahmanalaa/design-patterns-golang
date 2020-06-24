package main

type State interface {
	switchMe(*StateContext)
}

type StateContext struct {
	state State
}

func (s *StateContext) setState(newState State) {
	s.state = newState
}

func (s *StateContext) switchMe() {
	s.state.switchMe(s)
}

type On struct {
}

func (o *On) switchMe(s *StateContext) {
	println("state is on")
	s.setState(new(Off))
}

type Off struct {
}

func (o *Off) switchMe(s *StateContext) {
	println("state is off")
	s.setState(new(On))
}

func createNewContext() *StateContext {
	return &StateContext{state: new(Off)}
}

func main() {

	stateContextInstance := createNewContext()

	stateContextInstance.switchMe()
	stateContextInstance.switchMe()
	stateContextInstance.switchMe()

}
