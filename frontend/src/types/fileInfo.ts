interface TimeData {
	LowDateTime: number
	HighDateTime: number
}

export interface FileInfo {
	FileAttributes: number
	CreationTime: TimeData
	LastAccessTime: TimeData
	LastWriteTime: TimeData
	FileSizeHigh: number
	FileSizeLow: number
	Reserved0: number
}
