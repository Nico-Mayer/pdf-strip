<script lang="ts">
	import Breadcrumbs from './Breadcrumbs.svelte'
	import { currentFile } from '$stores/store'

	import Encryption from '$lib/Editor/Views/Encryption/Encryption.svelte'
</script>

<main class="px-4 flex-1 flex flex-col overflow-x-hidden">
	<Breadcrumbs />

	<div class="prose flex items-center gap-4">
		<div class="i-mdi-file-pdf icon" />
		<h1 class="text-3xl py-4">{$currentFile.name}</h1>
	</div>

	<Encryption />
</main>
