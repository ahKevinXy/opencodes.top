package upload

import "mime/multipart"

type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

func NewOss(uploads string) OSS {
	switch uploads {
	case "local": //本地
		return &Local{}
	}
	return nil
}
