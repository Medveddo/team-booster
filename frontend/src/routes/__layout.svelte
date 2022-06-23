<script lang="ts">
	import Navbar from '../components/navbar.svelte';
</script>

<Navbar />
<slot />

<style>
	@tailwind base;
	@tailwind components;
	@tailwind utilities;
</style>
