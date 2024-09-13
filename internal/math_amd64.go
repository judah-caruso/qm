//go:build amd64

//go:generate go run math_amd64.go -out ../math_amd64.s -pkg qm
package main

import (
	. "github.com/mmcloughlin/avo/build"
	"github.com/mmcloughlin/avo/buildtags"
)

func main() {
	defer Generate()
	Constraint(buildtags.Term("amd64"))

	// @todo: needs an update for fixed-point numbers?
	//TEXT("Sqrt", NOSPLIT|NOFRAME, "func(x float32) float32")
	//Pragma("nosplit")
	//Pragma("noescape")
	//{
	//	x := Load(Param("x"), XMM())
	//	SQRTPS(x, x)
	//	Store(x, ReturnIndex(0))
	//	RET()
	//}
	//
	//TEXT("InvSqrt", NOSPLIT|NOFRAME, "func(x float32) float32")
	//Pragma("nosplit")
	//Pragma("noescape")
	//{
	//	x := Load(Param("x"), XMM())
	//	SQRTPS(x, x)
	//	RCPSS(x, x)
	//	Store(x, ReturnIndex(0))
	//	RET()
	//}

	// Until the Go compiler inlines assembly functions,
	// this will be slower in almost all cases.
	/*

		arrayof_2float32_procs := []struct {
			name string
			inst func(operand.Op, operand.Op)
		}{
			{name: "add", inst: ADDSS},
			{name: "sub", inst: SUBSS},
			{name: "mul", inst: MULSS},
			{name: "div", inst: DIVSS},
		}

		for _, proc := range arrayof_2float32_procs {
			TEXT(fmt.Sprintf("%s_arrayof_2float32", proc.name), NOSPLIT|NOFRAME, "func(l, r [2]float32) [2]float32")
			Pragma("nosplit")
			Pragma("noescape")

			rx := Load(Param("r").Index(0), XMM())
			proc.inst(operand.NewParamAddr("l", 0), rx)
			Store(rx, ReturnIndex(0).Index(0))

			ry := Load(Param("r").Index(1), XMM())
			proc.inst(operand.NewParamAddr("l", 4), ry)
			Store(ry, ReturnIndex(0).Index(1))

			RET()
		}
	*/
}
