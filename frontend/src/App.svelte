<script lang="ts">
	import FileInput from '$lib/FileInput.svelte'
	import Sidebar from '$lib/Sidebar/Sidebar.svelte'
	import Editor from '$lib/Editor/Editor.svelte'
	import { filePaths } from '$stores/store'
	import { Toaster } from 'svelte-french-toast'
	import Menubar from '$lib/Menubar.svelte'
</script>

<main class="h-screen flex-col w-screen flex">
	<Menubar />
	<section class="flex flex-1 w-screen min-w-[900px] overflow-hidden">
		<Sidebar />

		<div class="flex flex-1 xl:justify-center">
			{#if $filePaths.size <= 0}
				<FileInput />
			{:else}
				<Editor />
			{/if}
		</div>
	</section>

	<Toaster />
</main>
