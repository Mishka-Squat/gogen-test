package lazy

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

func (s *StructWithLazy) lazy_PtrStringList() []string {
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

func (s *StructWithLazy) lazy_PtrStringMap() map[string]string {
	return map[string]string{
		"a": "Hello",
		"b": "World",
		"c": "!",
	}
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
