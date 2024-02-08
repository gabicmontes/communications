package telegram

import (
	"net/http"
	"os"
	"encoding/json"
	"github.com/joho/godotenv"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type RequestBody struct {
	GroupID int64 `json:"group_id"`
	Message string `json:"message"`
}

func Send(w http.ResponseWriter, r *http.Request) {

	godotenv.Load()

	token := os.Getenv("TELEGRAM_BOT_TOKEN")

	var errorResponse ErrorResponse

	if token == "" {
		errorResponse.Message = "Telegram bot token not found"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
	

	requestBody := RequestBody{}
	
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		errorResponse.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		errorResponse.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	alertText := tgbotapi.NewMessage(requestBody.GroupID, requestBody.Message)
	response, err := bot.Send(alertText)

	if err != nil {
		errorResponse.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	log.Printf("Message sent to group %d with message id %d", response.Chat.ID, response.MessageID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(requestBody)
	return
}
