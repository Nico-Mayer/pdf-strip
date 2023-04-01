import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [svelte()],
	resolve: {
		alias: {
			$lib: path.resolve('./src/lib'),
			$types: path.resolve('./src/types'),
			$utils: path.resolve('./src/utils'),
			$stores: path.resolve('./src/stores'),
			$go: path.resolve('./wailsjs/go/main'),
		},
	},
})
