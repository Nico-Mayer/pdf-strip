<script lang="ts">
	import { currentFile } from '../stores/store'
</script>

<main>
	{#if $currentFile}
		<p>{$currentFile.name}</p>
	{/if}
</main>
