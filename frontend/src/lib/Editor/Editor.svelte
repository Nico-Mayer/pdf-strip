<script lang="ts">
	import { onMount } from 'svelte'
	import { Metadata } from '$go/App'
	import Breadcrumbs from './Breadcrumbs.svelte'
	import { currentFile } from '$stores/store'
	import type { FileInfo } from '$types'

	let metadata: FileInfo[] = null

	onMount(async () => {
		metadata = await Metadata($currentFile.path)
		console.log(metadata)
	})
</script>

<main class="px-2">
	<Breadcrumbs />

	<section>
		<div>
			<span>GG</span>
		</div>
	</section>
</main>
