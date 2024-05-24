package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var router *gin.Engine
var token = "6706237172:AAFZyrXsYjMg2ion8MH2rG99Pf-Cjf-DjVw"

func init() {
	router = gin.Default()

	router.Any("/*path", func(c *gin.Context) {
		update := tgbotapi.Update{}
		err := c.BindJSON(&update)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 检查是否有新消息
		if update.Message != nil {
			// 处理用户发送的消息
			handleMessage(update.Message)
		}

		c.String(http.StatusOK, "OK")
	})
}

func handleMessage(message *tgbotapi.Message) {
	msg := message.Text
	userID := message.From.ID
	username := message.From.UserName
	msgID := message.MessageID
	// 根据用户发送的消息进行处理
	switch msg {
	case "/start":
		sendMessage(userID, "Hello, "+username+"! I'm your Telegram bot.")
	case "/help":
		sendMessage(userID, "How can I help you?")
	default:
		sendMessage(userID, "You said: "+msg, msgID)
	}
}

func sendMessage(userID int64, message string, replyMsgID ...int) {
	// 发送消息给用户
	bot, _ := tgbotapi.NewBotAPI(token)
	msg := tgbotapi.NewMessage(userID, message)
	if replyMsgID != nil {
		msg.ReplyToMessageID = replyMsgID[0]
	}
	bot.Send(msg)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
