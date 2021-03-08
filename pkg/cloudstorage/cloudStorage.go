package cloudstorage

import "bytes"

type CloudStorage interface {
	UploadFile(fileName string, bytes bytes.Buffer) error
	DownloadFile(fileName string) (bytes.Buffer, error)
}
