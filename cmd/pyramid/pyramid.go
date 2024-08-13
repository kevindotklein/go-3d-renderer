package pyramid

import "github.com/kevindotklein/go-3d-renderer/pkg/la"

func InitPyramidVertices() []la.Vector4 {
	return []la.Vector4{
		{X: -0.5, Y:  0.5, Z: -0.5, W: 1.0},
		{X:  0.5, Y:  0.5, Z: -0.5, W: 1.0},
		{X:  0.5, Y:  0.5, Z:  0.5, W: 1.0},
		{X: -0.5, Y:  0.5, Z:  0.5, W: 1.0}, 		
		{X:  0.0, Y: -0.5, Z:  0.0, W: 1.0},
		{X: -0.5, Y:  0.5, Z: -0.5, W: 1.0},
		{X:  0.5, Y:  0.5, Z: -0.5, W: 1.0},
		{X:  0.5, Y:  0.5, Z:  0.5, W: 1.0},
		{X: -0.5, Y:  0.5, Z:  0.5, W: 1.0},
	}
}