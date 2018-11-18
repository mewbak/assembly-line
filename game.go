package main

import "github.com/hajimehoshi/ebiten"

type Game struct {
	resources *ResourceManager

	objects []Object
}

func NewGame() (*Game, error) {
	g := &Game{
		resources: NewResourceManager(),
	}

	sprites := []string{
		"cell",
	}

	for _, s := range sprites {
		err := g.resources.Sprites.Load(s)
		if err != nil {
			return nil, err
		}
	}

	grid := NewGrid(
		10, 10,
		NewGridRenderer(g.resources.Sprites.Get("cell")),
	)
	g.AddObject(grid)

	return g, nil
}

func (g *Game) AddObject(obj Object) {
	g.objects = append(g.objects, obj)
}

func (g *Game) Update(screen *ebiten.Image) error {
	for _, obj := range g.objects {
		obj.Render(screen)
	}
	return nil
}
