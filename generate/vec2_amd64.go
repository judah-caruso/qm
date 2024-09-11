//go:build amd64

//go:generate go run . -out ../vec2_amd64.s -stubs ../vec2_amd64.go -pkg qm
package main

import (
	. "github.com/mmcloughlin/avo/build"
	"github.com/mmcloughlin/avo/buildtags"
)

func main() {
	Constraint(buildtags.Term("amd64"))
	defer Generate()
}
