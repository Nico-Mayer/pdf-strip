<script lang="ts">
	import SideBtn from './SideBtn.svelte'
	import { filePaths } from '$stores/store'
	import { getFilePaths } from '$utils'

	async function handleClick() {
		const paths = await getFilePaths()
		filePaths.add(paths)
	}
</script>

{#if $filePaths.size > 0}
	<main class="flex flex-col w-80 bg-base-200 shrink-0 p-4 justify-between">
		<ul class="menu gap-2">
			{#each Array.from($filePaths) as path}
				<SideBtn {path} />
			{/each}
		</ul>
		<button class="btn btn-primary" on:click={handleClick}>Add</button>
	</main>
{/if}
