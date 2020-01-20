package geomap

import "github.com/mitghi/geomap/storage"

// SECTION: Constants

const (
	minLon          float64 = (-180.0 * 1000)
	minLat          float64 = (-90.0 * 1000)
	latDegreeLength float64 = (111.0 * 1000)
	lonDegreeLength float64 = (85.0 * 1000)
	cEarthRadius    float64 = (6371.0 * 1000)
)

const (
	NorthEast int = iota
	East
	SouthEast
	South
	SouthWest
	West
	NorthWest
	North
)

// SECTION: Typedec

type DegreeDistance map[int]float64
type AcceptFn func(PointInterface) bool
type gridmap map[cell]storage.StoreInterface

// SECTION: Interface

type PointInterface interface {
	Id() string
	Lat() float64
	Lon() float64
	Meta() interface{}
}

// SECTION: Structs

type Point struct {
	PId  string  `json:"id"`
	PLat float64 `json:"lat"`
	PLon float64 `json:"lon"`
}

type sortedPoints struct {
	points []PointInterface
	point  PointInterface
	ld     DegreeDistance
}

type GeoMap struct {
	tiles    *grid
	position map[string]PointInterface
	ld       DegreeDistance
}

type grid struct {
	index      gridmap
	mkfn       storage.StoreBuilder
	resolution float64
}

type cell struct {
	x int
	y int
}
