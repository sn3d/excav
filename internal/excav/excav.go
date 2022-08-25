package excav

import "github.com/sn3d/excav/api"

// dispatcher is globally available for all  excav internals.
// Any excav part might emit event.
//
// Not sure if it's OK to have it as global butt currently
// it's more usable even globals are kind of antpattern.
var dispatcher = Dispatcher{}

func init() {
	dispatcher.listeners = make([]api.EventListener, 0)
}
