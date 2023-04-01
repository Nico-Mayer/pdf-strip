const { iconsPlugin, getIconCollections } = require('@egoist/tailwindcss-icons')
/** @type {import('tailwindcss').Config} */
export default {
	content: ['./index.html', './src/**/*.{js,ts,jsx,tsx,svelte}'],
	theme: {
		extend: {},
	},
	plugins: [
		require('daisyui'),
		iconsPlugin({
			// Select the icon collections you want to use
			collections: getIconCollections(['mdi']),
		}),
		require('@tailwindcss/typography'),
	],
}
