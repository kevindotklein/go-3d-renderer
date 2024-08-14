package la

import (
	"fmt"
	"math"
)

//https://www.opengl-tutorial.org/beginners-tutorials/tutorial-3-matrices/

type Vector4 struct {
	X, Y, Z, W float32
}

type Matrix4 struct {
	A, B, C, D Vector4
}

func (v *Vector4) Dot(m Matrix4) {
	X := (v.X * m.A.X) + (v.Y * m.A.Y) + (v.Z * m.A.Z) + (v.W * m.A.W)
	Y := (v.X * m.B.X) + (v.Y * m.B.Y) + (v.Z * m.B.Z) + (v.W * m.B.W)
	Z := (v.X * m.C.X) + (v.Y * m.C.Y) + (v.Z * m.C.Z) + (v.W * m.C.W)
	W := (v.X * m.D.X) + (v.Y * m.D.Y) + (v.Z * m.D.Z) + (v.W * m.D.W)

	v.X = X
	v.Y = Y
	v.Z = Z
	v.W = W
}

// https://en.wikipedia.org/wiki/Rotation_matrix#General_rotations
func (v *Vector4) Rotate(yaw, pitch, roll float64) {

	X := math.Cos(yaw) * math.Cos(pitch) * float64(v.X) + (math.Cos(yaw) * math.Sin(pitch) * math.Sin(roll) - math.Sin(yaw) * math.Cos(roll)) * 
			 float64(v.Y) + (math.Cos(yaw) * math.Sin(pitch) * math.Cos(roll) + math.Sin(yaw) * math.Sin(roll)) * float64(v.Z)

	Y := math.Sin(yaw) * math.Cos(pitch) * float64(v.X) + (math.Sin(yaw) * math.Sin(pitch) * math.Sin(roll) + math.Cos(yaw) * math.Cos(roll)) *
		   float64(v.Y) + (math.Sin(yaw) * math.Sin(pitch) * math.Cos(roll) - math.Cos(yaw) * math.Sin(roll)) * float64(v.Z)

	Z := -math.Sin(pitch) * float64(v.X) + math.Cos(pitch) * math.Sin(roll) * float64(v.Y) + math.Cos(pitch) * math.Cos(roll) * float64(v.Z)

	v.X = float32(X)
	v.Y = float32(Y)
	v.Z = float32(Z)
}

func (v Vector4) String() string {
	return fmt.Sprintf("|%.1f|\n|%.1f|\n|%.1f|\n|%.1f|\n", v.X, v.Y, v.Z, v.W)
}

func (m Matrix4) String() string {
	return fmt.Sprintf("|%.1f, %.1f, %.1f, %.1f|\n|%.1f, %.1f, %.1f, %.1f|\n|%.1f, %.1f, %.1f, %.1f|\n|%.1f, %.1f, %.1f, %.1f|\n", 
											m.A.X, m.A.Y, m.A.Z, m.A.W,
											m.B.X, m.B.Y, m.B.Z, m.B.W,
											m.C.X, m.C.Y, m.C.Z, m.C.W,
											m.D.X, m.D.Y, m.D.Z, m.D.W)
}
