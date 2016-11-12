package object

import (
	"go-component-entity-system/component"
	"go-component-entity-system/event"
)

//Object Base game object that contains all the components and blueprint builder
type Object struct {
	components []component.IComponent
}

func (o *Object) fireEvent(e event.Event) bool {
	for i := range o.components {
		c := o.components[i]
		if !c.FireEvent(e) {
			break
		}
	}
	return false
}
