package models

import "math"

type POI struct {
	ID   uint64 `json:"id"`
	Name string `json:"name" binding:"required"`
	X    uint64 `json:"x" binding:"required"`
	Y    uint64 `json:"y" binding:"required"`
}

type GPOI struct {
	D_Max 	uint64 `json:"d_max" binding:"required"`
	X    	uint64 `json:"x" binding:"required"`
	Y   	uint64 `json:"y" binding:"required"`
}

func (p *POI) WithinTheDistance(x, y, d_max uint64) bool {
	distance := math.Sqrt(float64(((x - p.X) * (x - p.X)) + ((y - p.Y) * (y - p.Y))))

	return distance <= float64(d_max)
}