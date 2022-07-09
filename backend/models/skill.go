package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Skill struct {
	ID     *primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name   string              `bson:"name" json:"name"`
	Levels []Level             `bson:"levels" json:"levels"`
}

func ShowSkill(s Skill) {
	fmt.Println("Skill:", s.Name)
	for _, l := range s.Levels {
		fmt.Println("  Level:", l.Level)
		fmt.Println("  Achievements:")
		for _, a := range l.Achievements {
			fmt.Println("    ", a.Description)
		}
	}
}
