package main

import (
	"context"
	"time"

	"github.com/javiertlopez/numeral/atlas"
	"github.com/javiertlopez/numeral/cloud"
	"github.com/javiertlopez/numeral/controller"
	"github.com/javiertlopez/numeral/router"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database keeps the database name
const Database = "numeral"

const mongoTimeout = 15 * time.Second

// App holds the handler, and logger
type App struct {
	router *mux.Router
}

// AppConfig struct with configuration variables
type AppConfig struct {
	commit   string
	version  string
	mongoURI string
}

// New returns an App
func New(cfg AppConfig) App {
	// Set client options
	clientOptions := options.Client().ApplyURI(cfg.mongoURI)

	// Context with timeout for establish connection with Mongo Atlas
	ctx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
	defer cancel()

	// Connect to Mongo Atlas
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	db := client.Database(Database)

	// S3
	awsCfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	s3Client := s3.NewFromConfig(awsCfg)

	repository := atlas.New(db)

	storage := cloud.New("picasso-numeral", s3Client)

	controller := controller.New(cfg.commit, cfg.version, repository, storage)

	router := router.New(controller)

	return App{
		router: router,
	}
}
