<script lang="ts">
	import type { PdfInfo } from '$types'
	import { editorOptions } from '$stores/store'

	export let info: PdfInfo

	function formatBytes(bytes: number): string {
		if (bytes < 1024) {
			return bytes + ' B'
		} else if (bytes < 1048576) {
			return Math.round(bytes / 1024) + ' KB'
		} else if (bytes < 1073741824) {
			return Math.round(bytes / 1048576) + ' MB'
		} else {
			return Math.round(bytes / 1073741824) + ' GB'
		}
	}
	function formatDate(dateString: string): string {
		return dateString
			.replace('D:', '')
			.replace(
				/(\d{4})(\d{2})(\d{2})(\d{2})(\d{2})(\d{2})/,
				'$1-$2-$3 $4:$5:$6'
			)
	}
</script>

<main
	class="flex flex-col my-4 w-full shadow border border-base-300 rounded-2xl bg-base-200">
	<div
		class="stats stats-vertical lg:stats-horizontal flex-shrink-0 bg-base-200 border-b border-base-content/10 rounded-b-none">
		<div class="stat">
			<div class="stat-title">PDF Version</div>
			<div class="stat-value">{info.Info['PDFversion']}</div>
		</div>

		<div class="stat">
			<div class="stat-title">Pages</div>
			<div class="stat-value">{info.Info['Pagecount']}</div>
		</div>

		<div class="stat">
			<div class="stat-title">Size</div>
			<div class="stat-value">{formatBytes(info.Size)}</div>
			<div class="stat-desc">{info.Size} bytes</div>
		</div>
	</div>

	<div class="collapse collapse-arrow">
		<input type="checkbox" bind:checked={$editorOptions.detailsOpen} />
		<div class="collapse-title font-semibold">Details</div>
		<div class="collapse-content">
			<div class="flex flex-col p-4 font-mono">
				{#each Object.entries(info.Info) as element}
					<div>
						{#if element[1]}
							{#if element[0] === 'Creationdate' || element[0] === 'Modificationdate'}
								{element[0]} : {formatDate(element[1])}
							{:else}
								{element[0]} : {element[1]}
							{/if}
						{/if}
					</div>
				{/each}
			</div>
		</div>
	</div>
</main>
