package gfx

import (
	"github.com/igadmg/gamemath/rect2"
	ecs "github.com/igadmg/goecs"
)

type DrawCallFn func(source rect2.Float32)

type DrawComponent struct {
	ecs.MetaTag `ecs:"component"`

	DrawCall DrawCallFn `gog:"new"`
}

type Layer struct {
	buffer int
	output int
}

type LayerComponent struct {
	ecs.MetaTag `ecs:"component"`

	*Layer
}

func (l *LayerComponent) Prepare() {
	l.Layer = nil
}

type DrawEntity struct {
	ecs.MetaTag `ecs:"archetype: { transient }"`
	ecs.Archetype

	Draw  *DrawComponent `gog:"new: drawCall"`
	Layer *LayerComponent
}

type DoubleDrawEntity struct {
	ecs.MetaTag `ecs:"archetype: { transient }"`
	DrawEntity  `gog:"new"`

	DrawAgain *DrawComponent `gog:"new: drawCall"`
}

type AnimatedSprite struct {
	ecs.MetaTag `ecs:"archetype"`
	DrawEntity  `gog:"new: '@.Draw', prepare: 'Named(LayerUI)'"`
	Bound       BoundComponent `gog:"new"`
}

func (s AnimatedSprite) Draw() {

}

type AnimatedSpriteStrange struct {
	ecs.MetaTag `ecs:"archetype"`
	DrawEntity  `gog:"new: '@.Draw', prepare: 'Named(LayerUI)'"`
	DrawCall    DrawComponent  `gog:"new"`
	Bound       BoundComponent `gog:"new"`
}

func (s AnimatedSpriteStrange) Draw() {

}

type BoundComponent struct {
	ecs.MetaTag `ecs:"component"`

	Bound rect2.Float32 `gog:"new"`
}

type BoundDrawEntity struct {
	ecs.MetaTag `ecs:"archetype: { transient }"`
	DrawEntity  `gog:"new"`

	Bound BoundComponent `gog:"new"`
}

// Here we removed base class new parameters and expect them to be imported automaticaly
type BoundDrawEntitySimple struct {
	ecs.MetaTag `ecs:"archetype: { transient }"`
	DrawEntity  `gog:"new"`

	Bound BoundComponent `gog:"new: bound"`
}
