export interface PdfInfo {
	Exists: boolean
	Encrypted: boolean
	Size: number
	Info: Record<string, string>
}
