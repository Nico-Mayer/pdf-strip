<script lang="ts">
	import { onMount } from 'svelte'
	import { Metadata } from '../../../wailsjs/go/main/App'
	import Breadcrumbs from './Breadcrumbs.svelte'
	import { currentFile } from '../../stores/store'
	import type { FileInfo } from 'src/types/fileInfo'

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
			<span>{metadata?.CreationTime.LowDateTime}</span>
		</div>
	</section>
</main>
