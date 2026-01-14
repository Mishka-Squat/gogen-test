package dbg

import (
	ecs "github.com/igadmg/goecs"
	"github.com/igadmg/gogen-test/gfx"
	"github.com/igadmg/gogen-test/input"
)

type SystemViewI interface {
	DrawMenu()
	DrawOverlay()
}

type SystemEntity struct {
	ecs.MetaTag `ecs:"archetype"`
	ecs.Archetype

	View  *SystemViewComponent  `ecs:"virtual"`
	Input *SystemInputComponent `ecs:"virtual"`
}

func (e SystemEntity) DrawMenu() {

}

func (e SystemEntity) DrawOverlay() {

}

type SystemViewComponent struct {
	ecs.MetaTag `ecs:"component"`

	SystemViewI `ecs:"self"`

	MenuView    ecs.Ref[gfx.DrawEntity]
	OverlayView ecs.Ref[gfx.DrawEntity]
}

type SystemInputComponent struct {
	ecs.MetaTag `ecs:"component"`
	input.InputSchemeComponent
}

type DbgSystemEntity struct {
	ecs.MetaTag `ecs:"archetype"`
	SystemEntity

	View *DbgSystemViewComponent
}

type DbgSystemViewComponent struct {
	ecs.MetaTag `ecs:"component"`

	SystemViewComponent

	LoadSavFn func(path string) error
	LoadFn    func(path string) error
	SaveFn    func(path string) error
}
