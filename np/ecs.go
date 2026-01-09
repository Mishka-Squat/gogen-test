package np

import (
	ecs "github.com/igadmg/goecs"
	"github.com/igadmg/gogen-test/game"
	"github.com/igadmg/gogen-test/gfx"
)

type NewSimpleComponent struct {
	ecs.MetaTag `ecs:"component"`

	Foo int     `gog:"new"`
	Bar float32 `gog:"new"`
	Baz float32 `gog:"new"`
}

type NewSimpleNamedComponent struct {
	ecs.MetaTag `ecs:"component"`

	Foo int     `gog:"new: fuu"`
	Bar float32 `gog:"new: brr"`
	Baz float32 `gog:"new: bzz"`
}

type NewEcsRefsComponent struct {
	ecs.MetaTag `ecs:"component"`

	Foo ecs.Ref[gfx.DrawEntity]    `ecs:"" gog:"new"`          // will generate foo_drawCall gfx.DrawCallFn parameter
	Bar ecs.Ref[gfx.DrawEntity]    `ecs:"reference" gog:"new"` // will generate bar ecs.Ref[gfx.DrawEntity] parameter
	Baz ecs.Ref[game.CursorEntity] `ecs:"" gog:"new"`          // will generate baz_xy vector2.Int, baz_world ecs.Ref[WorldEntity] parameters
}

type NewEcsRefsNamedComponent struct {
	ecs.MetaTag `ecs:"component"`

	Foo ecs.Ref[gfx.DrawEntity]    `ecs:"" gog:"new"`                  // will generate foo_drawCall gfx.DrawCallFn parameter
	Bar ecs.Ref[gfx.DrawEntity]    `ecs:"reference" gog:"new"`         // will generate bar ecs.Ref[gfx.DrawEntity] parameter
	Baz ecs.Ref[game.CursorEntity] `ecs:"" gog:"new: [baz_xy, world]"` // will generate baz_xy vector2.Int, world ecs.Ref[WorldEntity] parameters
}

type NewEcsRefsNamedOverlapComponent struct {
	ecs.MetaTag `ecs:"component"`

	World ecs.Ref[game.WorldEntity]  `ecs:"reference" gog:"new"`         // will generate world ecs.Ref[WorldEntity] parameter
	Foo   ecs.Ref[gfx.DrawEntity]    `ecs:"" gog:"new"`                  // will generate foo_drawCall gfx.DrawCallFn parameter
	Bar   ecs.Ref[gfx.DrawEntity]    `ecs:"reference" gog:"new"`         // will generate bar ecs.Ref[gfx.DrawEntity] parameter
	Baz   ecs.Ref[game.CursorEntity] `ecs:"" gog:"new: [baz_xy, world]"` // will generate baz_xy vector2.Int and reuse world ecs.Ref[WorldEntity]
}

type NewInheritedComponentBase1 struct {
	ecs.MetaTag `ecs:"component"`

	Player ecs.Ref[game.PlayerEntity] `ecs:"reference" gog:"new"` // will generate world ecs.Ref[WorldEntity] parameter
	World  ecs.Ref[game.WorldEntity]  `ecs:"reference" gog:"new"` // will generate world ecs.Ref[WorldEntity] parameter
	Foo    ecs.Ref[gfx.DrawEntity]    `ecs:"" gog:"new"`          // will generate foo_drawCall gfx.DrawCallFn parameter
	Bar    ecs.Ref[gfx.DrawEntity]    `ecs:"reference" gog:"new"` // will generate bar ecs.Ref[gfx.DrawEntity] parameter
}

type NewInheritedComponentBase2 struct {
	ecs.MetaTag `ecs:"component"`

	Colony ecs.Ref[game.ColonyEntity] `ecs:"" gog:"new: 'player'"` // will generate foo_drawCall gfx.DrawCallFn parameter
}

type NewInheritedComponent struct {
	ecs.MetaTag `ecs:"component"`

	NewInheritedComponentBase1 `ecs:"new"`
	NewInheritedComponentBase2 `ecs:"new"`
	Baz                        ecs.Ref[game.CursorEntity] `ecs:"" gog:"new: [baz_xy, world]"` // will generate baz_xy vector2.Int and reuse world ecs.Ref[WorldEntity]
}

type NewEcsEntity struct {
	ecs.MetaTag `ecs:"archetype"`
	ecs.Archetype

	Component1 *NewEcsRefsNamedOverlapComponent `gog:"new"`
}
