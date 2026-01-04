package ui

import ecs "github.com/igadmg/goecs"

type Context any

type LayoutComponentI interface {
	PrepareLayout() bool
	Layout(_lay *Context)
}

type LayoutComponent struct {
	ecs.MetaTag `ecs:"component"`

	Context
	Layout LayoutComponentI `gog:"new: 'layout'"`
}

type DragSourceQuery struct {
	ecs.MetaTag `ecs:"query"`
	ecs.Query

	DragSource *DragSourceComponent
}

type DragSourceComponent struct {
	ecs.MetaTag `ecs:"component"`

	Target int
}

type DragSourceComponent2 struct {
	ecs.MetaTag `ecs:"component"`
	DragSourceComponent
}

type DragSourceComponent3 struct {
	ecs.MetaTag `ecs:"component"`
	DragSourceComponent2
}

type DragSourceEntity struct {
	ecs.MetaTag `ecs:"archetype: { transient }"`
	ecs.Archetype

	DragSource *DragSourceComponent `ecs:"virtual"`
}

type DragSource2Entity struct {
	ecs.MetaTag `ecs:"archetype: { transient }"`
	DragSourceEntity

	DragSource *DragSourceComponent2 `ecs:"virtual"`
}

type DragSource3Entity struct {
	ecs.MetaTag `ecs:"archetype: { transient }"`
	DragSource2Entity

	DragSource *DragSourceComponent3 `ecs:"virtual"`
}
