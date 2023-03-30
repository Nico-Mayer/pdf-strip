<script lang="ts">
	import SideBtn from './SideBtn.svelte'
	export let files: File[] = null
</script>

{#if files && files[0]}
	<main class="bg-base-200 w-52 h-full">
		<ul class="flex flex-col gap-1 p-1">
			{#each files as file}
				<SideBtn {file} />
			{/each}
		</ul>
	</main>
{/if}
