<script lang="ts">
	import { Encrypt } from '$go/App'
	import { slide } from 'svelte/transition'
	import { currentFile } from '$stores/store'
	import toast from 'svelte-french-toast'
	import { createEventDispatcher } from 'svelte'

	const dispatch = createEventDispatcher()
	let provideUserPW = false
	let newPW = ''
	let confirmPW = ''
	let userPW = ''
	let confirmUserPW = ''

	$: submittable =
		newPW.length > 0 &&
		(!provideUserPW || userPW.length > 0) &&
		newPW === confirmPW &&
		(!provideUserPW || userPW === confirmUserPW)
	$: $currentFile, reset()

	function reset() {
		provideUserPW = false
		newPW = ''
		userPW = ''
	}

	async function handleEncrypt() {
		if (!submittable) return
		let showEncryptErr = false
		if (!provideUserPW) userPW = newPW

		try {
			await Encrypt($currentFile.path, userPW, newPW)
		} catch (err) {
			showEncryptErr = true
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
	<fieldset class="flex gap-2 w-full flex-col">
		<div class="form-control">
			<div class="flex items-center gap-2">
				<div class="i-mdi-lock w-6 h-6 text-success" />
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
			<div transition:slide class="flex flex-col gap-2">
				<input
					bind:value={userPW}
					type="password"
					placeholder="User Password"
					class="input input-bordered w-full" />
				<input
					bind:value={confirmUserPW}
					type="password"
					placeholder="Confirm User Password"
					class="input input-bordered w-full" />

				<div class="divider my-1" />
			</div>
		{/if}
		<input
			type="password"
			bind:value={newPW}
			placeholder="Password"
			class="input input-bordered w-full" />
		<input
			type="password"
			bind:value={confirmPW}
			placeholder="Confirm Password"
			class="input input-bordered w-full" />

		<button class="btn" class:btn-disabled={!submittable}>Encrypt</button>
	</fieldset>
</form>
