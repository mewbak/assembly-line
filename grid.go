package main

import "github.com/hajimehoshi/ebiten"

type Grid struct {
	w, h int

	objects  [][]Object
	renderer Renderer
}

func NewGrid(w, h int, renderer Renderer) *Grid {
	g := &Grid{w: w, h: h, renderer: renderer}
	g.objects = make([][]Object, h)
	for i := range g.objects {
		g.objects[i] = make([]Object, w)
	}
	return g
}

func (g *Grid) Place(x, y int, obj Object) {
	g.objects[y][x] = obj
}

func (g *Grid) Render(screen *ebiten.Image) {
	for y, row := range g.objects {
		for x, _ := range row {
			g.renderer.Render(screen, x, y)
		}
	}
}
