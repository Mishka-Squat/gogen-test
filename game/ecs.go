package game

import (
	"github.com/igadmg/gamemath/vector2"
	ecs "github.com/igadmg/goecs"
	"github.com/igadmg/gogen-test/gfx"
	"github.com/igadmg/gogen-test/ui"
)

type WorldModelComponent struct {
	ecs.MetaTag `ecs:"component"`

	Date string `ecs:"a, dto"`
}

type WorldEntity struct {
	ecs.MetaTag `ecs:"archetype"`
	ecs.Archetype

	Model *WorldModelComponent
}

type PlayerEntity struct {
	ecs.MetaTag `ecs:"archetype"`
	ecs.Archetype
}

type CursorComponent struct {
	ecs.MetaTag `ecs:"component"`

	xy    vector2.Int          `ecs:"a, dto" gog:"new"`
	world ecs.Ref[WorldEntity] `ecs:"a, reference" gog:"new"`
}

type CursorEntity struct {
	ecs.MetaTag `ecs:"archetype"`
	ecs.Archetype

	Cursor *CursorComponent `gog:"new"`
}

type ScreenLayoutComponent struct {
	ecs.MetaTag        `ecs:"component: { transient }"`
	ui.LayoutComponent `gog:"new"`
}

type ScreenViewComponent struct {
	ecs.MetaTag `ecs:"component"`

	background      ecs.Ref[gfx.DrawEntity] `ecs:"a"` // background is transient ref here, because DrawEntity s transient, should be created automaticaly on create transient step
	panelBackground ecs.Ref[gfx.DrawEntity] `ecs:"a"`
}

type ScreenModelComponent struct {
	ecs.MetaTag `ecs:"component"`

	world  ecs.Ref[WorldEntity]  `ecs:"a, reference" gog:"new"` // reference components should not be created by default, but also not recreated as transient refs
	Player ecs.Ref[PlayerEntity] `ecs:"a, reference" gog:"new"`
	Cursor ecs.Ref[CursorEntity] `ecs:"a" gog:"new"`
}

type ScreenEntity struct {
	ecs.MetaTag `ecs:"archetype"`
	ecs.Archetype

	layout *ScreenLayoutComponent `gog:"new: '@'"`
	View   *ScreenViewComponent   `gog:"new: { background: '@.DrawBackground' }"`
	Model  *ScreenModelComponent  `gog:"new: 'world, player, cursor'"`
}

func (s ScreenEntity) DrawBackground() {

}

type ComplexScreenLayoutComponent struct {
	ecs.MetaTag        `ecs:"component: { transient }"`
	ui.LayoutComponent `gog:"new"`
}

type ComplexScreenViewModelComponent struct {
	ecs.MetaTag `ecs:"component"`

	background ecs.Ref[gfx.DrawEntity] `ecs:"a"`                      // background is transient ref here, because DrawEntity s transient, should be created automaticaly on create transient step
	world      ecs.Ref[WorldEntity]    `ecs:"a, reference" gog:"new"` // reference components should not be created by default, but also not recreated as transient refs
	Player     ecs.Ref[PlayerEntity]   `ecs:"a, reference" gog:"new"`
	Cursor     ecs.Ref[CursorEntity]   `ecs:"a" gog:"new"`
}

type ComplexScreenEntity struct {
	ecs.MetaTag `ecs:"archetype"`
	ecs.Archetype

	layout    *ComplexScreenLayoutComponent    `gog:"new: '@'"`
	ViewModel *ComplexScreenViewModelComponent `gog:"new: { background: '@.DrawBackground', world, player, cursor }"`
}

func (s ComplexScreenEntity) DrawBackground() {

}
