package game

import (
	"github.com/Mishka-Squat/gamemath/vector2"
	ecs "github.com/Mishka-Squat/goecs"
	"github.com/Mishka-Squat/gogen-test/gfx"
	"github.com/Mishka-Squat/gogen-test/input"
	"github.com/Mishka-Squat/gogen-test/ui"
	rl "github.com/Mishka-Squat/raylib-go/raylib"
)

type SaveTag struct {
	ecs.MetaTag `ecs:"component"`
}

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

type TestEmptyModelComponent struct {
	ecs.MetaTag `ecs:"component"`
}

type TestEmptyModelEntity struct {
	ecs.MetaTag `ecs:"archetype"`
	ecs.Archetype

	Model *TestEmptyModelComponent `ecs:"save, virtual"`
}

type TestNationModelComponent struct {
	ecs.MetaTag `ecs:"component"`

	Id   int    `ecs:"a: NationId, dto"`
	Name string `ecs:"a, dto"`
}

type TestNationEntity struct {
	ecs.MetaTag `ecs:"archetype"`
	ecs.Archetype
	SaveTag

	Model *TestNationModelComponent `ecs:"save, virtual"`
}

type TestPlayerNationModelComponent struct {
	ecs.MetaTag `ecs:"component"`

	TestNationModelComponent
	TaxRate uint8  `ecs:"a, dto"`
	Gold    uint32 `ecs:"a, dto"`
}

type TestPlayerNationEntity struct {
	ecs.MetaTag `ecs:"archetype"`
	TestNationEntity

	Model *TestPlayerNationModelComponent `ecs:"save"`
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

type ControlLayoutComponent struct {
	ecs.MetaTag        `ecs:"component: { transient }"`
	ui.LayoutComponent `gog:"new"`
}

type ControlViewComponent struct {
	ecs.MetaTag `ecs:"component"`

	background      ecs.Ref[gfx.DrawEntity] `ecs:"a" gog:"new, prepare: 'Named(LayerBackground)'"` // background is transient ref here, because DrawEntity s transient, should be created automaticaly on create transient step
	panelBackground ecs.Ref[gfx.DrawEntity] `ecs:"a" gog:"new"`
}

type ControlModelComponent struct {
	ecs.MetaTag `ecs:"component"`

	world  ecs.Ref[WorldEntity]  `ecs:"a, reference" gog:"new"` // reference components should not be created by default, but also not recreated as transient refs
	Player ecs.Ref[PlayerEntity] `ecs:"a, reference" gog:"new"`
	Cursor ecs.Ref[CursorEntity] `ecs:"a" gog:"new: 'cursor_xy, world'"`
}

type ControlInputComponent struct {
	ecs.MetaTag `ecs:"component: { transient }"`
	input.InputSchemeComponent

	OnCursorPress func()
}

type ControlEntity struct {
	ecs.MetaTag `ecs:"archetype"`
	ecs.Archetype

	layout *ControlLayoutComponent `gog:"new: '@'"`
	View   *ControlViewComponent   `ecs:"virtual" gog:"new: {
		background: '@.DrawBackground',
		panelBackground: '@.DrawPanelBackground',
	}"`
	Model *ControlModelComponent `ecs:"virtual" gog:"new: 'world, player, cursor_xy'"`
	Input *ControlInputComponent `gog:""`
}

func (s ControlEntity) InputClickKey() rl.UnifiedKeyType {
	return rl.Keyboard_KeyEscape
}

/*
	gog:"input: {
			click: {
				desktop: {
					KeyInputEntity: {
						keyfn: 's.InputClickKey()',
					}
				},
				laptop: desktop,
				default: desktop
			}
		}"
*/
func (s ControlEntity) InputClick() {

}

type ScreenLayoutComponent struct {
	ecs.MetaTag        `ecs:"component: { transient }"`
	ui.LayoutComponent `gog:"new"`
}

type ScreenViewComponent struct {
	ecs.MetaTag `ecs:"component"`

	background       ecs.Ref[gfx.DrawEntity]     `ecs:"a" gog:"new"` // background is transient ref here, because DrawEntity s transient, should be created automaticaly on create transient step
	panelBackground  ecs.Ref[gfx.DrawEntity]     `ecs:"a" gog:"new"`
	loadingIndicator ecs.Ref[gfx.AnimatedSprite] `ecs:"a" gog:"new"`
	control          ecs.Ref[ControlEntity]      ``
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
		loadingIndicator: 'bound',
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
	ecs.MetaTag         `ecs:"component"`
	ScreenViewComponent `gog:"new"`

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
	Cursor     ecs.Ref[CursorEntity]   `ecs:"a" gog:"new: 'cursor_xy, world'"`
}

type ComplexScreenEntity struct {
	ecs.MetaTag `ecs:"archetype"`
	ecs.Archetype

	layout    *ComplexScreenLayoutComponent    `gog:"new: '@'"`
	ViewModel *ComplexScreenViewModelComponent `gog:"new: {
		background: '@.DrawBackground',
	}"`
}

func (s ComplexScreenEntity) DrawBackground() {

}

type SystemControlEntity struct {
	ecs.MetaTag `ecs:"archetype"`
	ecs.Archetype

	layout *ControlLayoutComponent `gog:"new: '@'"`
	View   *ControlViewComponent   `ecs:"virtual" gog:"new: {
		background: '@.DrawBackground',
		panelBackground: '@.DrawPanelBackground',
	}"`
	Input *ControlInputComponent `gog:""`
}

type SystemScreenViewComponent struct {
	ecs.MetaTag `ecs:"component"`

	background      ecs.Ref[gfx.DrawEntity]      `ecs:"a" gog:"new"` // background is transient ref here, because DrawEntity s transient, should be created automaticaly on create transient step
	panelBackground ecs.Ref[gfx.DrawEntity]      `ecs:"a"`
	control         ecs.Ref[SystemControlEntity] `ecs:"a" gog:"new"`
}

type SystemScreenEntity struct {
	ecs.MetaTag `ecs:"archetype"`
	ecs.Archetype

	View *SystemScreenViewComponent `ecs:"virtual" gog:"new: {
		background: '@.DrawBackground',
	}"`
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
		loadingIndicator: 'bound',
	}"`
	Model *ColonyScreenModelComponent `gog:"new"` // TODO(iga) FIX: broken if "new: 'world, colony, cursor'"
}

type Colony2ScreenModelComponent struct {
	ecs.MetaTag          `ecs:"component"`
	ScreenModelComponent `gog:"new"`

	Colony ecs.Ref[ColonyEntity] `ecs:"a, reference" gog:"new"`
}

type Colony2ScreenEntity struct {
	ecs.MetaTag  `ecs:"archetype"`
	ScreenEntity `gog:"new: 'bound, world, colony.Get().PlayerRef(), cursor_xy'"`

	Model *Colony2ScreenModelComponent `ecs:"virtual" gog:"new: 'world, colony.Get().PlayerRef(), cursor_xy, colony'"`
}

type TestMapScreen struct {
	ecs.MetaTag `ecs:"archetype"`
	ecs.Archetype

	layout *TestMapScreenLayout `gog:"new: '@'"`
	Input  *TestMapScreenInput  `gog:""`
	Model  *TestMapScreenModel  `gog:"new: 'world, player'"`
	View   *TestMapScreenView   `gog:"new: {
		background: '@.DrawBackground',
		world, player, '@.Cursor()'
	}"`
}

type TestMapScreenLayout struct {
	ecs.MetaTag        `ecs:"component: { transient }"`
	ui.LayoutComponent `gog:"new"`
}

type TestMapScreenInput struct {
	ecs.MetaTag `ecs:"component: { transient }"`
	input.InputSchemeComponent
}

type TestMapScreenModel struct {
	ecs.MetaTag `ecs:"component"`

	world  ecs.Ref[WorldEntity]  `ecs:"a, reference" gog:"new"`
	player ecs.Ref[PlayerEntity] `ecs:"a, reference" gog:"new"`
	Cursor ecs.Ref[CursorEntity] `ecs:"a" gog:"new: [cursor_xy, world]"`
}

type TestMapScreenView struct {
	ecs.MetaTag `ecs:"component"`

	background ecs.Ref[gfx.DrawEntity]            `ecs:"a" gog:"new"`
	mapView    ecs.Ref[TestMapViewEntity]         `ecs:"a" gog:"new: 'world, player, cursor'"`
	infoView   ecs.Ref[BaseTestMapInfoViewEntity] `ecs:"a" gog:"new: 'world, player, cursor'"`
}

type TestMapViewEntity struct {
	ecs.MetaTag `ecs:"archetype"`

	BaseTestMapViewEntity `gog:"new"`

	Model *TestMapModel `ecs:"virtual" gog:"new: 'world, player, cursor'"`
	View  *TestMapView  `ecs:"virtual" gog:"new: {
		layerFog: '@.DrawFogTiles',
		layerUi: '@.DrawUiTiles',
		layerOverlay: '@.DrawOverlayTiles',
	}"`
	Input *TestMapViewEntityInput `gog:""`
}

type TestMapView struct {
	ecs.MetaTag `ecs:"component"`

	BaseTestMapView
	//layerUnits   ecs.Ref[gfx.DrawCallEntity]       `gog:"new"`
	layerFog     ecs.Ref[gfx.DrawEntity] `gog:"new"`
	layerUi      ecs.Ref[gfx.DrawEntity] `gog:"new"`
	layerOverlay ecs.Ref[gfx.DrawEntity] `gog:"new"`
}

type TestMapModel struct {
	ecs.MetaTag `ecs:"component"`

	BaseTestMapModel `gog:"new: 'world, player'"`
	Cursor           ecs.Ref[CursorEntity] `ecs:"a, reference" gog:"new"`
}

type TestMapViewEntityInput struct {
	ecs.MetaTag `ecs:"component: { transient }"`
	input.InputSchemeComponent
}

type BaseTestMapViewEntity struct {
	ecs.MetaTag `ecs:"archetype"`
	ecs.Archetype

	Model *BaseTestMapModel `ecs:"abstract"`
	View  *BaseTestMapView  `ecs:"abstract" gog:"new: {
		layerTerrain: '@.DrawLayerTerrain',
		layerGame: '@.DrawLayerGame',
		layerDebug: '@.DrawLayerDebug'
	}"`
}

type BaseTestMapView struct {
	ecs.MetaTag `ecs:"component"`

	layerTerrain ecs.Ref[gfx.DrawEntity] `gog:"new"`
	layerGame    ecs.Ref[gfx.DrawEntity] `gog:"new"`
	layerDebug   ecs.Ref[gfx.DrawEntity] `gog:"new"`
}

type BaseTestMapModel struct {
	ecs.MetaTag `ecs:"component"`

	World  ecs.Ref[WorldEntity]  `ecs:"a, reference" gog:"new"`
	Player ecs.Ref[PlayerEntity] `ecs:"a, reference" gog:"new"`
}

type BaseTestMapInfoView struct {
	ecs.MetaTag `ecs:"component"`

	World  ecs.Ref[WorldEntity]  `ecs:"a, reference" gog:"new"`
	Player ecs.Ref[PlayerEntity] `ecs:"a, reference" gog:"new"`
	Cursor ecs.Ref[CursorEntity] `ecs:"a, reference" gog:"new"`

	//input InputScheme `gog:""`
	//click *BoundedInputComponent
}

type BaseTestMapInfoViewEntity struct {
	ecs.MetaTag `ecs:"archetype"`
	//gfx.BoundDrawEntity `gog:"new: '@.DrawPanel', prepare: { Layer: 'Named(LayerUI)' }"`
	gfx.DrawEntity `gog:"new: '@.DrawPanel', prepare: { Layer: 'Named(LayerUI)' }"`

	View *BaseTestMapInfoView `gog:"new: 'world, player, cursor'"`
}

type TestMessageBoxModel struct {
	ecs.MetaTag `ecs:"component"`

	resultFn func(int) `ecs:"a" gog:"new"`
}

type TestMessageBoxViewI interface {
}

type TestMessageBoxView struct {
	ecs.MetaTag `ecs:"component"`

	ui.Context
	Content TestMessageBoxViewI     `ecs:"a" gog:"new"`
	panel   ecs.Ref[gfx.DrawEntity] `ecs:"a" gog:"new"`
}

type TestMessageBox struct {
	ecs.MetaTag `ecs:"archetype"`
	ecs.Archetype

	Model *TestMessageBoxModel `ecs:"virtual" gog:"new"`
	View  *TestMessageBoxView  `ecs:"abstract" gog:"new: {
		panel: '@.Draw',
	}"`
	Input *TestMessageBoxInput `ecs:"virtual" gog:""`
}

type TestMessageBoxInput struct {
	ecs.MetaTag `ecs:"component: { transient }"`
	input.InputSchemeComponent
}

/*
	gog:"input: {
			confirm: {
				desktop: {
					KeyInputEntity: {
						key: [ rl.Keyboard_KeyEnter ],
					}
				},
				laptop: desktop,
				default: desktop
			}
		}"
*/
func (s TestMessageBox) InputConfirm() {
}

type TestMessageBoxWithChoiceModel struct {
	ecs.MetaTag `ecs:"component"`

	TestMessageBoxModel
}

type TestMessageBoxWithChoiceView struct {
	ecs.MetaTag `ecs:"component"`

	TestMessageBoxView
	menu ecs.Ref[PlayerEntity] `ecs:"a" gog:"new"`
}

type TestMessageBoxWithChoice struct {
	ecs.MetaTag `ecs:"archetype"`
	TestMessageBox

	Model *TestMessageBoxWithChoiceModel `ecs:"virtual" gog:"new"`
	View  *TestMessageBoxWithChoiceView  `ecs:"virtual" gog:"new: '@'"`
	Input *TestMessageBoxWithChoiceInput `ecs:"virtual" gog:""`
}

type TestMessageBoxWithChoiceInput struct {
	ecs.MetaTag `ecs:"component: { transient }"`
	TestMessageBoxInput
}
