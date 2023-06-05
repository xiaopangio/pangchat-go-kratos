package service_logic

type DownloadFileRequest struct {
	FilePath string `json:"file_path" form:"file_path" binding:"required"`
}
