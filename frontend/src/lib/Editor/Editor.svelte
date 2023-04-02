<script lang="ts">
	import Breadcrumbs from './Breadcrumbs.svelte'
	import { currentFile } from '$stores/store'
	import { IsEncrypted } from '$go/App'
	import Compression from '$lib/Editor/Views/Compression/Compression.svelte'
	import Encryption from '$lib/Editor/Views/Encryption/Encryption.svelte'

	let isEncrypted = false

	$: $currentFile, checkIfEncrypted()

	async function checkIfEncrypted() {
		isEncrypted = await IsEncrypted($currentFile.path)
	}
</script>

<main
	class="px-4 flex-1 flex flex-col overflow-x-hidden overflow-y-auto pb-4 hide-scrollbar xl:max-w-7xl">
	<Breadcrumbs />

	<section class="flex items-cente gap-4 mb-2">
		<div class="prose flex items-center gap-4">
			<div class="indicator">
				{#if isEncrypted}
					<div
						class="i-mdi-shield-lock-outline bg-error h-4 w-4 indicator-item" />
				{:else}
					<div
						class="i-mdi-shield-lock-open-outline bg-success h-4 w-4 indicator-item" />
				{/if}
				<div class="i-mdi-file-pdf icon" />
			</div>

			<h1 class="text-3xl py-4">{$currentFile.name}</h1>
		</div>
		<div class="flex items-center" />
	</section>

	<Compression />
	<div class="divider flex-shrink-0" />
	<Encryption {isEncrypted} on:checkStatus={checkIfEncrypted} />
</main>
