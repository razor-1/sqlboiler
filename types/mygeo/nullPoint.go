package mygeo

import (
	"database/sql/driver"
)

// NullPoint allows point to be null
type NullPoint struct {
	Point
	Valid bool `json:"valid"`
}

func NullPointFrom(p Point) NullPoint {
	return NullPoint{Point: p, Valid: true}
}

// Value for database
func (p NullPoint) Value() (driver.Value, error) {
	if !p.Valid {
		return nil, nil
	}

	return valuePoint(p.Point)
}

// Scan from sql query
func (p *NullPoint) Scan(src interface{}) error {
	if src == nil {
		p.Point, p.Valid = NewPoint(0, 0, 0), false
		return nil
	}

	p.Valid = true
	return scanPoint(&p.Point, src)
}

// Randomize for sqlboiler
func (p *NullPoint) Randomize(nextInt func() int64, fieldType string, shouldBeNull bool) {
	if shouldBeNull {
		p.Valid = false
		return
	}

	p.Valid = true
	p.Point = randPoint(nextInt)
}
