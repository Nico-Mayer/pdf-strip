<script lang="ts">
	import { filePaths } from '../stores/store'
	import { OpenFile } from '../../wailsjs/go/main/App'

	async function handleClick() {
		const path = await OpenFile()
		const oldPaths = $filePaths
		console.log(oldPaths)
		$filePaths = [...oldPaths, path]
	}
</script>

<main class="flex flex-col justify-center items-center gap-10">
	<div class="h-64 w-64 dots rounded-xl">568×320</div>
	<button on:click={handleClick}>Open</button>
</main>

<style scoped>
	.dots {
		background-image: radial-gradient(
			hsla(var(--bc) / 0.2) 0.5px,
			hsla(var(--b2) / 1) 0.5px
		);
		background-size: 5px 5px;
	}
</style>
