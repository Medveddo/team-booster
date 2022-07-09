package models

type Chalenge struct {
	AcuireTime  int32      `bson:"acuire_time" json:"acuire_time"` // in minutes
	Score       int32      `bson:"score" json:"score"`
	Description string     `bson:"description" json:"description"`
	Key         string     `bson:"key" json:"key"`
	Resources   []Resource `bson:"resources" json:"resources"`
	Checklist   []string   `bson:"checklist" json:"checklist"`
}
