import { writable } from 'svelte/store'
import type { Writable } from 'svelte/store'
import type { FileData } from '../types'

export const filePaths: Writable<string[]> = writable([])
export const currentFile: Writable<FileData> = writable({
	path: '',
	name: '',
	isDirectory: false,
	isFile: true,
} as FileData)
