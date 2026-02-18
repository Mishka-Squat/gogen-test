package lazy

import "github.com/igadmg/orderedmap/v4"

type StructWithLazy struct {
	//lazy_data_StructWithLazy
}

func (s StructWithLazy) lazy_String() string {
	return "Hello World!"
}

func (s *StructWithLazy) lazy_PtrString() string {
	return "Hello World!"
}

func (s StructWithLazy) lazy_StringList() []string {
	return []string{
		"Hello",
		"World",
		"!",
	}
}

func (s StructWithLazy) lazy_StringMap() map[string]string {
	return map[string]string{
		"a": "Hello",
		"b": "World",
		"c": "!",
	}
}

func (s *StructWithLazy) lazy_StringOrderedMap() orderedmap.Of[*int, []*float64] {
	return orderedmap.Make[*int, []*float64]()
}

// lazy:"set"
func (s StructWithLazy) lazy_StringWithSet() string {
	return "Hello World!"
}

// lazy:"override"
func (s StructWithLazy) lazy_StringWithOverride() string {
	return "Hello World!"
}

// lazy:"set, override"
func (s StructWithLazy) lazy_StringWithSetAndOverride() string {
	return "Hello World!"
}

func (s StructWithLazy) lazy_StringWithError() (string, error) {
	return "Hello World!", nil
}
