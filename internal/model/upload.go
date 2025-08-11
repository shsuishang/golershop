package model

// 上传得文件信息
type FileInfo struct {
	FileName         string `json:"file_name"`
	FileSize         int64  `json:"file_size"`
	FileUrl          string `json:"file_url"`
	Url              string `json:"url"`
	Type             string `json:"type"`
	FileType         string `json:"file_type"`
	FilePath         string `json:"file_path"`
	MimeType         string `json:"mime_type" ` // 素材类型
	UserId           uint   `json:"user_id" `
	MaterialDuration string `json:"material_duration" `
}
