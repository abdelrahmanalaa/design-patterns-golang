package main

type State interface {
	switchMe(*stateContext)
}

type stateContext struct {
	state State
}

func createNew() *stateContext {
	return &stateContext{state: new(Off)}
}

func (s *stateContext) setState(newState State) {
	s.state = newState
}

func main() {
	stateContextInstance := createNew()

	stateContextInstance.switchMe()
	stateContextInstance.switchMe()
	stateContextInstance.switchMe()

}

func (s *stateContext) switchMe() {
	s.state.switchMe(s)
}

type On struct {
}

func (o *On) switchMe(s *stateContext) {
	println("state is on")
	s.setState(new(Off))
}

type Off struct {
}

func (o *Off) switchMe(s *stateContext) {
	println("state is off")
	s.setState(new(On))
}
