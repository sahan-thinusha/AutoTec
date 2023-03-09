package main

import (
	"autotec/pkg/env"
	"autotec/pkg/rest_controller"
	"context"
	echo "github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"strings"
	"time"
)

func init() {
	if val, exist := os.LookupEnv(env.RESTPort); exist && !strings.EqualFold(val, "") {
		env.REST_Port = val
	} else {
		env.REST_Port = "8080"
	}

	if val, exist := os.LookupEnv(env.MongoURI); exist && !strings.EqualFold(val, "") {
		env.Mongo_URI = val
	} else {
		env.Mongo_URI = "mongodb+srv://root:sahan12345@clustertest.qegu11v.mongodb.net/?retryWrites=true&w=majority"
	}

	env.Encrypt_Key = "92AE31A79FEEB2A3"
	env.SigningKey = "JhbGciOiJub25lIiwidHlwIjoiS"
}

func mongoConnect() {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(env.Mongo_URI).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	env.MongoDBConnection = client.Database("Auto_Tec")
}

func main() {

	mongoConnect()
	e := echo.New()
	rest_controller.EchoController(e)
	e.Logger.Fatal(e.Start(":" + env.REST_Port))
}
