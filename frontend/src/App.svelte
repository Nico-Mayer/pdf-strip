<script lang="ts">
	import Navbar from './lib/Navbar.svelte'
	import FileInput from './lib/FileInput.svelte'
	/* import { Greet } from '../wailsjs/go/main/App.js'

	function greet(): void {
		Greet(name).then((result) => (resultText = result))
	} */
</script>

<main class="h-screen w-screen flex-col flex">
	<Navbar />
	<div class="flex justify-center items-center flex-1">
		<FileInput />
	</div>
</main>
