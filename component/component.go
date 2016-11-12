package component

import "go-component-entity-system/event"

//IComponent component interface
type IComponent interface {
	FireEvent(e event.Event) bool
}

//Component base component
type Component struct {
}

//FireEvent processes the event for the component
func (c *Component) FireEvent(e event.Event) bool {

	return false
}
