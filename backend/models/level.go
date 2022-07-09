package models

type Level struct {
	Level        string        `bson:"level" json:"level"`
	Achievements []Achievement `bson:"achievements" json:"achievements"`
	Chalenges    []Chalenge    `bson:"chalenges" json:"chalenges"`
}

const LevelBeginner = "beginner"
const LevelIntermediate = "intermediate"
const LevelAdvanced = "advanced"
const LevelExpert = "expert"
const LevelAny = "Any"
