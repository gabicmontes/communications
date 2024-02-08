package slack

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"net/http"
	"os"
)

type Response struct {
	Message string `json:"message"`
}

type RequestBody struct {
	Color   string `json:"color"`
	Title   string `json:"title"`
	Message string `json:"message"`
	Channel string `json:"channel"`
}

func Send(w http.ResponseWriter, r *http.Request) {

	godotenv.Load()

	token := os.Getenv("SLACK_TOKEN")

	var response Response

	if token == "" {
		response.Message = "Slack token not found"
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

	client := slack.New(token, slack.OptionDebug(true))
	attachment := slack.Attachment{
		Color:   requestBody.Color,
		Pretext: requestBody.Title,
		Text:    requestBody.Message,
	}
	_, timestamp, err := client.PostMessage(
		requestBody.Channel,
		slack.MsgOptionAttachments(attachment),
	)

	if err != nil {
		response.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Message = fmt.Sprintf("Message sent at %s", timestamp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return
}
