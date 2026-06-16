package bot

type State string

const (
	StateUnknown         State = "unknown"
	StateWaitName        State = "wait_name"
	StateWaitBio         State = "wait_bio"
	StateBrowsing        State = "browsing"
	StateMenu            State = "menu"
	StateConfirmDeletion State = "confirm_deletion"
)

type FSM struct {
	states map[int64]State
}

func NewFSM() *FSM {
	return &FSM{
		states: make(map[int64]State),
	}
}

func (f *FSM) GetState(userID int64) State {
	state, ok := f.states[userID]
	if !ok {
		return StateUnknown
	}
	return state
}

func (f *FSM) SetState(userID int64, state State) {
	f.states[userID] = state
}
