import { writable } from 'svelte/store'
import type { FileData } from '../types'
import { getNameFromPath } from '$utils'

export const filePaths = createPathStore()
export const currentFile = createCurrFileStore()

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

function createCurrFileStore() {
	const { subscribe, set, update } = writable<FileData>({
		path: '',
		name: '',
		isDirectory: false,
		isFile: true,
	} as FileData)

	return {
		subscribe,
		pick: (path: string) => {
			update(() => {
				const newFile: FileData = {
					path: path,
					name: getNameFromPath(path),
					isDirectory: false,
					isFile: true,
				}
				return newFile
			})
		},
		reset: () =>
			set({
				path: '',
				name: '',
				isDirectory: false,
				isFile: true,
			} as FileData),
	}
}
