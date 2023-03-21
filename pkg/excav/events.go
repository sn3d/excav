package excav

// You will implement this interface if you want listen and
// react to any event. You have to register your listener implementation
// to Dispatcher
type EventListener interface {
	OnEvent(ev Event)
}

// All events are represented by this interface. Yes
// currently any data struct can be event.
type Event interface {
}

//------------------------------------------------------------------------------
// Events
//------------------------------------------------------------------------------

type PatchingStarted struct {
	Repo string
}

type PatchApplied struct {
	Repo      string
	Branch    string
	CommitMsg string
	ErrorMsg  string
}

type TaskStarted struct {
	Task string
}

type TaskEnd struct {
	Task  string
	Error error
}

type Pushed struct {
	Repo            string
	MergeRequestURL string
	ErrorMsg        string
}

type ReposSelected struct {
	RepoNames []string
}

type RepoDiscarded struct {
	Repo  string
	Error error
}

type BulkDiscarded struct {
}

type RepoError struct {
	Repo     string
	ErrorMsg string
	Error    error
}
