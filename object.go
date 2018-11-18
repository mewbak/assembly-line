package main

import "github.com/hajimehoshi/ebiten"

type Object interface {
	Render(screen *ebiten.Image)
}

type Renderer interface {
	Render(screen *ebiten.Image, x, y int)
}
