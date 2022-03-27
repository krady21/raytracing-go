package utils

import (
	"fmt"
	"math"
)

type Vec3 struct {
	I, J, K float64
}

func (v *Vec3) Plus(w Vec3) {
	v.I += w.I
	v.J += w.J
	v.K += w.K
}

func (v *Vec3) Times(t float64) {
	v.I *= t
	v.J *= t
	v.K *= t
}

func (v *Vec3) DividedBy(t float64) {
	v.Times(1 / t)
}

func (v Vec3) LenSquared() float64 {
	return v.I*v.I + v.J*v.J + v.K*v.K
}

func (v Vec3) Len() float64 {
	return math.Sqrt(v.LenSquared())
}

func (v Vec3) Neg() Vec3 {
	return Vec3{-v.I, -v.J, -v.K}
}

func (v *Vec3) String() string {
	return fmt.Sprintf("%g %g %g", v.I, v.J, v.K)
}

func Plus(v, w Vec3) Vec3 {
	return Vec3{v.I + w.I, v.J + w.J, v.K + w.K}
}

func Minus(v, w Vec3) Vec3 {
	return Vec3{v.I - w.I, v.J - w.J, v.K - w.K}
}

func Multiply(v, w Vec3) Vec3 {
	return Vec3{v.I * w.I, v.J * w.J, v.K * w.K}
}

func Dot(v, w Vec3) float64 {
	return v.I * w.I + v.J * w.J + v.K * w.K
}

func Cross(v, w Vec3) Vec3 {
	return Vec3 {
		v.J * w.K - v.K * w.J, 
		v.K * w.I - v.I * w.K,
		v.I * w.J - v.J * w.I,
	}
}

type Point3 = Vec3 // 3d point
type Color = Vec3  // rgb color
