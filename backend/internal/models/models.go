package models

import "fmt"

type Resource struct {
	Name     string `bson:"name" json:"name"`
	Location string `bson:"location" json:"location"`
}

type Achievement struct {
	AcuireTime  int32      `bson:"acuire_time" json:"acuire_time"` // in minutes
	Score       int32      `bson:"score" json:"score"`
	Description string     `bson:"description" json:"description"`
	Key         string     `bson:"key" json:"key"`
	Resources   []Resource `bson:"resources" json:"resources"`
}

type Chalenge struct {
	AcuireTime  int32      `bson:"acuire_time" json:"acuire_time"` // in minutes
	Score       int32      `bson:"score" json:"score"`
	Description string     `bson:"description" json:"description"`
	Key         string     `bson:"key" json:"key"`
	Resources   []Resource `bson:"resources" json:"resources"`
	Checklist   []string     `bson:"checklist" json:"checklist"`
}

type Level struct {
	Level        string        `bson:"level" json:"level"`
	Achievements []Achievement `bson:"achievements" json:"achievements"`
	Chalenges    []Chalenge    `bson:"chalenges" json:"chalenges"`
}

type Skill struct {
	Name   string  `bson:"name" json:"name"`
	Levels []Level `bson:"levels" json:"levels"`
}

const BEGINNER_LEVEL = "beginner"
const INTERMEDIATE_LEVEL = "intermediate"
const ADVANCED_LEVEL = "advanced"
const EXPERT_LEVEL = "expert"
const FREE_TO_LEARN = "free_to_learn"

func ShowSkill (s Skill) {
	fmt.Println("Skill:", s.Name)
	for _, l := range s.Levels {
		fmt.Println("  Level:", l.Level)
		fmt.Println("  Achievements:")
		for _, a := range l.Achievements {
			fmt.Println("    ", a.Description)
		}
	}
}