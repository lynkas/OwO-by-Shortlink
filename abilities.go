package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"math/big"
	"net/http"
	"time"
)

func getURL(c *gin.Context)  {
	var path Path
	err:=c.ShouldBindUri(&path)
	if err!=nil {
		ErrorMessage(c,http.StatusForbidden,err.Error())
		return
	}

	id := FullPathExtract(path.Fullpath)
	if id==nil {
		ErrorMessage(c,http.StatusNotFound,"not found")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := Client.Database(DB).Collection(COdata)
	result := collection.FindOne(ctx,bson.M{"serial":id.Serial,"prefix":id.Prefix})
	var record Record
	err=result.Decode(&record)
	if err!=nil {
		ErrorMessage(c,http.StatusNotFound,err.Error())
		return
	}

	if record.Serial==0 {
		ErrorMessage(c,http.StatusNotFound,"not found")
		return
	} else {
		c.JSON(http.StatusOK,record.toMessageGroup())
		c.Abort()
		return
	}
}

func newRecord(c *gin.Context){
	var newRecord NewRecord
	err:=c.ShouldBindJSON(&newRecord)
	if err!=nil {
		ErrorMessage(c,http.StatusForbidden,err.Error())
		return
	}
	from,ok:=Config.Token[newRecord.Token]
	if !ok {
		ErrorMessage(c,http.StatusForbidden,"not ok")
		return
	}
	id := UnID
	UnID+=1
	record:=Record{
		From:     from,
		Identify: newRecord.Identify,
		Format: newRecord.Format,
		Content:newRecord.Content,
		TextSerial: big.NewInt(id).Text(62),
		Prefix:     GetRandom5(),
		Serial:     id,

	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := Client.Database(DB).Collection(COdata)
	_,err =collection.InsertOne(ctx,record)
	if err!=nil {
		ErrorMessage(c,http.StatusForbidden,err.Error())
		return
	}
	c.JSON(http.StatusOK,record.toIDGroup())
	return
}