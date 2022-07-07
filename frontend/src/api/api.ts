import type { Skill } from './entities';

export class API {
	static async fetchRoot(): Promise<Array<Skill>> {
		const url = 'http://localhost:8000/';
		const resp = await fetch(url);
		const data = await resp.json();
		return data;
	}
}
