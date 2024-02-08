package telegram

import (
	"net/http"
	"os"
	"encoding/json"
	"github.com/joho/godotenv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Response struct {
	Message string `json:"message"`
}

type RequestBody struct {
	GroupID int64 `json:"group_id"`
	Message string `json:"message"`
}

func Send(w http.ResponseWriter, r *http.Request) {

	godotenv.Load()

	token := os.Getenv("TELEGRAM_BOT_TOKEN")

	var response Response

	if token == "" {
		response.Message = "Telegram bot token not found"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	

	requestBody := RequestBody{}
	
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		response.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		response.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	alertText := tgbotapi.NewMessage(requestBody.GroupID, requestBody.Message)
	resp, err := bot.Send(alertText)

	if err != nil {
		response.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Message = "Message sent to group " + resp.Chat.Title

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return
}
