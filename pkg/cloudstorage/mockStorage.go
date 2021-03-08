package cloudstorage

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"

	"go.uber.org/zap"
)

type MockCloudStorage struct {
	dataPath string

	logger *zap.SugaredLogger
}

func NewMockCloudStorage(dataPath string, logger *zap.SugaredLogger) *MockCloudStorage {
	return &MockCloudStorage{
		dataPath: dataPath,
		logger:   logger,
	}
}

func (s *MockCloudStorage) UploadFile(fileName string, bytes bytes.Buffer) error {
	// open output file
	fo, err := os.Create(s.dataPath + "/" + fileName)
	if err != nil {
		s.logger.Debugf("Err upload: %v", err)
		return err
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			s.logger.Debugf("Err close: %v", err)
			panic(err)
		}
	}()

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := bytes.Read(buf)
		if err != nil && err != io.EOF {
			s.logger.Debugf("Err upload: %v", err)
			return err
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := fo.Write(buf[:n]); err != nil {
			s.logger.Debugf("Err upload: %v", err)
			return err
		}
	}
	return nil
}

func (s *MockCloudStorage) DownloadFile(fileName string) (bytes.Buffer, error) {
	fi, err := os.Open(s.dataPath + "/" + fileName)
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	all, err := ioutil.ReadAll(fi)

	if err != nil {
		panic(err)
	}

	return *bytes.NewBuffer(all), nil
}
