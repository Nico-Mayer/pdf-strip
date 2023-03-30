<script lang="ts">
	import { currentFile } from '../../stores/store'

	export let file = null

	function setCurrentFile() {
		currentFile.set(file)
	}
</script>

<button class="btn btn-ghost flex" on:click={setCurrentFile}
	>{file?.name}</button>
