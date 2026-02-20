package input

import (
	"github.com/Mishka-Squat/gamemath/rect2"
	ecs "github.com/Mishka-Squat/goecs"
	rl "github.com/Mishka-Squat/raylib-go/raylib"
)

type InputScheme interface {
	ecs.IsDeferable

	Init()
	Prepare()
	Layout(rect rect2.Float32)
	Rebuild()
}

type InputSchemeQuery struct {
	ecs.MetaTag `ecs:"query"`
	ecs.Query

	InputScheme *InputSchemeComponent
}

type InputSchemeComponent struct {
	ecs.MetaTag `ecs:"component: { input }"`

	input InputScheme `gog:""`
}

func (o *InputSchemeComponent) Defer() {
	if o.input != nil {
		o.input.Defer()
		o.input = nil
	}
}

type KeySet []rl.KeyType
type KeyChord [][]rl.KeyType // TODO(iga): optimize that to C structure

type OnPressFn = func(mask int)

type KeyInputComponent struct {
	ecs.MetaTag `ecs:"component"`

	Key       KeySet
	Chord     KeyChord
	Delay     float32
	Frequency float32
	next_time float64

	OnPress OnPressFn
}

type KeyInputEntity struct {
	ecs.MetaTag `
		ecs:"archetype: { transient }"
		gog:"input: {
			Input: {
				Key: [
					'input.KeySet${key}',
					'${keyfn}',
				],
				Chord: 'input.KeyChord${chord}',
				Delay: '${delay}',
				Frequency: '${frequency}',
				'func()': {
					OnPress: 'func(_ int) { i_func() }'
				},
				'func(mask int)': {
					OnPress: 'func(mask int) { i_func(mask) }'
				}
			}
		}"`
	ecs.Archetype

	Input *KeyInputComponent
}
