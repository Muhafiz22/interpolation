package parsing

import "fmt"

// State string in terms of FSM
type State string

// Event or function to be performed
type Event string

type Action struct {
	Destination State
}

type Transition map[Event]Action

type StateMap map[State]Transition

type FSM struct {
	InitialState State
	CurrentState State
	StateMap     StateMap
}

type IFSM interface {
	Current() State
	Transition(event Event) error
}

func (fsm *FSM) Current() State {
	if fsm.CurrentState == "" {
		return fsm.InitialState
	}
	return fsm.CurrentState
}

func (fsm *FSM) Transition(event Event) error {
	action := fsm.StateMap[fsm.CurrentState][event]
	if fmt.Sprint(action) != fmt.Sprint(Action{}) {
		fsm.CurrentState = action.Destination
		return nil
	}
	return fmt.Errorf("transition invalid")

}
