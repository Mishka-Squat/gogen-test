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

type SubSimpleType struct {
	SimpleType

	SubBar SimpleType
}

type SubComplexType struct {
	ComplexType

	SubFoo SimpleType
}

type UltimateType struct {
	SubSimpleType
	SubComplexType

	UltimateBaz float32
}
