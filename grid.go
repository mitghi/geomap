package geomap

import "github.com/mitghi/geomap/storage"

func newGrid(resolution float64, mkfn storage.StoreBuilder) *grid {
	var g *grid = &grid{
		index:      make(map[cell]storage.StoreInterface),
		mkfn:       mkfn,
		resolution: resolution,
	}
	return g
}

func (g *grid) Clone() *grid {
	var (
		clone *grid
	)
	clone = &grid{
		index:      make(map[cell]storage.StoreInterface),
		mkfn:       g.mkfn,
		resolution: g.resolution,
	}
	for k, v := range g.index {
		s, ok := v.(storage.StoreInterface)
		if !ok {
			panic("unreachable")
		}
		clone.index[k] = s.Clone()
	}
	return clone
}

func (g *grid) AddEntry(point PointInterface) storage.StoreInterface {
	var (
		tile  storage.StoreInterface
		index cell = CalcCell(point, g.resolution)
	)
	if v, ok := g.index[index]; ok {
		return v
	} else {
		tile = g.mkfn()
		g.index[index] = tile
		return tile
	}
}

func (g *grid) GetEntry(point PointInterface) (tile storage.StoreInterface) {
	var (
		index cell = CalcCell(point, g.resolution)
		ok    bool
	)
	tile, ok = g.index[index]
	if !ok {
		return g.mkfn()
	}

	return tile
}

func (g *grid) Range(topLeft PointInterface, bottomRight PointInterface) []storage.StoreInterface {
	var (
		tlIndex cell = CalcCell(topLeft, g.resolution)
		brIndex cell = CalcCell(bottomRight, g.resolution)
	)
	return g.get(brIndex.x, tlIndex.x, tlIndex.y, brIndex.y)
}

func (g *grid) Cells(minx int, maxx int, miny int, maxy int) (cells []cell) {
	cells = make([]cell, 0)
	for x := minx; x <= maxx; x++ {
		for y := miny; y <= maxy; y++ {
			cells = append(cells, cell{x, y})
		}
	}
	return cells
}

func (g *grid) get(minx int, maxx int, miny int, maxy int) (entries []storage.StoreInterface) {
	entries = make([]storage.StoreInterface, 0, 0)
	for x := minx; x <= maxx; x++ {
		for y := miny; y <= maxy; y++ {
			if index, ok := g.index[cell{x, y}]; ok {
				entries = append(entries, index)
			}
		}
	}
	return entries
}
