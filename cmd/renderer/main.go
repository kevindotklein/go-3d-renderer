package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"os"

	"github.com/kevindotklein/go-3d-renderer/cmd/cube"
	"github.com/kevindotklein/go-3d-renderer/cmd/pyramid"
	"github.com/kevindotklein/go-3d-renderer/cmd/ui"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/kevindotklein/go-3d-renderer/pkg/la"
	"golang.org/x/image/font/basicfont"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
	fov          = 45
	xSpeed       = 0.02
	ySpeed       = 0.02
	zSpeed       = 0.08
)

type model uint8

const (
	cubeModel model = iota
	pyramidModel
)

var distance float32
var inputLabel string
var toggleInfo bool
var currentModel model
var rotationRadians float64

type Game struct {
	vertices     []la.Vector4
	keyStates    map[ebiten.Key]int
}

func point3Dto2D(point, z float32) float32 {
	radians := (float64(fov) / float64(180)) * math.Pi
	return point / (z * float32(math.Tan(radians/2)))
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		for i := range g.vertices {
			g.vertices[i].Rotate(0, 0, -rotationRadians)
		}
		inputLabel = ui.ArrowUpLabel
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		for i := range g.vertices {
			g.vertices[i].Rotate(0, 0, rotationRadians)
		}
		inputLabel = ui.ArrowDownLabel
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		for i := range g.vertices {
			g.vertices[i].Rotate(0, rotationRadians, 0)
		}
		inputLabel = ui.ArrowLeftLabel
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		for i := range g.vertices {
			g.vertices[i].Rotate(0, -rotationRadians, 0)
		}
		inputLabel = ui.ArrowRightLabel
	}

	if ebiten.IsKeyPressed(ebiten.KeyJ) {
		for i := range g.vertices {
			g.vertices[i].Rotate(-rotationRadians, 0, 0)
		}
		inputLabel = ui.JLabel
	} else if ebiten.IsKeyPressed(ebiten.KeyK) {
		for i := range g.vertices {
			g.vertices[i].Rotate(rotationRadians, 0, 0)
		}
		inputLabel = ui.KLabel
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		distance -= zSpeed
		inputLabel = ui.WLabel
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		distance += zSpeed
		inputLabel = ui.SLabel
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		m := la.Matrix4{A: la.Vector4{X: 1.0}, B: la.Vector4{Y: 1.0, W: ySpeed}, C: la.Vector4{Z: 1.0}, D: la.Vector4{W: 1.0}}
		for i := range g.vertices {
			g.vertices[i].Dot(m)
		}
		inputLabel = ui.SpaceLabel
	} else if ebiten.IsKeyPressed(ebiten.KeyShiftLeft) {
		m := la.Matrix4{A: la.Vector4{X: 1.0}, B: la.Vector4{Y: 1.0, W: -ySpeed}, C: la.Vector4{Z: 1.0}, D: la.Vector4{W: 1.0}}
		for i := range g.vertices {
			g.vertices[i].Dot(m)
		}
		inputLabel = ui.LshiftLabel
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		m := la.Matrix4{A: la.Vector4{X: 1.0, W: xSpeed}, B: la.Vector4{Y: 1.0}, C: la.Vector4{Z: 1.0}, D: la.Vector4{W: 1.0}}
		for i := range g.vertices {
			g.vertices[i].Dot(m)
		}
		inputLabel = ui.ALabel
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		m := la.Matrix4{A: la.Vector4{X: 1.0, W: -xSpeed}, B: la.Vector4{Y: 1.0}, C: la.Vector4{Z: 1.0}, D: la.Vector4{W: 1.0}}
		for i := range g.vertices {
			g.vertices[i].Dot(m)
		}
		inputLabel = ui.DLabel
	}

	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if inpututil.IsKeyJustPressed(k) {
			g.keyStates[k] = 1
		} else if ebiten.IsKeyPressed(k) {
			g.keyStates[k]++
		} else {
			g.keyStates[k] = 0
		}

		if k == ebiten.KeyI && g.keyStates[k] == 1 {
			toggleInfo = !toggleInfo
			inputLabel = ui.ILabel
		}

		if k == ebiten.KeyP && g.keyStates[k] == 1 {
			g.vertices = pyramid.InitPyramidVertices()
			currentModel = pyramidModel
			inputLabel = ui.PLabel
		} else if k == ebiten.KeyC && g.keyStates[k] == 1 {
			g.vertices = cube.InitCubeVertices()
			currentModel = cubeModel
			inputLabel = ui.CLabel
		}

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch currentModel {
		case cubeModel:
			renderCube(screen, g)
		case pyramidModel:
			renderPyramid(screen, g)
		default:
			renderCube(screen, g)
	}

	text.Draw(screen, ui.UiMessage, basicfont.Face7x13, 20, 30, color.White)
	text.Draw(screen, inputLabel, basicfont.Face7x13, screenWidth-180, screenHeight-100, color.White)
	if toggleInfo {
		text.Draw(screen, fmt.Sprintf("distance: %.1f", distance), basicfont.Face7x13, screenWidth-130, 30, color.White)
	}
}

func renderCube(screen *ebiten.Image, g *Game) {
	for i := 0; i < 4; i++ {
		drawLine(screen, g.vertices[i], g.vertices[(i+1)%4], color.RGBA{255, 0, 0, 255})
	}

	for i := 4; i < 8; i++ {
		drawLine(screen, g.vertices[i], g.vertices[4+(i-4+1)%4], color.RGBA{255, 0, 0, 255})
	}

	// vertical lines
	for i := 0; i < 4; i++ {
		drawLine(screen, g.vertices[i], g.vertices[i+4], color.RGBA{255, 0, 0, 255})
	}

	for _, v := range g.vertices {
		drawVertex(screen, v, color.RGBA{0, 255, 0, 255}, toggleInfo)
	}
}

func renderPyramid(screen *ebiten.Image, g *Game) {
	for i := 0; i < 4; i++ {
		drawLine(screen, g.vertices[i], g.vertices[(i+1)%4], color.RGBA{255, 0, 0, 255})
	}

	for i := 0; i < 4; i++ {
		drawLine(screen, g.vertices[i], g.vertices[4], color.RGBA{255, 0, 0, 255})
	}

	// vertical lines
	for _, v := range g.vertices {
		drawVertex(screen, v, color.RGBA{0, 255, 0, 255}, toggleInfo)
	}
}

func drawLine(screen *ebiten.Image, v1, v2 la.Vector4, clr color.RGBA) {
	vector.StrokeLine(screen,
		point3Dto2D(v1.X, v1.Z+distance)*screenWidth+screenWidth/2,
		point3Dto2D(v1.Y, v1.Z+distance)*screenHeight+screenHeight/2,
		point3Dto2D(v2.X, v2.Z+distance)*screenWidth+screenWidth/2,
		point3Dto2D(v2.Y, v2.Z+distance)*screenHeight+screenHeight/2,
		1, clr, false)
}

func drawVertex(screen *ebiten.Image, v la.Vector4, clr color.RGBA, coord bool) {
	x := point3Dto2D(v.X, v.Z+distance)*screenWidth + screenWidth/2
	y := point3Dto2D(v.Y, v.Z+distance)*screenHeight + screenHeight/2
	vector.DrawFilledCircle(screen, x, y, 2, clr, false)
	if coord {
		text.Draw(screen, fmt.Sprintf("(%.1f, %.1f, %.1f)", v.X, v.Y, v.Z+distance), basicfont.Face7x13, int(x+10), int(y-20), color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("3d renderer")
	distance        = 5
	toggleInfo      = true
	currentModel    = cubeModel
	rotationRadians = (float64(1) / float64(180)) * math.Pi

	game := &Game{vertices: cube.InitCubeVertices(), keyStates: make(map[ebiten.Key]int)}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
