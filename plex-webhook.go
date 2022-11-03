package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"plex-webhook/action"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/plex", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Yes , This is a plex wobhook server.",
		})
	})
	r.POST("/1bb7537b450781daa736cab535819c57/plex", plexWebhook) // random string path for public
	r.Run()                                                       // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func plexWebhook(c *gin.Context) {
	// body, _ := ioutil.ReadAll(c.Request.Body)
	// println(string(body))
	// c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

	//mf, exist := c.GetPostFormMap("payload")
	mf, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	//fmt.Println(mf.Value)
	p := mf.Value["payload"][0]
	// fmt.Println(p)

	var r Payload
	//if err := mapstructure.Decode(p, &r); err != nil {
	if err := json.Unmarshal([]byte(p), &r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	var name string
	if r.Metadata.LibrarySectionType == LibrarySectionShow {
		name = r.Metadata.GrandparentTitle + " " + r.Metadata.ParentTitle + " " + r.Metadata.Title
	} else {
		name = r.Metadata.Title
	}
	if r.Event == EventTypePlay || r.Event == EventTypeResume {
		// fmt.Println("fuck")
		msg := fmt.Sprintf("用户：%s 正在 %s 通过 %s 播放 %s %s ", r.Account.Title, r.Player.PublicAddress, r.Player.Title, r.Metadata.LibrarySectionTitle, name)
		// fmt.Println(msg)
		err := action.PushTgNotification(msg)
		if err != nil {
			log.Panicf("Send Notifaction failed with %s .\n", err)
		}
		log.Println(msg)
	}
	c.JSON(http.StatusOK, gin.H{"message": "got it"})
}
