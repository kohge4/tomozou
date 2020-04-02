package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tomozou/adapter/chatdata"
	"tomozou/domain"
	"tomozou/usecase"

	"github.com/gin-gonic/gin"
)

type ChatApplicationImpl struct {
	UseCase usecase.ChatApplication
}

func (ch *ChatApplicationImpl) ChatList(c *gin.Context) {
	id, _ := c.Get("tomozou-id")
	userID, _ := id.(int)
	println(userID)
}

func (ch *ChatApplicationImpl) UserChat(c *gin.Context) {
	var jsonBody interface{}
	c.BindJSON(&jsonBody)

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

	cR := []domain.UserChat{*chat}
	response, _ := chatResponse(cR)
	fmt.Println(response)
	c.JSON(200, response)
}

func (ch *ChatApplicationImpl) DisplayChatRoom(c *gin.Context) {
	//userID, _ := c.Get("tomozou-id")
	chatRooms, err := ch.UseCase.ChatRooms(1)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	c.JSON(200, chatRooms)
}

func (ch *ChatApplicationImpl) DisplayChatListByArtist(c *gin.Context) {
	artistID := c.Param("artistID")
	id, _ := strconv.Atoi(artistID)
	chatList, err := ch.UseCase.ChatListByArtistID(id)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	response, _ := chatResponse(chatList)
	c.JSON(200, response)
}

func chatResponse(chatList []domain.UserChat) ([]ChatResponse, error) {
	var chatResp []ChatResponse
	var chat ChatResponse
	for _, i := range chatList {
		chat = ChatResponse{
			ID:        i.ID,
			Comment:   i.Comment,
			Name:      "kokog",
			Image:     "https://i.scdn.co/image/ab6775700000ee85b2bd4f64bd8c250aedd13123",
			UserID:    i.UserID,
			ArtistID:  i.ArtistID,
			CreatedAt: i.CreatedAt,
		}
		chatResp = append(chatResp, chat)
	}
	return chatResp, nil
}
