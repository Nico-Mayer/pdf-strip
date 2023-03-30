import { writable } from 'svelte/store'
import type { Writable } from 'svelte/store'

export const files: Writable<File[]> = writable(null)
export const currentFile: Writable<File> = writable(null)
