package lazy

type StructWithLazy struct {
	//lazy_data_StructWithLazy
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
