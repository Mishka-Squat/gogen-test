package gfx

import ecs "github.com/igadmg/goecs"

type DrawComponent struct {
	_ ecs.MetaTag `ecs:"component"`
}

type DrawEntity struct {
	_  ecs.MetaTag `ecs:"archetype"`
	Id ecs.Id

	Component *DrawComponent
}
