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
			// 获取用户发送的消息内容
			message := update.Message.Text
			// 获取用户的 ID
			userID := update.Message.From.ID
			// 获取用户的用户名
			username := update.Message.From.UserName

			// 处理用户发送的消息
			handleMessage(userID, username, message)
		}

		c.String(http.StatusOK, "OK")
	})
}

func handleMessage(userID int64, username string, message string) {
	// 根据用户发送的消息进行处理
	switch message {
	case "/start":
		sendMessage(userID, "Hello, "+username+"! I'm your Telegram bot.")
	case "/help":
		sendMessage(userID, "How can I help you?")
	default:
		sendMessage(userID, "You said: "+message)
	}
}

func sendMessage(userID int64, message string) {
	// 发送消息给用户
	bot, _ := tgbotapi.NewBotAPI(token)
	msg := tgbotapi.NewMessage(userID, message)
	bot.Send(msg)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
