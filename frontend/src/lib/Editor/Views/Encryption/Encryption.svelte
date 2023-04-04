<script lang="ts">
	import Encrypt from './Encrypt.svelte'
	import { editorOptions } from '$stores/store'
</script>

<div
	class="collapse collapse-arrow border border-base-300 bg-base-200 rounded-box flex-shrink-0 shadow">
	<input type="checkbox" bind:checked={$editorOptions.encryptionOpen} />
	<div class="collapse-title text-xl font-medium flex items-center gap-2">
		<div class="i-mdi-security" />
		<span>Encryption</span>
	</div>
	<div class="collapse-content">
		<Encrypt on:checkStatus />
	</div>
</div>
