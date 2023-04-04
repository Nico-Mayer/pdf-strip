<script lang="ts">
	import type { PdfInfo } from '$types'

	export let info: PdfInfo
	let infoObj: Object = {}

	$: info, createInfoObj()
	$: console.log(infoObj)

	function createInfoObj() {
		if (!info.Info) return
		let temp = {}
		for (const item of info.Info) {
			const [key, value] = item.split(':').map((str) => str.trim())
			temp[key.replace(/\s+/g, '')] = value
		}
		infoObj = temp
	}
</script>

<div
	class="stats stats-vertical lg:stats-horizontal shadow my-4 bg-base-200 flex-shrink-0 border border-base-300">
	<div class="stat">
		<div class="stat-title">PDF Version</div>
		<div class="stat-value" />
	</div>

	<div class="stat">
		<div class="stat-title" />
		<div class="stat-value" />
	</div>

	<div class="stat">
		<div class="stat-title">New Registers</div>
		<div class="stat-value">1,200</div>
		<div class="stat-desc">↘︎ 90 (14%)</div>
	</div>
</div>
