package controller

import (
	"encoding/json"
	"net/http"
	"tomozou/adapter/chatdata"
	"tomozou/usecase"

	"github.com/gin-gonic/gin"
)

type ChatApplicationImpl struct {
	UseCase usecase.ChatApplication
}

func (ch *ChatApplicationImpl) UserChat(c *gin.Context) {
	var jsonBody interface{}
	c.BindJSON(&jsonBody)
	//log.Debug().Interface("Body", jsonBody).Msg("RequestBody")

	var chatIn chatdata.ChatIn

	jsonByte, _ := json.Marshal(jsonBody)
	err := json.Unmarshal(jsonByte, &chatIn)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	chat, err := chatIn.UserChat()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	ch.UseCase.UserComment(chat)

	c.JSON(200, chat)
}

func (ch *ChatApplicationImpl) DisplayChatRoom(c *gin.Context) {
	//userID, _ := c.Get("tomozou-id")
	chatRooms, err := ch.UseCase.ChatRooms(1)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	c.JSON(200, chatRooms)
}
