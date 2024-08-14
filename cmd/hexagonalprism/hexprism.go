package hexagonalprism

import (
	"github.com/kevindotklein/go-3d-renderer/pkg/la"
)

func InitHexagonPrismVertices() []la.Vector4 {
	return []la.Vector4{
		{X:  0.5,  Y:  0.0,   Z:  0.5, W: 1.0},
		{X:  0.25, Y:  0.433, Z:  0.5, W: 1.0},
		{X: -0.25, Y:  0.433, Z:  0.5, W: 1.0},
		{X: -0.5,  Y:  0.0,   Z:  0.5, W: 1.0},
		{X: -0.25, Y: -0.433, Z:  0.5, W: 1.0},
		{X:  0.25, Y: -0.433, Z:  0.5, W: 1.0},
		{X:  0.5,  Y:  0.0,   Z: -0.5, W: 1.0},
		{X:  0.25, Y:  0.433, Z: -0.5, W: 1.0},
		{X: -0.25, Y:  0.433, Z: -0.5, W: 1.0},
		{X: -0.5,  Y:  0.0,   Z: -0.5, W: 1.0},
		{X: -0.25, Y: -0.433, Z: -0.5, W: 1.0},
		{X:  0.25, Y: -0.433, Z: -0.5, W: 1.0},
	}
}
