package uix

import ecs "github.com/igadmg/goecs"

type Component1_1_1 struct {
	ecs.MetaTag `ecs:"component"`

	a int `ecs:"a"`
	b int `ecs:"a: 'Bee'"`
}

type Archetype1_1_1 struct {
	ecs.MetaTag `ecs:"archetype"`
	ecs.Archetype

	Component1 Component1_1_1
}
