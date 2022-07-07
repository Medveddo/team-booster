<script context="module">
	export function load({ error, status }) {
		return {
			props: {
				title: `${status}: ${error.message}`,
                status: status
			}
		};
	}
</script>

<script>
	export let title;
	export let status;
</script>

{#if status == 404}
	
    <div class="flex justify-center">
        <div class="mockup-window border bg-base-300 w-96">
            <div class="flex px-4 py-16 bg-base-200 justify-center">
                <span class="text-2xl mr-5">404 Not found</span>
                <a class="btn btn-info" href="/">Go home</a>
            </div>
          </div>
    </div>
{:else}
	<h1>{title}</h1>
{/if}
