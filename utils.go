package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"math/big"
	"math/rand"
	"time"
)

var Config struct{
	Url string `json:"db"`

	Token map[string]string `json:"token"`
	//token : from

}

var RandomCharSize = 5

type Path struct{
	Fullpath string `uri:"fullpath" binding:"required"`
}

type info struct {
	LastID int64 `json:"last_id"`
}

var Client *mongo.Client

var configFile = "config.json"

func LoadConfig()  {
	data, err := ioutil.ReadFile(configFile)
    if err != nil {
      panic(err)
    }
    err = json.Unmarshal(data, &Config)
    if err != nil {
        panic(err)
    }
}

func ConnectDB()  {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(Config.Url))
    if err != nil {
    	log.Fatal(err)
    }

    //test connection
    err=client.Ping(ctx,nil)
    if err != nil {
    	log.Fatal(err,", connect failed, check db link in config.json")
    }
    Client=client
}

var DB = "owo"
var COInfo = "info"
var COdata = "data"

func DBBasicData()  {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	count,err := Client.Database(DB).Collection(COdata).CountDocuments(ctx,bson.M{})
	if err!=nil {
		log.Fatal(err)
	}
	UnID = count+1
	defer cancel()
}


func DBSettingInit()  {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := Client.Database(DB).Collection(COInfo)
	_,err:=collection.InsertOne(ctx,info{LastID:0})
	if err!=nil {
		log.Fatal(err)
	}
}

func ErrorMessage(c *gin.Context,code int,message interface{})  {
	c.JSON(code,gin.H{"message":message})
	c.Abort()
}

func FullPathExtract(fullpath string) *IDGroup {
	if len(fullpath)<RandomCharSize+1 {
		return nil
	}

	TextSerial:= fullpath[RandomCharSize:]

	var ID big.Int
	_,ok:=ID.SetString(TextSerial,62)

	if !ok {
		return nil
	}

	return &IDGroup{
		TextSerial: TextSerial,
		Prefix:     fullpath[0:RandomCharSize],
		Serial:		ID.Int64(),
	}
}

func GetRandom5() string {
	rand.Seed(time.Now().Unix())
	return randSeq(5)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}
