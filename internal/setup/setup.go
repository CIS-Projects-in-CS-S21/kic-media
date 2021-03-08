package setup

import (
	"context"
	"github.com/kic/media/pkg/cloudstorage"
	"net"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/kic/media/internal/server"
	"github.com/kic/media/pkg/database"
	pbmedia "github.com/kic/media/pkg/proto/media"
)

func DBRepositorySetup(logger *zap.SugaredLogger, dbPrefix string) (database.Repository, *mongo.Client) {
	MongoURI := os.Getenv("MONGO_URI")
	IsProduction := os.Getenv("PRODUCTION") != ""

	ctx := context.Background()

	mongoCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	mongoClient, err := mongo.Connect(mongoCtx, options.Client().ApplyURI(MongoURI))
	if err != nil {
		logger.Fatalf("Couldn't connect to mongo: %v", err)
	}

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		logger.Fatalf("Couldn't ping mongo: %v", err)
	}

	var dbName string
	if IsProduction {
		dbName = dbPrefix + "-prod"
	} else {
		dbName = dbPrefix + "-test"
	}

	repository := database.NewMongoRepository(mongoClient, logger)
	repository.SetCollections(dbName)
	return repository, mongoClient
}

func GRPCSetup(logger *zap.SugaredLogger, db database.Repository) *grpc.Server {
	IsProduction := os.Getenv("PRODUCTION") != ""
	ListenAddress := ":" + os.Getenv("PORT")

	listener, err := net.Listen("tcp", ListenAddress)
	if err != nil {
		logger.Fatalf("Unable to listen on %v: %v", ListenAddress, err)
	}

	grpcServer := grpc.NewServer()

	var bucketName string
	if IsProduction {
		bucketName = "media-prod"
	} else {
		bucketName = "media-test"
	}
	cloudStorage, err := cloudstorage.NewGoogleStorageBucket(bucketName)

	if err != nil {
		logger.Fatalf("Unable to connect to cloud store: %v", err)
	}

	mediaService := server.NewMediaStorageServer(db, cloudStorage, logger)
	pbmedia.RegisterMediaStorageServer(grpcServer, mediaService)

	reflection.Register(grpcServer)

	go func() {
		defer listener.Close()
		if err := grpcServer.Serve(listener); err != nil {
			logger.Fatalf("Failed to serve: %v", err)
		}
	}()

	logger.Infof("Server started on %v", ListenAddress)

	return grpcServer
}
