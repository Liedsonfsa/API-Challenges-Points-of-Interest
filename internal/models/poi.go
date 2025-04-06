package models

type POI struct {
	Name 	string `json:"name" binding:"required"`
	X 		uint64 `json:"x" binding:"required"`
	Y 		uint64 `json:"y" binding:"required"`
}