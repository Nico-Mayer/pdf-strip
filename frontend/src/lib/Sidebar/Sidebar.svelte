<script lang="ts">
	import SideBtn from './SideBtn.svelte'
	import { filePaths, currentFile } from '$stores/store'
	import { getFilePaths } from '$utils'

	async function handleClick() {
		const paths = await getFilePaths()
		if ($filePaths.size === 0) {
			currentFile.pick(paths[0])
		}
		filePaths.add(paths)
	}
</script>

<main class="flex flex-col w-80 bg-base-200">
	<section class="px-4 pb-4 prose">
		<h1 class="px-4 text-primary text-3xl">
			PDF <span class="text-base-content">Strip</span>
		</h1>
	</section>

	<section class="flex flex-col overflow-y-auto p-4 justify-between h-full">
		<ul class="menu gap-2">
			{#each Array.from($filePaths) as path}
				<SideBtn {path} />
			{/each}
		</ul>
		<button class="btn btn-primary mt-4" on:click={handleClick}
			>open</button>
	</section>
</main>
