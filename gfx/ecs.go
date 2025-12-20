package gfx

import (
	"github.com/igadmg/gamemath/rect2"
	ecs "github.com/igadmg/goecs"
)

type DrawCallFn func(source rect2.Float32)

type DrawComponent struct {
	_ ecs.MetaTag `ecs:"component"`

	DrawCall DrawCallFn
}

type DrawEntity struct {
	_  ecs.MetaTag `ecs:"archetype: { transient }"`
	Id ecs.Id

	Component *DrawComponent `gog:"new: drawCall"`
}
