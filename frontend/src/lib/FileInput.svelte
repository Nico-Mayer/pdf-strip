<script lang="ts">
</script>

<main class="flex flex-col justify-center items-center gap-10">
	<div class="h-64 w-64 dots rounded-xl">568×320</div>
	<input
		type="file"
		accept="pd"
		class="file-input file-input-bordered file-input-success w-full max-w-xs" />
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
