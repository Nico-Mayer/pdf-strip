<script lang="ts">
	import { editorOptions } from '$stores/store'

	let compressionLevel = 25
	let saveAsCopy = false

	$: level =
		compressionLevel === 25
			? 'low'
			: compressionLevel === 50
			? 'medium'
			: compressionLevel === 75
			? 'high'
			: 'ultra'
</script>

<div
	class="collapse collapse-arrow border border-base-300 bg-base-100 rounded-box flex-shrink-0">
	<input type="checkbox" bind:checked={$editorOptions.compressionOpen} />
	<div class="collapse-title text-xl font-medium">Compression</div>

	<div class="collapse-content">
		<section class="flex flex-col gap-4">
			<div class="flex justify-between">
				<span class="font-bold">Level:</span>
				<span
					class="badge"
					class:badge-info={level === 'low'}
					class:badge-success={level === 'medium'}
					class:badge-warning={level === 'high'}
					class:badge-error={level === 'ultra'}>{level}</span>
			</div>

			<input
				type="range"
				min="25"
				max="100"
				bind:value={compressionLevel}
				class="range"
				step="25" />
			<div class="w-full flex justify-between text-xs px-2">
				<span>|</span>
				<span>|</span>
				<span>|</span>
				<span>|</span>
			</div>

			<div class="flex gap-4">
				<input
					type="checkbox"
					class="toggle"
					bind:checked={saveAsCopy} />
				<span>Save as Copy</span>

				{#if saveAsCopy}
					<input
						type="text"
						class="input input-bordered"
						placeholder="Filename" />
				{/if}
			</div>

			<button class="btn"> Compress </button>
		</section>
	</div>
</div>
