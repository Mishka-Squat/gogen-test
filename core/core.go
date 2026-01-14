package core

type ProfessionId uint8

type SimpleType struct {
	Foo  int
	Bar  float32
	Baz  []float32
	Prof ProfessionId
}

type ComplexType struct {
	Foo SimpleType
	Bar []SimpleType
}
