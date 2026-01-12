package core

type SimpleType struct {
	Foo int
	Bar float32
	Baz []float32
}

type ComplexType struct {
	Foo SimpleType
	Bar []SimpleType
}
