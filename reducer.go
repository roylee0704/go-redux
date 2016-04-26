package main

// Action is an object that holds an action_type
type Action struct {
	typ string
}

// counter is a reducer function that accepts state, action and
// returns new states.
func counter(state int, action *Action) int {

	switch action.typ {
	case "INCREMENT":
		return state + 1
	case "DECREMENT":
		return state - 1
	default:
		return state
	}

}
