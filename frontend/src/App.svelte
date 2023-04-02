<script lang="ts">
	import FileInput from '$lib/FileInput.svelte'
	import Sidebar from '$lib/Sidebar/Sidebar.svelte'
	import Editor from '$lib/Editor/Editor.svelte'
	import { filePaths } from '$stores/store'
	import { Toaster } from 'svelte-french-toast'
</script>

<main class="h-screen w-screen overflow-hidden flex min-w-[900px]">
	<Sidebar />

	{#if $filePaths.size <= 0}
		<FileInput />
	{:else}
		<Editor />
	{/if}
	<Toaster />
</main>
