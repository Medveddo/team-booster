export class Resource {
	name: string;
	location: string;

	constructor(name: string, location: string) {
		this.name = name;
		this.location = location;
	}
}

export class Achievement {
	acuire_time: number;
	score: number;
	description: string;
	key: string;
	resources: Array<Resource>;

	constructor(
		acuire_time: number,
		score: number,
		description: string,
		key: string,
		resources: Array<Resource>
	) {
		this.acuire_time = acuire_time;
		this.score = score;
		this.description = description;
		this.key = key;
		this.resources = resources;
	}
}

export class Chalenge {
	acuire_time: number;
	score: number;
	description: string;
	key: string;
	resources: Array<Resource>;
	checklist: Array<string>;

	constructor(
		acuire_time: number,
		score: number,
		description: string,
		key: string,
		resources: Array<Resource>,
		checklist: Array<string>
	) {
		this.acuire_time = acuire_time;
		this.score = score;
		this.description = description;
		this.key = key;
		this.resources = resources;
		this.checklist = checklist;
	}
}

export class Level {
	level: string;
	achievements: Array<Achievement>;
	chalenges: Array<Chalenge>;

	constructor(level: string, achievements: Array<Achievement>, chalenges: Array<Chalenge>) {
		this.level = level;
		this.achievements = achievements;
		this.chalenges = chalenges;
	}
}

export class Skill {
	name: string;
	levels: Array<Level>;

	constructor(name: string, levels: Array<Level>) {
		this.name = name;
		this.levels = levels;
	}
}
