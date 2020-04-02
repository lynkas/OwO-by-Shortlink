package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func init()  {
	LoadConfig()
	ConnectDB()
	DBBasicData()
	rand.Seed(time.Now().UnixNano())

}

func router(r *gin.Engine){
		r.GET("/query/:fullpath/", getURL)
		r.POST("/new/",newRecord)
}

func main() {
	r := gin.Default()
	router(r)
	r.Run()
}