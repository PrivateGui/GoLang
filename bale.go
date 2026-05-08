package main
//simple script for bale messenger using tgbotapi library via Golang
import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	token := ""

	rand.Seed(time.Now().UnixNano())

	bot, err := tgbotapi.NewBotAPIWithAPIEndpoint(token, "https://tapi.bale.ai/bot%s/%s")
	if err != nil {
		log.Fatal(err)
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Text == "/start" {

			randomToken := rand.Intn(1000000)

			keyboard := map[string]interface{}{
				"inline_keyboard": [][]map[string]interface{}{
					{
						{
							"text": "📋 رونوشت توکن",
							"copy_text": map[string]string{
								"text": fmt.Sprintf("%d", randomToken),
							},
						},
					},
				},
			}

			keyboardJSON, _ := json.Marshal(keyboard)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "════ ✧ توکن اختصاصی شما ✧ ════")

			msg.ReplyMarkup = json.RawMessage(keyboardJSON)

			_, err := bot.Send(msg)
			if err != nil {
				log.Println("send error:", err)
			}
		}
	}
}
