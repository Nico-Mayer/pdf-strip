<script lang="ts">
	import { files } from '../stores/store'

	function handleFileChange(event) {
		const fileList: FileList = event.target.files
		const fileArray = Array.from(fileList)
		files.set(fileArray)
	}
</script>

<main class="flex flex-col justify-center items-center gap-10">
	<div class="h-64 w-64 dots rounded-xl">568×320</div>
	<input
		name="userFile"
		accept="application/pdf, application/vnd.ms-excel"
		type="file"
		multiple
		on:change={handleFileChange}
		class="file-input file-input-bordered file-input-success w-full max-w-xs" />

	{#if files && files[0]}
		<p>
			{files[0].name}
		</p>
	{/if}
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
