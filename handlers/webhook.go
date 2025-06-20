package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Event struct {
	Type       string `json:"type`
	ReplyToken string `json:"replyToken"`
	Source     struct {
		UserId string `json:"userId"`
		Type   string `json:"type"`
	} `json:"source"`
	Message struct {
		ID   string `json:"id"`
		Type string `json:"type"`
		Text string `json:text`
	} `json:message`
}

type WebhookRequest struct {
	Events []Event `json:"evnets`
}

func LineWebhookHandler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Red Request Fail:", err)
		c.Status((http.StatusInternalServerError))
		return
	}

	var req WebhookRequest
	if err := json.Unmarshal(body, &req); err != nil {
		log.Println("Invaild Json:", err)
		return
	}

	for _, event := range req.Events {
		if event.Type == "message" && event.Message.Type == "text" {
			replyToLine(event.ReplyToken, "You just said:"+event.Message.Text)
		}
	}

	c.Status((http.StatusOK))
}

func replyToLine(replyToken string, message string) {
	endpoint := "https://api.line.me/v2/bot/message/reply"
	accessToken := os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")

	payload := map[string]interface{}{
		"replayToken": replyToken,
		"message": []map[string]string{
			{
				"type": "text",
				"text": message,
			},
		},
	}

	jsonBody, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBody))

	req.Header.Set("Context-Type", "application/json")
	req.Header.Set("authorization", "bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Send Message Fail:", err)
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	log.Println("Line respone:", respBody)

}
