package excav

import "github.com/sn3d/excav/api"

// Dispatcher is responsible for delivery events to registered
// listeners. The code is sending events via Notify and reacting
// on events by implementing EventListener.
type Dispatcher struct {
	listeners []api.EventListener
}

func RegisterListener(listener api.EventListener) {
	dispatcher.listeners = append(dispatcher.listeners, listener)
}

// Send event
func (d *Dispatcher) Notify(ev api.Event) {
	if d == nil {
		return
	}

	for _, listener := range d.listeners {
		listener.OnEvent(ev)
	}
}
