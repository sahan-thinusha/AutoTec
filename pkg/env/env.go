package env

import mongo "go.mongodb.org/mongo-driver/mongo"

var MongoDBConnection *mongo.Database

const RESTPort = "REST_PORT"

var REST_Port string

const MongoURI = "MongoURI"

var Mongo_URI string

var SigningKey string

var Encrypt_Key string
