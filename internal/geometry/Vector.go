package geometry

type Vector2D struct {
	X, Y int32
}

func (v *Vector2D) Distance(neighbor Vector2D) {
}

func (v *Vector2D) GridRegion(RegionSize uint8) GridRegion {
	return GridRegion{
		Row:    int8(v.X) / int8(RegionSize),
		Column: int8(v.Y) / int8(RegionSize),
	}
}
