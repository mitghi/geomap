package geomap

import (
	"fmt"
	"testing"
)

func TestPointIndex(t *testing.T) {
	var (
		gm       *GeoMap = NewGeoMap(1.0 * 1000)
		p        PointInterface
		emb, oxf *Point
	)
	emb = NewPoint("e", -0.122367, 51.507312)
	oxf = NewPoint("o", -0.141700, 51.515110)
	gm.Add(oxf)
	gm.Add(emb)
	p = gm.Get("o")
	if p == nil {
		t.Fatal("inconsistent state.")
	}
	if p.Id() != "o" {
		t.Fatal("assertion failed.")
	}
	gm.Remove("test2")
	if np := gm.Get("test2"); np != nil {
		t.Fatal("inconsistent state.")
	}
	ritems := gm.Range(emb, oxf)
	if len(ritems) == 0 {
		t.Fatal("assertion failed.", ritems)
	}
}

func TestNearest(t *testing.T) {
	var (
		points []PointInterface = []PointInterface{
			NewPoint("e", -0.122367, 51.507212),
			NewPoint("o", -0.141700, 51.515110),
		}
		gm *GeoMap = NewGeoMap(Km(1.0))
	)
	for _, point := range points {
		gm.Add(point)
	}
	fmt.Println(gm)
	np := gm.Nearest(points[0], 4, Km(2.0), func(p PointInterface) bool { return true })
	fmt.Println("np:", np)
	if len(np) == 0 {
		t.Fatal("inconsistent state")
	}
}
