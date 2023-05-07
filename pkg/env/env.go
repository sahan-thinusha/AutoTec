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
const BaseURL = "http://165.22.214.115:8082/"
const Finished = "Finished"
