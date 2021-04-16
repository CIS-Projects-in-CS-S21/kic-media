/*
This is for running integration tests in a production like environment
*/

package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	pbcommon "github.com/kic/media/pkg/proto/common"
	pbmedia "github.com/kic/media/pkg/proto/media"
	pbusers "github.com/kic/media/pkg/proto/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func shouldCheckForFile(authCtx context.Context, client pbmedia.MediaStorageClient) {
	in := &pbmedia.CheckForFileRequest{
		FileInfo: &pbcommon.File{
			FileName:     "testerino",
			FileLocation: "testerino",
			Metadata:     map[string]string{"test": "test"},
		},
	}
	res, err := client.CheckForFileByName(authCtx, in)

	if err != nil {
		log.Fatalf("Failed to check for file with: %v", err)
	}

	if res.Exists != true {
		log.Fatalf("Got improper response in check for file")
	}

	log.Printf("shouldCheckForFile Success")
}

func shouldUploadFile(authCtx context.Context, client pbmedia.MediaStorageClient) {
	buffer, err := ioutil.ReadFile("Makefile")

	if err != nil {
		log.Fatal("cannot read file: ", err)
	}

	req := &pbmedia.UploadFileRequest{
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
		FileURI: string(buffer),
	}

	resp, err := client.UploadFile(authCtx, req)

	if err != nil {
		log.Fatal("cannot upload image: ", err)
	}

	if resp.BytesRead != 538 {
		log.Fatalf("Reported %v bytes but should be 538", resp.BytesRead)
	}

	log.Printf("shouldUploadFile Success")
}

func shouldDownloadFile(authCtx context.Context, client pbmedia.MediaStorageClient) {
	dlStream, err := client.DownloadFileByName(authCtx, &pbmedia.DownloadFileRequest{
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
		buff.Write([]byte(recv.GetChunk()))
	}

	fi, err := os.Create("test_data/Makefile")
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()
	writer := bufio.NewWriter(fi)

	n, err := writer.Write(buff.Bytes())

	writer.Flush()

	if n != 538 {
		log.Fatalf("Reported %v bytes but should be 538", n)
	}

	log.Printf("shouldDownloadFile Success")
}

func shouldUpdateFile(authCtx context.Context, client pbmedia.MediaStorageClient) {
	updateReq := &pbmedia.UpdateFilesWithMetadataRequest{
		Strictness: pbmedia.MetadataStrictness_CASUAL,
		UpdateFlag: pbmedia.UpdateFlag_OVERWRITE,
		FilterMetadata: map[string]string{
			"rsc": "3711",
		},
		DesiredMetadata: map[string]string{
			"rsc": "42",
		},
	}

	updateRes, err := client.UpdateFilesWithMetadata(authCtx, updateReq)

	if err != nil {
		log.Fatal("could not update file: ", err)
	}

	if updateRes.NumFilesUpdated != 1 {
		log.Fatalf("Got %v files updated by should be 1", updateRes.NumFilesUpdated)
	}

	log.Printf("shouldUpdateFile Success")
}

func shouldDeleteFile(authCtx context.Context, client pbmedia.MediaStorageClient) {
	deleteReq := &pbmedia.DeleteFilesWithMetaDataRequest{Metadata: map[string]string{
		"rsc": "42",
	},
		Strictness: pbmedia.MetadataStrictness_CASUAL,
	}

	deleteRes, err := client.DeleteFilesWithMetaData(authCtx, deleteReq)
	if err != nil {
		log.Fatal("could not delete file: ", err)
	}

	if deleteRes.Success != true {
		log.Fatal("Failed to delete file")
	}

	log.Printf("shouldDeleteFile Success")
}

func main() {
	conn, err := grpc.Dial("test.api.keeping-it-casual.com:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pbmedia.NewMediaStorageClient(conn)

	// User client for auth

	usersClient := pbusers.NewUsersClient(conn)

	// get JWT
	tokRes, err := usersClient.GetJWTToken(context.Background(), &pbusers.GetJWTTokenRequest{
		Username: "testuser",
		Password: "testpass",
	})

	// creating auth context
	md := metadata.Pairs("Authorization", fmt.Sprintf("Bearer %v", tokRes.Token))
	authCtx := metadata.NewOutgoingContext(context.Background(), md)

	shouldCheckForFile(authCtx, client)
	shouldUploadFile(authCtx, client)
	shouldDownloadFile(authCtx, client)
	shouldUpdateFile(authCtx, client)
	shouldDeleteFile(authCtx, client)

}
