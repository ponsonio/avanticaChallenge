package models

import "github.com/Kamva/mgm/v3"

type Point struct {
	X float64 `json:"x" bson:"x"`
	Y float64 `json:"y" bson:"y"`
}

type Spot struct {
	mgm.DefaultModel `bson:",inline"`
	Point            Point   `json:"point" bson:"point"`
	Name             string  `json:"name" bson:"name"`
	Gold             float64 `json:"gold" bson:"gold"`
	Quadrant         string  `json:"quadrant" bson:"quadrant"`
}

type Path struct {
	mgm.DefaultModel `bson:",inline"`
	Distance         float64 `json:"distance" bson:"distance"`
	SpotA            string  `json:"a" bson:"spotA"`
	SpotB            string  `json:"b" bson:"spotB"`
}

type Quadrant struct {
	mgm.DefaultModel `bson:",inline"`
	Center           Point `json:"center" bson:"center"`
	Number           int   `json:"number" bson:"number"`
}
