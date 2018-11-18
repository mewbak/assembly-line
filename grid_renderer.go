package main

import "github.com/hajimehoshi/ebiten"

type GridRenderer struct {
	sprite *ebiten.Image
}

func NewGridRenderer(sprite *ebiten.Image) *GridRenderer {
	return &GridRenderer{sprite: sprite}
}

func (r *GridRenderer) Render(screen *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	w, h := r.sprite.Size()
	op.GeoM.Translate(float64(x*w), float64(y*h))
	screen.DrawImage(r.sprite, op)
}
