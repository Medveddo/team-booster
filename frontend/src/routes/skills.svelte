<script>
    let skills = [];
	const fetchData = async () => {
		const url = 'http://localhost:8000/';
		const resp = await fetch(url);
		const data = await resp.json();
		console.log(data);
        skills = data;
	};
    fetchData();
</script>

{#each skills as s}
    <h1>{s.name}</h1>
    {#each s.levels as l}
        <h2>____{l.level}</h2>
        {#each l.achievements as a}
            <h3>_________{a.description}</h3>
        {/each}
        <br>
        {#each l.chalenges as c}
            <h3>_________{c.description}</h3>
        {/each}
    {/each}
{/each}