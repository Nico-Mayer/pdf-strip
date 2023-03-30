<script lang="ts">
	import SideBtn from './SideBtn.svelte'
	import { filePaths } from '../../stores/store'
</script>

{#if $filePaths && $filePaths[0]}
	<main class="bg-base-200 w-52 h-full">
		<ul class="flex flex-col gap-1 p-1">
			{#each $filePaths as path}
				<SideBtn {path} />
			{/each}
		</ul>
	</main>
{/if}
