import { writable } from 'svelte/store'
import type { Writable } from 'svelte/store'
import type { FileData } from '../types'

export const filePaths = createPathStore()

function createPathStore() {
	const { subscribe, set, update } = writable<Set<string>>(new Set([]))

	return {
		subscribe,
		add: (paths: string[]) => {
			update((prevPaths: Set<string>) => {
				const pathsSet: Set<string> = new Set(paths)
				const newPaths: Set<string> = new Set([
					...prevPaths,
					...pathsSet,
				])
				return newPaths
			})
		},
		remove: (path: string) => {
			update((prevPaths) => {
				if (prevPaths.has(path)) {
					const newPaths = prevPaths
					newPaths.delete(path)
					return newPaths
				}
			})
		},
		reset: () => set(new Set([])),
	}
}

export const currentFile: Writable<FileData> = writable({
	path: '',
	name: '',
	isDirectory: false,
	isFile: true,
} as FileData)
