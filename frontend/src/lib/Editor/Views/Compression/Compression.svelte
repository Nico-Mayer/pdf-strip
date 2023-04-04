<script lang="ts">
	import { editorOptions } from '$stores/store'
	import { OpenDir } from '$go/App'

	let compressionLevel = 25
	let saveAsCopy = false
	let safeTo = ''

	$: level =
		compressionLevel === 25
			? 'low'
			: compressionLevel === 50
			? 'medium'
			: compressionLevel === 75
			? 'high'
			: 'ultra'

	$: submittable = (safeTo.length > 0 && saveAsCopy) || !saveAsCopy

	async function chooseDir() {
		const path = await OpenDir()
		safeTo = path
	}
</script>

<div
	class="collapse collapse-arrow border border-base-300 bg-base-100 rounded-box flex-shrink-0">
	<input type="checkbox" bind:checked={$editorOptions.compressionOpen} />
	<div class="collapse-title text-xl font-medium flex items-center gap-2">
		<div class="i-mdi-zip-box" />
		<span> Compression</span>
	</div>

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

			<div class="flex gap-4 items-center">
				<div class="flex flex-col gap-2 flex-shrink-0">
					<span>Save as Copy</span>
					<input
						type="checkbox"
						class="toggle toggle-success"
						bind:checked={saveAsCopy} />
				</div>

				<input
					type="text"
					readonly
					on:click={chooseDir}
					class="input input-success visible w-full"
					placeholder="Choose a directory"
					class:invisible={!saveAsCopy}
					bind:value={safeTo} />
			</div>

			<button class="btn" disabled={!submittable}> Compress </button>
		</section>
	</div>
</div>
