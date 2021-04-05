/*
This is for running integration tests in a production like environment
*/

package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"io/ioutil"
	"log"
	"os"

	pbcommon "github.com/kic/media/pkg/proto/common"
	pbmedia "github.com/kic/media/pkg/proto/media"
	pbusers "github.com/kic/media/pkg/proto/users"
)

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

	// -----------------

	in := &pbmedia.CheckForFileRequest{
		FileInfo: &pbcommon.File{
			FileName:     "testerino",
			FileLocation: "testerino",
			Metadata:     map[string]string{"test": "test"},
		},
	}
	res, err := client.CheckForFileByName(authCtx, in)

	fmt.Printf("res: %v\nerr: %v\n", res, err)


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
		File: buffer,
	}


	resp, err := client.UploadFile(authCtx, req)

	if err != nil {
		log.Fatal("cannot upload image: ", err)
	}


	log.Printf("image uploaded with id: %s, size: %d", resp.FileID, resp.BytesRead)

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

	// Updating files --------------------

	// creating update request
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

	fmt.Printf("Update res: %v\nUpdate err: %v\n", updateRes, err)
	// ------------------------------------

	// Deleting files -----------------------

	deleteReq := &pbmedia.DeleteFilesWithMetaDataRequest{Metadata: map[string]string{
		"rsc": "42",
	},
	Strictness: pbmedia.MetadataStrictness_CASUAL,
	}

	deleteRes, err := client.DeleteFilesWithMetaData(authCtx, deleteReq)
	if err != nil {
		log.Fatal("could not delete file: ", err)
	}

	log.Printf("Delete res: %v\nDelete err: %v\n", deleteRes, err)

	// --------------------------------------



}
