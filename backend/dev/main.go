package main

import (
	"github.com/Medveddo/team-booster/backend/api"
)

func main() {
	api := api.NewAPI()
	_ = api
	// ss := storage.NewSkillStorage()
	// defer ss.Disconnect()
	// // ss.InsertSkill(&models.RedisSkill)
	// skills, _ := ss.GetSkills()

	// for _,s := range skills {
	// 	models.ShowSkill(s)
	// }
}
