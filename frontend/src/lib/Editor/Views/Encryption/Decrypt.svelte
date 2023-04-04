<script lang="ts">
	import { Decrypt } from '$go/App'
	import { currentFile } from '$stores/store'
	import { createEventDispatcher } from 'svelte'
	import { toast } from 'svelte-french-toast'

	const dispatch = createEventDispatcher()
	let pw = ''
	let asOwner = true

	$: submittable = pw.length > 0
	$: $currentFile, reset()

	function reset() {
		pw = ''
		asOwner = false
	}

	async function handleDecrypt() {
		if (!submittable) return
		let decryptErr = false

		try {
			await Decrypt($currentFile.path, pw, asOwner)
		} catch (err) {
			decryptErr = true
		}

		if (decryptErr) {
			toast.error('Provide valid password.', {
				position: 'bottom-right',
				style: 'background-color: hsl(var(--er)); color: hsl(var(--n));',
			})
		} else {
			toast.success('File decrypted!', {
				position: 'bottom-right',
				style: 'background-color: hsl(var(--su)); color: hsl(var(--n));',
			})
		}

		pw = ''
		dispatch('checkStatus')
	}
</script>

<div
	class="border border-base-300 bg-base-200 rounded-box flex-shrink-0 shadow mt-4">
	<div class="collapse-title text-xl font-medium flex items-center gap-2">
		<div class="i-mdi-lock-open-variant w-6 h-6 text-warning" />
		<h1 class="font-semibold">Decrypt:</h1>
	</div>

	<div class="px-4 py-1 text-sm">
		<span>To edit this file, you must first decrypt it.</span>
	</div>

	<div class="p-1">
		<form
			on:submit|preventDefault={handleDecrypt}
			class="rounded-box border-base-200 gap-2 flex flex-col border p-2">
			<input
				type="password"
				bind:value={pw}
				placeholder="Password"
				class="input input-bordered w-full" />

			<button class="btn" class:btn-disabled={!submittable}
				>Decrypt</button>
		</form>
	</div>
</div>
