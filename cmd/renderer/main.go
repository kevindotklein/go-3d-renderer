package main

import (
	"image/color"
	"log"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/kevindotklein/go-3d-renderer/pkg/la"
	"golang.org/x/image/font/basicfont"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
	fov          = 45

	uiMessage = 
	"[esc]    to exit\n" +
	"[w]      to move camera foward\n"+
	"[s]      to move camera backward\n"+
	"[a]      to move camera to the left\n"+
	"[d]      to move camera to the right\n"+
	"[space]  to move camera upward\n"+
	"[lshift] to move camera downward\n"
)

var cubeDistance float32 

func initCubeVertices() []la.Vector4 {
	return []la.Vector4{
		{X: -0.5, Y:  0.5, Z: -0.5, W: 1.0},
		{X:  0.5, Y:  0.5, Z: -0.5, W: 1.0},
		{X:  0.5, Y:  0.5, Z:  0.5, W: 1.0},
		{X: -0.5, Y:  0.5, Z:  0.5, W: 1.0},
		{X: -0.5, Y: -0.5, Z: -0.5, W: 1.0},
		{X:  0.5, Y: -0.5, Z: -0.5, W: 1.0},
		{X:  0.5, Y: -0.5, Z:  0.5, W: 1.0},
		{X: -0.5, Y: -0.5, Z:  0.5, W: 1.0},
		{X: -0.5, Y: -0.5, Z:  0.5, W: 1.0},
		{X:  0.5, Y: -0.5, Z:  0.5, W: 1.0},
		{X:  0.5, Y:  0.5, Z:  0.5, W: 1.0},
		{X: -0.5, Y:  0.5, Z:  0.5, W: 1.0},
		{X: -0.5, Y: -0.5, Z: -0.5, W: 1.0},
		{X:  0.5, Y: -0.5, Z: -0.5, W: 1.0},
		{X:  0.5, Y:  0.5, Z: -0.5, W: 1.0},
		{X: -0.5, Y:  0.5, Z: -0.5, W: 1.0},
	}
}

type Game struct {
	cubeVertices []la.Vector4
}

func point3Dto2D(point, z float32) float32 {
	radians := (float64(fov) / float64(180)) * math.Pi
	return point / (z * float32(math.Tan(radians/2)))
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		cubeDistance -= 0.2
	}else if ebiten.IsKeyPressed(ebiten.KeyS) {
		cubeDistance += 0.2
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		m := la.Matrix4{A: la.Vector4{X: 1.0}, B: la.Vector4{Y: 1.0, W: 0.02}, C: la.Vector4{Z: 1.0}, D: la.Vector4{W: 1.0}}
		for i := range g.cubeVertices {
			g.cubeVertices[i].Dot(m)
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyShiftLeft) {
		m := la.Matrix4{A: la.Vector4{X: 1.0}, B: la.Vector4{Y: 1.0, W: -0.02}, C: la.Vector4{Z: 1.0}, D: la.Vector4{W: 1.0}}
		for i := range g.cubeVertices {
			g.cubeVertices[i].Dot(m)
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		m := la.Matrix4{A: la.Vector4{X: 1.0, W: 0.02}, B: la.Vector4{Y: 1.0}, C: la.Vector4{Z: 1.0}, D: la.Vector4{W: 1.0}}
		for i := range g.cubeVertices {
			g.cubeVertices[i].Dot(m)
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		m := la.Matrix4{A: la.Vector4{X: 1.0, W: -0.02}, B: la.Vector4{Y: 1.0}, C: la.Vector4{Z: 1.0}, D: la.Vector4{W: 1.0}}
		for i := range g.cubeVertices {
			g.cubeVertices[i].Dot(m)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < 4; i++ {
		drawLine(screen, g.cubeVertices[i], g.cubeVertices[(i+1)%4], color.RGBA{255, 0, 0, 255})
	}

	for i := 4; i < 8; i++ {
		drawLine(screen, g.cubeVertices[i], g.cubeVertices[4+(i-4+1)%4], color.RGBA{255, 0, 0, 255})
	}

	// vertical lines
	for i := 0; i < 4; i++ {
		drawLine(screen, g.cubeVertices[i], g.cubeVertices[i+4], color.RGBA{255, 0, 0, 255})
	}

	for _, v := range g.cubeVertices {
		drawVertex(screen, v, color.RGBA{0, 255, 0, 255})
	}

	text.Draw(screen, uiMessage, basicfont.Face7x13, 20, 30, color.White)	
}

func drawLine(screen *ebiten.Image, v1, v2 la.Vector4, clr color.RGBA) {
	vector.StrokeLine(screen,
		point3Dto2D(v1.X, v1.Z+cubeDistance)*screenWidth+screenWidth/2,
		point3Dto2D(v1.Y, v1.Z+cubeDistance)*screenHeight+screenHeight/2,
		point3Dto2D(v2.X, v2.Z+cubeDistance)*screenWidth+screenWidth/2,
		point3Dto2D(v2.Y, v2.Z+cubeDistance)*screenHeight+screenHeight/2,
		1, clr, false)
}

func drawVertex(screen *ebiten.Image, v la.Vector4, clr color.RGBA) {
	vector.DrawFilledCircle(screen,
		point3Dto2D(v.X, v.Z+cubeDistance)*screenWidth+screenWidth/2,
		point3Dto2D(v.Y, v.Z+cubeDistance)*screenHeight+screenHeight/2,
		2, clr, false)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("3d renderer")
	cubeDistance = 10

	game := &Game{cubeVertices: initCubeVertices()}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
