package main

import (
	"github.com/Medveddo/team-booster/backend/internal/models"
	"github.com/Medveddo/team-booster/backend/internal/storage"
)

func main() {
	ss := storage.NewSkillStorage()
	defer ss.Disconnect()
	// ss.InsertSkill(&models.RedisSkill)
	skills, _ := ss.GetSkills()

	for _,s := range skills {
		models.ShowSkill(s)
	}
	
}
