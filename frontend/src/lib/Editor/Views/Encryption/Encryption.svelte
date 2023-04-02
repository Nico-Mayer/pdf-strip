<script lang="ts">
	import Encrypt from './Encrypt.svelte'
	import { editorOptions } from '$stores/store'
	import Decrypt from './Decrypt.svelte'
</script>

<div
	class="collapse collapse-arrow border border-base-300 bg-base-100 rounded-box mt-2">
	<input type="checkbox" bind:checked={$editorOptions.encryptionOpen} />
	<div class="collapse-title text-xl font-medium">Encryption</div>
	<div class="collapse-content">
		<Encrypt />
		<div class="divider" />
		<Decrypt />
	</div>
</div>
