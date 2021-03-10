/*
This is for running integration tests in a production like environment
*/

package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	pbcommon "github.com/kic/media/pkg/proto/common"
	pbmedia "github.com/kic/media/pkg/proto/media"
)

func main() {
	conn, err := grpc.Dial("test.api.keeping-it-casual.com:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pbmedia.NewMediaStorageClient(conn)

	in := &pbmedia.CheckForFileRequest{
		FileInfo: &pbcommon.File{
			FileName:     "testerino",
			FileLocation: "testerino",
			Metadata:     map[string]string{"test": "test"},
		},
	}
	res, err := client.CheckForFileByName(context.Background(), in)

	fmt.Printf("res: %v\nerr: %v\n", res, err)

	file, err := os.Open("Makefile")
	if err != nil {
		log.Fatal("cannot open image file: ", err)
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := client.UploadFile(ctx)
	if err != nil {
		log.Fatal("cannot upload image: ", err)
	}

	req := &pbmedia.UploadFileRequest{
		Data: &pbmedia.UploadFileRequest_FileInfo{
			FileInfo: &pbcommon.File{
				FileName:     "Makefile",
				FileLocation: "test",
				Metadata: map[string]string{
					"rsc": "3711",
					"r":   "2138",
					"gri": "1908",
					"adg": "912",
				},
			},
		},
	}

	stream.Send(req)

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot read chunk to buffer: ", err)
		}

		req := &pbmedia.UploadFileRequest{
			Data: &pbmedia.UploadFileRequest_Chunk{
				Chunk: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			log.Fatal("cannot send chunk to server: ", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("cannot receive response: ", err)
	}

	log.Printf("image uploaded with id: %s, size: %d", resp.FileID, resp.BytesRead)

	dlStream, err := client.DownloadFileByName(context.Background(), &pbmedia.DownloadFileRequest{
		FileInfo: &pbcommon.File{
			FileName: "Makefile",
		},
	})

	if err != nil {
		log.Fatal("cannot receive response in download: ", err)
	}

	var buf []byte
	buff := bytes.NewBuffer(buf)

	for {
		recv, err := dlStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot receive response in download: ", err)
		}
		buff.Write(recv.GetChunk())
	}

	fi, err := os.Create("test_data/Makefile")
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()
	writer := bufio.NewWriter(fi)

	n, err := writer.Write(buff.Bytes())

	writer.Flush()

	log.Printf("wrote %v bytes", n)

}
