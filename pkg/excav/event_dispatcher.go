package excav

// dispatcher is globally available for all  excav internals.
// Any excav part might emit event.
//
// Not sure if it's OK to have it as global butt currently
// it's more usable even globals are kind of antpattern.
var dispatcher = Dispatcher{}

func init() {
	dispatcher.listeners = make([]EventListener, 0)
}

// Dispatcher is responsible for delivery events to registered
// listeners. The code is sending events via Notify and reacting
// on events by implementing EventListener.
type Dispatcher struct {
	listeners []EventListener
}

func RegisterListener(listener EventListener) {
	dispatcher.listeners = append(dispatcher.listeners, listener)
}

// Send event
func (d *Dispatcher) Notify(ev Event) {
	if d == nil {
		return
	}

	for _, listener := range d.listeners {
		listener.OnEvent(ev)
	}
}
