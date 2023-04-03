package mygeo

import (
	"database/sql/driver"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/ewkb"
)

const defaultSRID = 4326

// Point is the fundamental two-dimensional building block for geometric types.
// X and Y are the respective coordinates, as floating-point numbers
type Point struct {
	orb.Point
	// X    float64 `json:"x"`
	// Y    float64 `json:"y"`
	SRID int `json:"srid"`
}

// NewPoint creates a point
func NewPoint(X, Y float64, SRID int) Point {
	return Point{
		Point: orb.Point{X, Y},
		SRID:  SRID,
	}
}

// Value representation for database
func (p Point) Value() (driver.Value, error) {
	return valuePoint(p)
}

// Scan from query
func (p *Point) Scan(src interface{}) error {
	return scanPoint(p, src)
}

func valuePoint(p Point) (driver.Value, error) {
	return ewkb.ValuePrefixSRID(p.Point, p.SRID).Value()
}

func scanPoint(p *Point, src interface{}) error {
	if src == nil {
		*p = NewPoint(0, 0, 0)
		return nil
	}

	gs := ewkb.ScannerPrefixSRID(p)
	return gs.Scan(src)
}

func newRandNum(nextInt func() int64) float64 {
	return float64(nextInt())
}

func randPoint(nextInt func() int64) Point {
	return NewPoint(newRandNum(nextInt), newRandNum(nextInt), defaultSRID)
}

// Randomize for sqlboiler
func (p *Point) Randomize(nextInt func() int64, fieldType string, shouldBeNull bool) {
	*p = randPoint(nextInt)
}
