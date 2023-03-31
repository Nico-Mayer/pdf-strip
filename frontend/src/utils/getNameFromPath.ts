export function getNameFromPath(path: string): string {
	const separator = path.includes('/') ? '/' : '\\'
	const pathArray = path.split(separator)
	return pathArray[pathArray.length - 1]
}
