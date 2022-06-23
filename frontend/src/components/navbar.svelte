<script lang="ts">
	let is_logged_in = true;
</script>

<div class="navbar bg-base-100">
	<div class="flex-1">
		<a href="/" class="btn btn-ghost normal-case text-xl">TeamBooster</a>
	</div>
	<div class="flex-none">
		<ul class="menu menu-horizontal p-0">
			<li><a href="/skills">Skills</a></li>
			<li><a href="/challenges">Challenges</a></li>
			<li><a href="/team">Team</a></li>
			{#if is_logged_in}
				<li><a href="/account">Account</a></li>
			{:else}
				<li><a href="/account">Sign Up / Login</a></li>
			{/if}
		</ul>
	</div>
</div>
