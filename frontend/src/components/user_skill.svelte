<script lang="ts">
	export let name = 'Skill';
	export let value = 33;
	export let color = 'primary';
	export let max = 100;
</script>

<div class="my-2 flex">
	<div class="text-left w-1/4">
		<span class="mx-2 text-sm">{name}</span>
	</div>
	<div class="text-left">
		<progress class="progress progress-{color} w-56" value="{value}" max="{max}" />
	</div>
</div>
