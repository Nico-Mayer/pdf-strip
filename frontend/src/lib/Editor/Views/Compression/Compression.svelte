<script lang="ts">
	import { editorOptions } from '$stores/store'
	import { OpenDir, Compress } from '$go/App'
	import { currentFile, os } from '$stores/store'
	import { createEventDispatcher } from 'svelte'
	import { toast } from 'svelte-french-toast'

	const dispatch = createEventDispatcher()
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
	async function handleCompress() {
		let separator = $os === 'windows' ? '\\' : '/'
		try {
			let outPath = $currentFile.path

			if (saveAsCopy) {
				const newName = $currentFile.name.replace(
					'.pdf',
					'_compressed.pdf'
				)
				outPath = `${safeTo}${separator}${newName}`
			}

			await Compress($currentFile.path, outPath, 3)
			dispatch('checkStatus')
			toast.success('Compression successful')
		} catch (error) {
			console.log(error)
			toast.error('Compression failed')
		}
	}
</script>

<div
	class="collapse collapse-arrow border rounded-box flex-shrink-0 bg-base-200 border-base-300 shadow">
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

			<button
				class="btn"
				disabled={!submittable}
				on:click={handleCompress}>
				Compress
			</button>
		</section>
	</div>
</div>
