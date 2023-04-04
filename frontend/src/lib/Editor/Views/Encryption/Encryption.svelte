<script lang="ts">
	import Encrypt from './Encrypt.svelte'
	import { editorOptions } from '$stores/store'
	import Decrypt from './Decrypt.svelte'

	export let isEncrypted: boolean = false
</script>

<div
	class="collapse collapse-arrow border border-base-300 bg-base-100 rounded-box flex-shrink-0">
	<input type="checkbox" bind:checked={$editorOptions.encryptionOpen} />
	<div class="collapse-title text-xl font-medium flex items-center gap-2">
		<div class="i-mdi-security" />
		<span>Encryption</span>
	</div>
	<div class="collapse-content">
		<Encrypt on:checkStatus {isEncrypted} />
		<div class="divider flex-shrink-0" />
		<Decrypt on:checkStatus {isEncrypted} />
	</div>
</div>
