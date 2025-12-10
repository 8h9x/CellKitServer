package entity

import "github.com/cellkit/server/internal/geometry"

type Entity interface {
	GetMass() uint32
	SetMass(uint32)

	GetPosition() geometry.Vector2D
	SetPosition(geometry.Vector2D)

	CanEat(Entity) bool
}
