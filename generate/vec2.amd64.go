//go:build amd64

//go:generate go run vec2.amd64.go -out ../vec2.amd64.s -stubs ../vec2.amd64.go -pkg qm
package main

import (
	. "github.com/mmcloughlin/avo/build"
	"github.com/mmcloughlin/avo/buildtags"
	"github.com/mmcloughlin/avo/operand"
)

func main() {
	defer Generate()
	Constraint(buildtags.Term("amd64"))

	TEXT("add_array_2_float32", NOSPLIT|NOFRAME, "func(l, r [2]float32) [2]float32")
	Pragma("nosplit")
	Pragma("noescape")
	{
		rx := Load(Param("r").Index(0), XMM())
		ADDSS(operand.NewParamAddr("l", 0), rx)
		Store(rx, ReturnIndex(0).Index(0))

		ry := Load(Param("r").Index(1), XMM())
		ADDSS(operand.NewParamAddr("l", 4), ry)
		Store(ry, ReturnIndex(0).Index(1))

		RET()
	}

	TEXT("sub_array_2_float32", NOSPLIT|NOFRAME, "func(l, r [2]float32) [2]float32")
	Pragma("nosplit")
	{
		rx := Load(Param("r").Index(0), XMM())
		SUBSS(operand.NewParamAddr("l", 0), rx)
		Store(rx, ReturnIndex(0).Index(0))

		ry := Load(Param("r").Index(1), XMM())
		SUBSS(operand.NewParamAddr("l", 4), ry)
		Store(ry, ReturnIndex(0).Index(1))
		RET()
	}

	TEXT("mul_array_2_float32", NOSPLIT|NOFRAME, "func(l, r [2]float32) [2]float32")
	Pragma("nosplit")
	{
		rx := Load(Param("r").Index(0), XMM())
		MULSS(operand.NewParamAddr("l", 0), rx)
		Store(rx, ReturnIndex(0).Index(0))

		ry := Load(Param("r").Index(1), XMM())
		MULSS(operand.NewParamAddr("l", 4), ry)
		Store(ry, ReturnIndex(0).Index(1))

		RET()
	}

	TEXT("div_array_2_float32", NOSPLIT|NOFRAME, "func(l, r [2]float32) [2]float32")
	Pragma("nosplit")
	{
		rx := Load(Param("r").Index(0), XMM())
		DIVSS(operand.NewParamAddr("l", 0), rx)
		Store(rx, ReturnIndex(0).Index(0))

		ry := Load(Param("r").Index(1), XMM())
		DIVSS(operand.NewParamAddr("l", 4), ry)
		Store(ry, ReturnIndex(0).Index(1))

		RET()
	}
}
