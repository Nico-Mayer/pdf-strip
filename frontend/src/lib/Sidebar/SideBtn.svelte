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
		class="btn btn-xs btn-outline btn-error invisible group-hover:visible self"
		on:click={handleClose}>
		<div class="i-mdi-close" />
	</button>
</button>
