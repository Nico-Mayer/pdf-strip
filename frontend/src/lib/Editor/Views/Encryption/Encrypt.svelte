<script lang="ts">
	import { Encrypt } from '$go/App'
	import { fade } from 'svelte/transition'
	import { currentFile } from '$stores/store'

	let showEncryptErr = false
	let showEncryptSuccess = false
	let provideUserPW = false

	let newPW = ''
	let userPW = ''

	$: submittable = newPW.length > 0 && (!provideUserPW || userPW.length > 0)
	$: $currentFile, reset()

	function reset() {
		showEncryptErr = false
		showEncryptSuccess = false
		provideUserPW = false
		newPW = ''
		userPW = ''
	}

	async function handleEncrypt() {
		if (!submittable) return

		if (provideUserPW) {
			showEncryptErr = !(await Encrypt($currentFile.path, userPW, newPW))
		} else {
			showEncryptErr = !(await Encrypt($currentFile.path, newPW, newPW))
		}

		if (!showEncryptErr) showEncryptSuccess = true
		newPW = ''
		userPW = ''

		setTimeout(() => {
			showEncryptSuccess = false
			showEncryptErr = false
		}, 3000)
	}
</script>

<form
	class="flex gap-2 w-full flex-col"
	on:submit|preventDefault={handleEncrypt}>
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
</form>

{#if showEncryptErr}
	<p class="badge badge-error mt-2" transition:fade>
		Already encrypted, decrypt first!
	</p>
{/if}
{#if showEncryptSuccess}
	<p class="badge badge-success mt-2" transition:fade>
		Encrypted successfully!
	</p>
{/if}
