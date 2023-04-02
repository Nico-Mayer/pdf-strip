<script lang="ts">
	import { Decrypt } from '$go/App'
	import { currentFile } from '$stores/store'

	let pw = ''
	let asOwner = false
	let decryptErr = false

	$: submittable = pw.length > 0
	$: $currentFile, reset()

	function reset() {
		pw = ''
		asOwner = false
		decryptErr = false
	}

	async function handleDecrypt() {
		if (!submittable) return

		decryptErr = !(await Decrypt($currentFile.path, pw, asOwner))
		console.log(decryptErr)

		pw = ''
	}
</script>

<form
	class="flex gap-2 w-full flex-col"
	on:submit|preventDefault={handleDecrypt}>
	<div class="form-control">
		<div class="flex items-center gap-2">
			<div class="i-mdi-lock-open-variant-outline w-6 h-6 text-warning" />
			<h1 class="font-semibold">Decrypt:</h1>
		</div>

		<label class="label cursor-pointer">
			<span class="label-text">Unlock as Owner</span>
			<input bind:checked={asOwner} type="checkbox" class="checkbox" />
		</label>
	</div>

	<input
		type="password"
		bind:value={pw}
		placeholder="Password"
		class="input input-bordered w-full" />

	<button class="btn" class:btn-disabled={!submittable}>Decrypt</button>
</form>
