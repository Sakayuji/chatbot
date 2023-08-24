package api

import (
	"github.com/Sakayuji/chatbot/chatgpt"
	"github.com/Sakayuji/chatbot/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// SendMessage to chatbot
func SendMessage(c *gin.Context) {
	uId := c.Query("user_id")
	userId, _ := strconv.ParseInt(uId, 10, 64)
	content := c.PostForm("content")

	// sentiment analysis using chatgpt
	tag := chatgpt.GptParse(content)
	temp := &db.Templates{}
	// Pick appropriate reply
	db.DB.Where("tag = ?", db.TagMp[tag]).Find(&temp)

	rest := db.DB.Create(&db.Messages{UserId: int32(userId), Content: content, Tag: db.TagMp[tag]})
	if rest.Error != nil {
		c.JSON(http.StatusOK, gin.H{"Error": rest.Error})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Message": "Succeed", "response": temp.Content, "tag": tag})
	}

}

// CreateTemplate for chatbot to reply customers
func CreateTemplate(c *gin.Context) {
	uId := c.Query("user_id")
	userId, _ := strconv.ParseInt(uId, 10, 64)
	content := c.PostForm("content")
	tagStr := c.PostForm("tag")
	tag, _ := strconv.ParseInt(tagStr, 10, 64)
	categoryStr := c.PostForm("category")
	category, _ := strconv.ParseInt(categoryStr, 10, 64)

	rest := db.DB.Create(&db.Templates{CreatorId: int32(userId), Content: content, Tag: int8(tag),
		Category: int8(category)})
	if rest.Error != nil {
		c.JSON(http.StatusOK, gin.H{"Error": rest.Error})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Message": "Succeed"})
	}

}

// GetMessage from history
func GetMessage(c *gin.Context) {
	uId := c.Query("user_id")
	userId, _ := strconv.ParseInt(uId, 10, 64)
	tagStr := c.Query("tag")
	tag, _ := strconv.ParseInt(tagStr, 10, 64)

	messages := []db.Messages{}
	var q *gorm.DB
	q = db.DB
	if userId != 0 {
		q = q.Where("user_id = ?", userId)
	}
	if tag != 0 {
		q = q.Where("tag = ?", tag)
	}
	rest := q.Find(&messages)

	if rest.Error != nil {
		c.JSON(http.StatusOK, gin.H{"Error": rest.Error})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Message": "Succeed", "data": messages})
	}

}

// GetTemplate which created
func GetTemplate(c *gin.Context) {
	uId := c.Query("user_id")
	userId, _ := strconv.ParseInt(uId, 10, 64)
	tagStr := c.Query("tag")
	tag, _ := strconv.ParseInt(tagStr, 10, 64)

	templates := []db.Templates{}
	var q *gorm.DB
	q = db.DB
	if userId != 0 {
		q = q.Where("creator_id = ?", userId)
	}
	if tag != 0 {
		q = q.Where("tag = ?", tag)
	}
	rest := q.Find(&templates)

	if rest.Error != nil {
		c.JSON(http.StatusOK, gin.H{"Error": rest.Error})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Message": "Succeed", "data": templates})
	}

}
