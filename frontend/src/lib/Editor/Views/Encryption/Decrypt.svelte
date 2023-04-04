<script lang="ts">
	import { Decrypt } from '$go/App'
	import { currentFile } from '$stores/store'
	import { createEventDispatcher } from 'svelte'
	import { toast } from 'svelte-french-toast'

	export let isEncrypted: boolean = false

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

		const decryptErr = !(await Decrypt($currentFile.path, pw, asOwner))
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

<form on:submit|preventDefault={handleDecrypt}>
	<fieldset
		disabled={!isEncrypted}
		class="gap-2 flex flex-col"
		class:opacity-30={!isEncrypted}>
		<div class="form-control">
			<div class="flex items-center gap-2">
				<div class="i-mdi-lock-open-variant w-6 h-6 text-warning" />
				<h1 class="font-semibold">Decrypt:</h1>
			</div>
		</div>

		<input
			type="password"
			bind:value={pw}
			placeholder="Password"
			class="input input-bordered w-full" />

		<button class="btn" class:btn-disabled={!submittable}>Decrypt</button>
	</fieldset>
</form>
