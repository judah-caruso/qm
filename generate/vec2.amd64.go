//go:build amd64

//go:generate go run vec2.amd64.go -out ../vec2.amd64.s -stubs ../vec2.amd64.go -pkg qm
package main

import (
	"fmt"
	. "github.com/mmcloughlin/avo/build"
	"github.com/mmcloughlin/avo/buildtags"
	"github.com/mmcloughlin/avo/operand"
)

func main() {
	defer Generate()
	Constraint(buildtags.Term("amd64"))

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
}
