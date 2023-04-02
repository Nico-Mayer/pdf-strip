<script lang="ts">
	import { Encrypt } from '$go/App'
	import Breadcrumbs from './Breadcrumbs.svelte'
	import { currentFile } from '$stores/store'
	import { fade } from 'svelte/transition'

	let showEncryptErr = false
	let showEncryptSuccess = false
	let showOwnerPW = false

	async function handleEncrypt(e: SubmitEvent) {
		const newPwInput = (e.target as HTMLFormElement).elements.namedItem(
			'newPW'
		) as HTMLInputElement
		const ownerPwInput = (e.target as HTMLFormElement).elements.namedItem(
			'ownerPW'
		) as HTMLInputElement

		if (!newPwInput.value) return

		const newPW = newPwInput.value
		const ownerPW = ownerPwInput ? ownerPwInput.value : newPW

		showEncryptErr = !(await Encrypt($currentFile.path, newPW, ownerPW))

		if (!showEncryptErr) showEncryptSuccess = true

		newPwInput.value = ''
		if (ownerPwInput) ownerPwInput.value = ''

		setTimeout(() => {
			showEncryptSuccess = false
			showEncryptErr = false
		}, 3000)
	}
</script>

<main class="px-4 flex-1 flex flex-col overflow-x-hidden">
	<Breadcrumbs />

	<div class="prose">
		<h1 class="text-3xl py-4">{$currentFile.name}</h1>
	</div>

	<section class="py-4">
		<form
			class="flex gap-2 w-full flex-col"
			on:submit|preventDefault={handleEncrypt}>
			<h2 class="font-bold text-lg">Encrypt:</h2>
			<div class="form-control">
				<label class="label cursor-pointer">
					<span class="label-text">Specify owner password</span>
					<input
						bind:checked={showOwnerPW}
						type="checkbox"
						class="checkbox checkbox-primary" />
				</label>
			</div>
			{#if showOwnerPW}
				<input
					transition:fade
					type="password"
					name="ownerPW"
					placeholder="Owner password"
					class="input input-bordered w-full" />
			{/if}
			<input
				type="password"
				name="newPW"
				placeholder="New password"
				class="input input-bordered w-full" />

			<button class="btn">Encrypt</button>
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
	</section>
</main>
