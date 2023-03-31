<script lang="ts">
	import { onMount } from 'svelte'
	import { Metadata } from '../../../wailsjs/go/main/App'
	import Breadcrumbs from './Breadcrumbs.svelte'
	import { currentFile } from '../../stores/store'

	onMount(async () => {
		const metadata = await Metadata($currentFile.path)
		console.log(metadata)
	})
</script>

<main class="px-2">
	<Breadcrumbs />
</main>
