<script lang="ts">
	import { currentFile } from '$stores/store'

	$: currentPath = $currentFile?.path
	$: fileName = $currentFile?.name
	$: breadcrumbs = formatPath(currentPath)

	function formatPath(path): string[] {
		if (!path) return []
		const separator = path.includes('/') ? '/' : '\\'
		return path.split(separator).slice(0, -1)
	}
</script>

<div class="text-base breadcrumbs overflow-x-auto hide-scrollbar">
	<ul>
		{#each breadcrumbs as breadcrumb, index}
			<li>
				{#if index === 0}
					<div class="i-mdi-harddisk mr-2" />
				{:else}
					<div class="i-mdi-folder mr-2" />
				{/if}

				{breadcrumb}
			</li>
		{/each}
		<li>
			<div class="i-mdi-file-pdf mr-2" />
			{fileName}
		</li>
	</ul>
</div>
