package uix

import ecs "github.com/igadmg/goecs"

type Component1_1_1 struct {
	_ ecs.MetaTag `ecs:"component"`

	a int `ecs:"a"`
	b int `ecs:"a: 'Bee'"`
}

type Archetype1_1_1 struct {
	_  ecs.MetaTag `ecs:"archetype"`
	Id ecs.Id

	Component1 Component1_1_1
}
