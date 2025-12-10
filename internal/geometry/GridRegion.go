package geometry

import "fmt"

type GridRegion struct {
	Row, Column int8
}

func (r GridRegion) String() string {
	return fmt.Sprintf("%d,%d", r.Row, r.Column)
}
