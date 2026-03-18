package dbg

import (
	ecs "github.com/Mishka-Squat/goecs"
	"github.com/Mishka-Squat/gogen-test/gfx"
	"github.com/Mishka-Squat/gogen-test/input"
	"github.com/Mishka-Squat/gogen-test/ui"
)

type TestOverlayViewComponent struct {
	ecs.MetaTag `ecs:"component"`

	layer gfx.LayerId `gog:"new"`
}

type TestOverlayLayout struct {
	ecs.MetaTag        `ecs:"component: { transient }"`
	ui.LayoutComponent `gog:"new"`
}

type OverlayEntity struct {
	ecs.MetaTag    `ecs:"archetype"`
	gfx.DrawEntity `gog:"new: '@.Draw', prepare: { Layer: 'Named(@.view.layer)' }"`

	layout *TestOverlayLayout        `gog:"new: '@'"`
	view   *TestOverlayViewComponent `gog:"new"`
}

func (e OverlayEntity) PrepareLayout() *ui.Context {
	return &e.layout.Context
}

func (e OverlayEntity) Layout(_lay *ui.Context) {
}

func (e OverlayEntity) Draw(layer *gfx.Layer) {
}

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
	ecs.MetaTag `ecs:"component: { transient }"`
	input.InputSchemeComponent
}

type DbgSystemEntity struct {
	ecs.MetaTag `ecs:"archetype"`
	SystemEntity

	View  *DbgSystemViewComponent
	Input *DbgSystemInputComponent
}

type DbgSystemViewComponent struct {
	ecs.MetaTag `ecs:"component"`

	SystemViewComponent

	LoadSavFn func(path string) error
	LoadFn    func(path string) error
	SaveFn    func(path string) error
}

type DbgSystemInputComponent struct {
	ecs.MetaTag `ecs:"component: { transient }"`
	SystemInputComponent
}
