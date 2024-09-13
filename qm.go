package qm

// VecElementIndex represents the index of an element for Vector types.
// This allows array access similar to a struct member. (i.e. v[X] vs. v.X)
type VecElementIndex uint8

const (
	X = VecElementIndex(0) // X, 0th element for Vec2, Vec3, Vec4
	Y = VecElementIndex(1) // Y, 1st element for Vec2, Vec3, Vec4
	Z = VecElementIndex(2) // Z, 2nd element for Vec3, Vec4
	W = VecElementIndex(3) // W, 3rd element for Vec4

	// Used to access Vector types as colors.
	R = VecElementIndex(0)
	G = VecElementIndex(1)
	B = VecElementIndex(2)
	A = VecElementIndex(3)

	// Used to access Vector types as UV coordinates.
	// Note: W isn't defined as that would incorrectly overlap with Vec3's Z element.
	U = VecElementIndex(0)
	V = VecElementIndex(1)

	// Used to access Vector types as bounds.
	Width  = VecElementIndex(0)
	Height = VecElementIndex(1)
)
