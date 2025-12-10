package geometry

func NewSpatialGrid(BoundWidth, BoundHeight int, RegionSize uint8) *SpatialGrid {
	return &SpatialGrid{
		BoundWidth:  BoundWidth,
		BoundHeight: BoundHeight,
		RegionSize:  RegionSize,
		Buckets:     make(map[string]map[any]struct{}),
	}
}

type SpatialGrid struct {
	BoundWidth  int
	BoundHeight int
	RegionSize  uint8
	Buckets     map[string]map[any]struct{}
}

func (g *SpatialGrid) Insert(coords Vector2D, entity any) {
	region := coords.GridRegion(g.RegionSize)
	key := region.String()

	bucket, exists := g.Buckets[key]
	if !exists {
		bucket = make(map[any]struct{}, 8) // init with a bit of length to ideally help bursty areas
		g.Buckets[key] = bucket
	}

	bucket[entity] = struct{}{}
}

func (g *SpatialGrid) Remove(entity Vector2D) {
	region := entity.GridRegion(g.RegionSize)
	key := region.String()

	bucket, exists := g.Buckets[key]
	if !exists {
		return
	}

	delete(bucket, entity)

	if len(bucket) == 0 {
		delete(g.Buckets, key)
	}
}

func (g *SpatialGrid) Reposition() {

}
