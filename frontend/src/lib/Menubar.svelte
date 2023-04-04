<script lang="ts">
	import {
		WindowMinimise,
		WindowToggleMaximise,
		Quit,
	} from '../../wailsjs/runtime/runtime'
	import { os } from '$stores/store'

	async function minimise() {
		await WindowMinimise()
	}
	async function toggleMaximise() {
		await WindowToggleMaximise()
	}
	async function quit() {
		await Quit()
	}
</script>

<main style="--wails-draggable:drag" class="flex justify-between max-h-10 h-10">
	<section class="p-2 w-80 bg-base-200" />

	{#if $os === 'windows'}
		<section class="flex" style="--wails-draggable:no-drag">
			<button class="menu-btn hover:bg-white/10" on:click={minimise}>
				<div class="i-mdi-window-minimize" />
			</button>
			<button
				class="menu-btn hover:bg-white/10"
				on:click={toggleMaximise}>
				<div class="i-mdi-window-maximize" />
			</button>
			<button
				class="menu-btn hover:bg-red-500 hover:text-white"
				on:click={quit}>
				<div class="i-mdi-window-close" />
			</button>
		</section>
	{/if}
</main>

<style>
	.menu-btn {
		@apply h-10 w-12 flex flex-shrink-0 justify-center items-center;
	}
</style>
