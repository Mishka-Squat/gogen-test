package game

import (
	"github.com/igadmg/gamemath/vector2"
	ecs "github.com/igadmg/goecs"
	"github.com/igadmg/gogen-test/gfx"
	"github.com/igadmg/gogen-test/input"
	"github.com/igadmg/gogen-test/ui"
)

type WorldModelComponent struct {
	ecs.MetaTag `ecs:"component"`

	Date     string                  `ecs:"a, dto"`
	Colonies []ecs.Ref[ColonyEntity] `ecs:"a, dto"`
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

type ColonyModelComponent struct {
	ecs.MetaTag `ecs:"component"`

	Name   string                `ecs:"a, dto"`
	Player ecs.Ref[PlayerEntity] `ecs:"a, dto, reference" gog:"new"`
}

type ColonyEntity struct {
	ecs.MetaTag `ecs:"archetype"`
	ecs.Archetype

	Model *ColonyModelComponent `gog:"new"`
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

	background       ecs.Ref[gfx.DrawEntity]     `ecs:"a" gog:"new"` // background is transient ref here, because DrawEntity s transient, should be created automaticaly on create transient step
	panelBackground  ecs.Ref[gfx.DrawEntity]     `ecs:"a" gog:"new"`
	loadingIndicator ecs.Ref[gfx.AnimatedSprite] `ecs:"a"`
}

type ScreenModelComponent struct {
	ecs.MetaTag `ecs:"component"`

	world  ecs.Ref[WorldEntity]  `ecs:"a, reference" gog:"new"` // reference components should not be created by default, but also not recreated as transient refs
	Player ecs.Ref[PlayerEntity] `ecs:"a, reference" gog:"new"`
	Cursor ecs.Ref[CursorEntity] `ecs:"a" gog:"new: 'cursor_xy, world'"`
}

type ScreenInputComponent struct {
	ecs.MetaTag `ecs:"component: { transient }"`
	input.InputSchemeComponent

	OnCursorPress func()
}

type ScreenEntity struct {
	ecs.MetaTag `ecs:"archetype"`
	ecs.Archetype

	layout *ScreenLayoutComponent `gog:"new: '@'"`
	View   *ScreenViewComponent   `ecs:"virtual" gog:"new: {
		background: '@.DrawBackground',
		panelBackground: '@.DrawPanelBackground',
	}"`
	Model *ScreenModelComponent `ecs:"virtual" gog:"new: 'world, player, cursor_xy'"`
	Input *ScreenInputComponent `gog:""`
}

func (s ScreenEntity) DrawBackground() {

}

func (s ScreenEntity) DrawPanelBackground() {

}

/*
	gog:"input: {
			click: {
				desktop: {
					KeyInputEntity: {
						key: [ rl.Keyboard_KeyEscape ],
					}
				},
				laptop: desktop,
				default: desktop
			}
		}"
*/
func (s ScreenEntity) InputClick() {

}

type SubScreenViewComponent struct {
	ecs.MetaTag `ecs:"component"`
	ScreenViewComponent

	foreground      ecs.Ref[gfx.DrawEntity] `ecs:"a" gog:"new"`
	panelForeground ecs.Ref[gfx.DrawEntity] `ecs:"a" gog:"new"`
}

type SubScreenEntity struct {
	ecs.MetaTag `ecs:"archetype"`
	ScreenEntity

	View *SubScreenViewComponent `ecs:"virtual" gog:"new: {
		foreground: '@.DrawForeground',
		panelForeground: '@.DrawPanelForeground',
	}"`
}

func (s SubScreenEntity) DrawForeground() {

}

func (s SubScreenEntity) DrawPanelForeground() {

}

type ComplexScreenLayoutComponent struct {
	ecs.MetaTag        `ecs:"component: { transient }"`
	ui.LayoutComponent `gog:"new"`
}

type ComplexScreenViewModelComponent struct {
	ecs.MetaTag `ecs:"component"`

	background ecs.Ref[gfx.DrawEntity] `ecs:"a" gog:"new"`            // background is transient ref here, because DrawEntity s transient, should be created automaticaly on create transient step
	world      ecs.Ref[WorldEntity]    `ecs:"a, reference" gog:"new"` // reference components should not be created by default, but also not recreated as transient refs
	Player     ecs.Ref[PlayerEntity]   `ecs:"a, reference" gog:"new"`
	Cursor     ecs.Ref[CursorEntity]   `ecs:"a" gog:"new: cursor_xy"`
}

type ComplexScreenEntity struct {
	ecs.MetaTag `ecs:"archetype"`
	ecs.Archetype

	layout    *ComplexScreenLayoutComponent    `gog:"new: '@'"`
	ViewModel *ComplexScreenViewModelComponent `gog:"new: {
		background: '@.DrawBackground',
		world, player, cursor_xy,
	}"`
}

func (s ComplexScreenEntity) DrawBackground() {

}

type SystemScreenViewComponent struct {
	ecs.MetaTag `ecs:"component"`

	background      ecs.Ref[gfx.DrawEntity] `ecs:"a"` // background is transient ref here, because DrawEntity s transient, should be created automaticaly on create transient step
	panelBackground ecs.Ref[gfx.DrawEntity] `ecs:"a"`
}

type SystemScreenEntity struct {
	ecs.MetaTag `ecs:"archetype"`
	ecs.Archetype

	View *SystemScreenViewComponent
}

type ColonyScreenModelComponent struct {
	ecs.MetaTag          `ecs:"component"`
	ScreenModelComponent `gog:"new: 'world, colony.Get().PlayerRef(), cursor_xy'"`

	Colony ecs.Ref[ColonyEntity] `ecs:"a, reference" gog:"new"`
}

type ColonyScreenEntity struct {
	ecs.MetaTag `ecs:"archetype"`
	ecs.Archetype

	layout *ScreenLayoutComponent `gog:"new: '@'"`
	View   *ScreenViewComponent   `ecs:"virtual" gog:"new: {
		background: '@.DrawBackground',
		panelBackground: '@.DrawPanelBackground',
	}"`
	Model *ColonyScreenModelComponent `gog:"new"` // TODO: FIX broken if "new: 'world, colony, cursor'"
}

type Colony2ScreenModelComponent struct {
	ecs.MetaTag          `ecs:"component"`
	ScreenModelComponent `gog:"new"`

	Colony ecs.Ref[ColonyEntity] `ecs:"a, reference" gog:"new"`
}

type Colony2ScreenEntity struct {
	ecs.MetaTag  `ecs:"archetype"`
	ScreenEntity `ecs:"virtual" gog:"new: 'world, colony.Get().PlayerRef(), cursor_xy'"`

	Model *Colony2ScreenModelComponent `ecs:"virtual" gog:"new: 'world, colony.Get().PlayerRef(), cursor_xy, colony'"`
}
