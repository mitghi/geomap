package geomap

import "math"

// SECTION: Internal

func toDegrees(x float64) float64 {
	return x * 180.0 / math.Pi
}

func toRadians(x float64) float64 {
	return x * math.Pi / 180.0
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func euclidDistance(disl DegreeDistance, lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {
	var (
		latLen float64 = math.Abs(lat1-lat2) * latDegreeLength
		lonLen float64 = math.Abs(lon1-lon2) * disl.avgLatitudeLength(lat1, lat2)
	)
	return (latLen*latLen + lonLen*lonLen)
}

func pCoordinate(lat float64, lon float64, res float64) (x int, y int) {
	x = int((-minLat + lat) * latDegreeLength / res)
	y = int((-minLon + lon) * lonDegreeLength / res)
	return x, y
}

func directionTo(bearing float64) int {
	var (
		index float64 = bearing - 22.5
	)
	if index < 0 {
		index += 360
	}
	return int(index / 45.0)
}

func bearingTo(lat1, lon1, lat2, lon2 float64) float64 {
	var (
		disLon float64 = toRadians(lon2 - lon1)
		rlat1  float64 = toRadians(lat1)
		rlat2  float64 = toRadians(lat2)
		x      float64 = math.Cos(rlat1)*math.Sin(rlat2) - math.Sin(rlat1)*math.Cos(rlat2)*math.Cos(disLon)
		y      float64 = math.Sin(disLon) * math.Cos(rlat2)
	)
	return toDegrees(math.Atan2(y, x))
}

func distance(lat1, lon1, lat2, lon2 float64) float64 {
	var (
		sLat float64 = math.Sin(toRadians(lat2-lat1) / 2)
		sLon float64 = math.Sin(toRadians(lon2-lon1) / 2)
		a    float64
		c    float64
	)
	a = math.Pow(sLat, 2) + math.Pow(sLon, 2)*math.Cos(toRadians(lat1)*math.Cos(toRadians(lat2)))
	c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return cEarthRadius * c
}

func isBetween(value float64, min float64, max float64) bool {
	return (value >= min) && (value <= max)
}

func CalcCell(p PointInterface, res float64) cell {
	var (
		x int
		y int
	)
	x, y = pCoordinate(p.Lat(), p.Lon(), res)
	return cell{x, y}
}

// - MARK: DegreeDistance section.

func (dd DegreeDistance) get(lat float64) (dist float64) {
	var (
		latIndex int     = int(lat * 10)
		latRnd   float64 = float64(latIndex) / 10
	)
	if v, ok := dd[latIndex]; ok {
		return v
	}
	dist = distance(latRnd, 0.0, latRnd, 1.0)
	dd[latIndex] = dist
	return dist
}

func (dd DegreeDistance) avgLatitudeLength(lat1, lat2 float64) float64 {
	var (
		avg float64 = (lat1 + lat2) / 2.0
	)
	return dd.get(avg)
}

// SECTION: Public

func Km(km float64) float64 {
	return km * 1000
}

func EuclidDistance(disl DegreeDistance, p1 PointInterface, p2 PointInterface) float64 {
	return euclidDistance(disl, p1.Lat(), p1.Lon(), p2.Lat(), p2.Lon())
}

func DirectionTo(p1 PointInterface, p2 PointInterface) int {
	var (
		bearing float64 = BearingTo(p1, p2)
	)
	return directionTo(bearing)
}

func BearingTo(p1 PointInterface, p2 PointInterface) float64 {
	return bearingTo(p1.Lat(), p1.Lon(), p2.Lat(), p2.Lon())
}

func Distance(p1 PointInterface, p2 PointInterface) float64 {
	return distance(p1.Lat(), p1.Lon(), p2.Lat(), p2.Lon())
}
