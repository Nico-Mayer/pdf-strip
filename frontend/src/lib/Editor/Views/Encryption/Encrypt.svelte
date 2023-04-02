<script lang="ts">
	import { Encrypt } from '$go/App'
	import { fade } from 'svelte/transition'
	import { currentFile } from '$stores/store'
	import toast from 'svelte-french-toast'
	import { createEventDispatcher } from 'svelte'

	export let isEncrypted: boolean = false

	const dispatch = createEventDispatcher()
	let provideUserPW = false
	let newPW = ''
	let userPW = ''

	$: submittable = newPW.length > 0 && (!provideUserPW || userPW.length > 0)
	$: $currentFile, reset()

	function reset() {
		provideUserPW = false
		newPW = ''
		userPW = ''
	}

	async function handleEncrypt() {
		if (!submittable) return
		let showEncryptErr = false

		if (provideUserPW) {
			showEncryptErr = !(await Encrypt($currentFile.path, userPW, newPW))
		} else {
			showEncryptErr = !(await Encrypt($currentFile.path, newPW, newPW))
		}

		if (showEncryptErr) {
			toast.error('Already encrypted, decrypt first.', {
				position: 'bottom-right',
				style: 'background-color: hsl(var(--er)); color: hsl(var(--n));',
			})
		} else {
			toast.success('File encrypted!', {
				position: 'bottom-right',
				style: 'background-color: hsl(var(--su)); color: hsl(var(--n));',
			})
		}

		newPW = ''
		userPW = ''
		dispatch('checkStatus')
	}
</script>

<form on:submit|preventDefault={handleEncrypt}>
	<fieldset
		class="flex gap-2 w-full flex-col"
		disabled={isEncrypted}
		class:opacity-30={isEncrypted}>
		<div class="form-control">
			<div class="flex items-center gap-2">
				<div class="i-mdi-lock-outline w-6 h-6 text-success" />
				<h1 class="font-semibold">Encrypt:</h1>
			</div>

			<label class="label cursor-pointer">
				<span class="label-text">Specify user password</span>
				<input
					bind:checked={provideUserPW}
					type="checkbox"
					class="checkbox" />
			</label>
		</div>
		{#if provideUserPW}
			<input
				transition:fade
				bind:value={userPW}
				type="password"
				name="ownerPW"
				placeholder="User password"
				class="input input-bordered w-full" />
		{/if}
		<input
			type="password"
			bind:value={newPW}
			name="newPW"
			placeholder="New password"
			class="input input-bordered w-full" />

		<button class="btn" class:btn-disabled={!submittable}>Encrypt</button>
	</fieldset>
</form>
