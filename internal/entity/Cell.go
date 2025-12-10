package entity

import (
	"image/color"

	"github.com/cellkit/server/internal/geometry"
)

type Cell struct {
	mass     uint32
	position geometry.Vector2D
	color    color.Color
}

func (c *Cell) CanEat(e Entity) bool {
	return e.GetMass() < c.GetMass()
}

func (c *Cell) GetMass() uint32 {
	return c.mass
}

func (c *Cell) SetMass(mass uint32) {
	c.mass = mass
}

func (c *Cell) GetPosition() geometry.Vector2D {
	return c.position
}

func (c *Cell) SetPosition(pos geometry.Vector2D) {
	c.position = pos
}

func (c *Cell) GetColor() color.Color {
	return c.color
}

func (c *Cell) SetColor(color color.Color) {
	c.color = color
}
