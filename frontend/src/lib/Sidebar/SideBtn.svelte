<script lang="ts">
	import { currentFile, filePaths } from '$stores/store'
	import { getNameFromPath } from '$utils'

	export let path = null

	$: active = $currentFile.path === path

	function setCurrentFile() {
		currentFile.pick(path)
	}
	function handleClose() {
		filePaths.remove(path)
		if ($filePaths.size > 0) {
			const firstPath = [...$filePaths][0]
			currentFile.pick(firstPath)
		}
	}
</script>

<button
	class="btn-ghost btn w-full gap-2 justify-start flex-nowrap group"
	on:click={setCurrentFile}
	class:btn-active={active}>
	<div class="i-mdi-file-pdf icon" />

	<p class="truncate flex-1 text-start">{getNameFromPath(path)}</p>

	<button
		class="btn btn-square btn-xs hover:btn-error group-hover:visible invisible"
		on:click|stopPropagation={handleClose}>
		<svg
			xmlns="http://www.w3.org/2000/svg"
			class="h-4 w-4"
			fill="none"
			viewBox="0 0 24 24"
			stroke="currentColor"
			><path
				stroke-linecap="round"
				stroke-linejoin="round"
				stroke-width="2"
				d="M6 18L18 6M6 6l12 12" /></svg>
	</button>
</button>
