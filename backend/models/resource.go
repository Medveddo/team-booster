package models

type Resource struct {
	Name     string `bson:"name" json:"name"`
	Location string `bson:"location" json:"location"`
}
