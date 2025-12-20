package game

import (
	ecs "github.com/igadmg/goecs"
	"github.com/igadmg/gogen-test/gfx"
)

type WorldModelComponent struct {
	_ ecs.MetaTag `ecs:"component"`

	Date string `ecs:"a, dto"`
}

type WorldEntity struct {
	_  ecs.MetaTag `ecs:"archetype"`
	Id ecs.Id

	Model *WorldModelComponent
}

type ScreenViewComponent struct {
	_ ecs.MetaTag `ecs:"component"`

	background ecs.Ref[gfx.DrawEntity] `ecs:"a"` // background is transient ref here, because DrawEntity s transient, should be created automaticaly on create transient step
}

type ScreenModelComponent struct {
	_ ecs.MetaTag `ecs:"component"`

	worldMap ecs.Ref[WorldEntity] `ecs:"a, reference"` // reference components should not be created by default, but also not recreated as transient refs
}

type ScreenEntity struct {
	_  ecs.MetaTag `ecs:"archetype"`
	Id ecs.Id

	View  *ScreenViewComponent
	Model *ScreenModelComponent
}

func (s ScreenEntity) DrawBackground() {

}
