<script lang="ts">
	import { GetPDFInfo } from '$go/App'
	import Compression from '$lib/Editor/Views/Compression/Compression.svelte'
	import Encryption from '$lib/Editor/Views/Encryption/Encryption.svelte'
	import { currentFile } from '$stores/store'
	import type { PdfInfo } from '$types'
	import Breadcrumbs from './Breadcrumbs.svelte'
	import Decrypt from '$lib/Editor/Views/Encryption/Decrypt.svelte'
	import Info from './Info.svelte'

	let info: PdfInfo
	let container: HTMLElement

	$: $currentFile, handelFileChange()
	$: encrypted = info?.Encrypted
	$: exists = info?.Exists

	async function handelFileChange() {
		try {
			const newInfo = await GetPDFInfo($currentFile.path)
			info = { ...newInfo }
		} catch (err) {
			info = null
		}
		if (container) container.scrollTop = 0
	}
</script>

<main
	class="px-4 flex-1 flex flex-col overflow-x-hidden overflow-y-auto pb-4 hide-scrollbar xl:max-w-7xl"
	bind:this={container}>
	<Breadcrumbs />
	<section class="flex items-cente gap-4">
		<div class="prose flex items-center gap-4">
			<div class="indicator">
				{#if encrypted}
					<div
						class="i-mdi-shield-lock-outline bg-warning h-4 w-4 indicator-item" />
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

	{#if info && exists && !encrypted}
		<Info {info} />
		<section class="flex flex-col gap-4">
			<Compression />
			<Encryption on:checkStatus={handelFileChange} />
		</section>
	{:else if exists && encrypted}
		<Decrypt on:checkStatus={handelFileChange} />
	{:else}
		<section class="alert alert-error shadow-lg">
			<div>
				<svg
					xmlns="http://www.w3.org/2000/svg"
					class="stroke-current flex-shrink-0 h-6 w-6"
					fill="none"
					viewBox="0 0 24 24"
					><path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
				<span>File does not exist</span>
				<br />
			</div>
			<div>
				<span class="badge">{$currentFile.path}</span>
			</div>
		</section>
	{/if}
</main>
