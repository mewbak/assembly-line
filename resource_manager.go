package main

import (
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten"
)

type SpritesManager struct {
	sprites map[string]*ebiten.Image
}

func (sm *SpritesManager) Get(name string) *ebiten.Image {
	return sm.sprites[name]
}

func (sm *SpritesManager) Load(name string) error {
	reader, err := os.Open("data/sprites/" + name + ".png")
	if err != nil {
		return err
	}

	img, _, err := image.Decode(reader)
	if err != nil {
		return err
	}

	sm.sprites[name], err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	return err
}

type ResourceManager struct {
	Sprites *SpritesManager
}

func NewResourceManager() *ResourceManager {
	rm := &ResourceManager{}
	rm.Sprites = &SpritesManager{
		sprites: map[string]*ebiten.Image{},
	}
	return rm
}
