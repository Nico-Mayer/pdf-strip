<script lang="ts">
	import Navbar from '$lib/Navbar.svelte'
	import FileInput from '$lib/FileInput.svelte'
	import Sidebar from '$lib/Sidebar/Sidebar.svelte'
	import Editor from '$lib/Editor/Editor.svelte'
	import { filePaths } from '$stores/store'
</script>

<main class="h-screen w-screen flex-col flex overflow-hidden">
	<Navbar />

	<section class="flex w-full h-full">
		<Sidebar />

		{#if !$filePaths || $filePaths.size === 0}
			<FileInput />
		{:else}
			<Editor />
		{/if}
	</section>
</main>
