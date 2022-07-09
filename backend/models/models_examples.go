package models

var resource Resource = Resource{Name: "Redis Data types tutorial", Location: "https://redis.io/docs/manual/data-types/data-types-tutorial/"}
var resource2 Resource = Resource{Name: "The little redis book - Chapter 2 - The Data Structures", Location: "https://github.com/karlseguin/the-little-redis-book/blob/master/en/redis.md#chapter-2---the-data-structures"}

var achievement Achievement = Achievement{
	AcuireTime:  10,
	Score:       5,
	Resources:   []Resource{resource, resource2},
	Description: "Know basic Redis commands for read and write data into it",
	Key:         "redis_basic_commands",
}

var achievement2 Achievement = Achievement{
	AcuireTime:  10,
	Score:       5,
	Resources:   []Resource{},
	Description: "Know about Redis expire, how to set it and check for TTL",
	Key:         "redis_expire_basic",
}

var chalenge Chalenge = Chalenge{
	Key:         "redis_docker_install",
	Score:       10,
	AcuireTime:  2,
	Description: "Get up and running Redis in Doker. Connect to redis-server via redis-cli",
	Checklist:   []string{},
	Resources:   []Resource{},
}

var chalenge2 Chalenge = Chalenge{
	Key: "basic_data_types_practice",
	Checklist: []string{
		"Установить ключ, прочтитать, удалить",
		"Создать список, пушануть в него слева справа, вывести список, вытащить значение",
		"Увеличить счётчик на 1, 10",
		"Установить хеш, вывести",
	},
	Description: "Play with Redis basic types",
	Score:      10,
	AcuireTime: 2,
	Resources:   []Resource{},
}

var beginner_level Level = Level{
	Level:        LevelBeginner,
	Achievements: []Achievement{achievement, achievement2},
	Chalenges:    []Chalenge{chalenge, chalenge2},
}

var inter_level Level = Level{
	Level: LevelIntermediate,
	Achievements: []Achievement{},
	Chalenges:    []Chalenge{},
}

var advanced_level = Level{
	Level:        LevelAdvanced,
	Achievements: []Achievement{},
	Chalenges:    []Chalenge{},
}

var RedisSkill Skill = Skill{Name: "Redis", Levels: []Level{beginner_level, inter_level, advanced_level}}
