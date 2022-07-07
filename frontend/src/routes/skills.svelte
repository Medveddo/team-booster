<script lang="ts">
	import { API } from '../api/api';
	import type { Skill } from '../api/entities';

	let skills: Array<Skill> = [];

	API.fetchRoot().then((data: Array<Skill>) => {
		console.log(data);
		skills = data;
	});
</script>

{#each skills as s}
	<h1>{s.name}</h1>
	{#each s.levels as l}
		<h2>__{l.level}</h2>
		{#each l.achievements as a}
			<h3>___{a.description}</h3>
            {#each a.resources as r}
            <h4>____{r.location}</h4>
            {/each}
		{/each}
		<br />
		{#each l.chalenges as c}
			<h3>___{c.description}</h3>
		{/each}
	{/each}
{/each}
