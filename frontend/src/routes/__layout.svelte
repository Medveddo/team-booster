<script lang="ts">
	import Navbar from '../components/navbar.svelte';
</script>

<svelte:head>
	<title>TeamBooster</title>
</svelte:head>

<Navbar />
<slot />

<style>
	@tailwind base;
	@tailwind components;
	@tailwind utilities;
</style>
