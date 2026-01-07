package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	Type        string     `bson:"type" json:"type"`
	Coordinates [2]float64 `bson:"coordinates" json:"coordinates"`
}

type RumahSakit struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Nama     string             `bson:"nama" json:"nama"`
	Location Location           `bson:"location" json:"location"`
}
