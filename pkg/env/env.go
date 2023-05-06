package env

import mongo "go.mongodb.org/mongo-driver/mongo"

var MongoDBConnection *mongo.Database

const RESTPort = "REST_PORT"

var REST_Port string

const MongoURI = "MongoURI"

var Mongo_URI string

var SigningKey string

var Encrypt_Key string

const CONFIRMED = "CONFIRMED"
const APPROVED = "APPROVED"
const REJECTED = "REJECTED"
const BaseURL = "http://127.0.0.1:8082/"
