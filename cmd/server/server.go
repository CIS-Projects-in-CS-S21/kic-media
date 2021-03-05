package main

import (
	"context"
	"go.uber.org/zap"
	"net"
	"os"
	"os/signal"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"

	"github.com/kic/media/internal/server"
	"github.com/kic/media/pkg/database"
	"github.com/kic/media/pkg/logging"
	pbmedia "github.com/kic/media/pkg/proto/media"
	"google.golang.org/grpc/reflection"

)

func dbRepositorySetup(logger *zap.SugaredLogger) (database.Repository, *mongo.Client) {
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
		dbName = "media-prod"
	} else {
		dbName = "media-test"
	}

	repository := database.NewMongoRepository(mongoClient, logger)
	repository.SetCollections(dbName)
	return repository, mongoClient
}

func grpcSetup(logger *zap.SugaredLogger, db database.Repository) *grpc.Server {
	ListenAddress := ":" + os.Getenv("PORT")

	listener, err := net.Listen("tcp", ListenAddress)
	if err != nil {
		logger.Fatalf("Unable to listen on %v: %v", ListenAddress, err)
	}

	grpcServer := grpc.NewServer()

	mediaService := server.NewMediaStorageServer(db, logger)
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

func main() {
	logger := logging.CreateLogger(zapcore.DebugLevel)

	repo, mongoClient := dbRepositorySetup(logger)

	serv := grpcSetup(logger, repo)

	defer serv.Stop()
	defer mongoClient.Disconnect(context.Background())

	// the server is listening in a goroutine so hang until we get an interrupt signal
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c
}
