package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pbcommon "github.com/kic/media/pkg/proto/common"
	pbmedia "github.com/kic/media/pkg/proto/media"
)

func main() {
	conn, err := grpc.Dial("localhost:9001",  grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pbmedia.NewMediaStorageClient(conn)

	in := &pbmedia.CheckForFileRequest{
		FileInfo: &pbcommon.File{
			FileName:     "test",
			FileLocation: "test",
			Metadata: map[string]string{"test": "test"},
		},
	}
	res, err := client.CheckForFileByName(context.Background(), in)

	fmt.Printf("res: %v\nerr: %v", res, err)
}
