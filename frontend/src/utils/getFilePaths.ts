import { OpenFiles } from '../../wailsjs/go/main/App'

export async function getFilePaths(): Promise<string[]> {
	const res = await OpenFiles()
	if (res === null) return []
	return res
}
