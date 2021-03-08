package cloudstorage

import "bytes"

// CloudStorage -  provides and interface for cloud storage provider specific implementations
type CloudStorage interface {
	UploadFile(fileName string, bytes bytes.Buffer) error
	DownloadFile(fileName string) (bytes.Buffer, error)
}
