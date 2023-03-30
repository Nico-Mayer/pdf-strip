<script lang="ts">
	import Navbar from './lib/Navbar.svelte'
	import FileInput from './lib/FileInput.svelte'
	import Sidebar from './lib/Sidebar/Sidebar.svelte'
	import Editor from './lib/Editor.svelte'
	import { files, currentFile } from './stores/store'

	$: console.log($files)
	$: console.log($currentFile)
</script>

<main class="h-screen w-screen flex-col flex">
	<Navbar />

	<section class="flex w-full h-full">
		<Sidebar files={$files} />

		{#if !$files || $files.length === 0}
			<FileInput />
		{:else}
			<Editor />
		{/if}
	</section>
</main>
