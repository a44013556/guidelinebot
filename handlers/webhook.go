package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"guidelinebot/config"
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
	Events []Event `json:"events`
}

type Action struct {
	Type  string `json:"type"`
	Label string `json:"label"`
	Text  string `json:"text"`
}
type QuickReplyItem struct {
	Type   string `json:"type"`
	Action Action `json:"action"`
}

type QuickReply struct {
	Items []QuickReplyItem `json:"items"`
}
type Message struct {
	Type       string     `json:"type"`
	Text       string     `json:"text"`
	QuickReply *QuickReply `json:"quickReply,omitempty"`
}
type Payload struct {
	ReplyToken string    `json:"replyToken"`
	Messages   []Message `json:"messages"`
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
			switch event.Message.Text {
			case "查詢行程":
				replyReginOptions(event.ReplyToken)
			case "北海道", "東北", "關東", "中部", "關西", "中國", "四國", "九州", "沖繩":
				replyToCheckTourist(event.ReplyToken, fmt.Sprintf("這是 %s 的行程資訊！\n1. 景點A\n2. 景點B", event.Message.Text))
			default:
				replyToCheckTourist(event.ReplyToken, "請輸入「查詢行程」 或是 日本地域")
			}

		}
	}

	c.Status((http.StatusOK))
}

func replyReginOptions(replyToken string) {
	regions := []string{"北海道", "東北", "關東", "中部", "關西", "中國", "四國", "九州", "沖繩"}
	var items []QuickReplyItem
	for _, region := range regions {
		items = append(items, QuickReplyItem{
			Type: "action",
			Action: Action{
				Type:  "message",
				Label: region,
				Text:  region,
			},
		})
	}

	payload := Payload{
		ReplyToken: replyToken,
		Messages: []Message{
			{
				Type: "text",
				Text: "請選擇地區:",
				QuickReply: &QuickReply{
					Items: items,
				},
			},
		},
	}
	replyToLine(payload, "replyReginOptions")
}

func replyToCheckTourist(replyToken string, message string) {

	payload := Payload{
		ReplyToken: replyToken,
		Messages: []Message{
			{
				Type: "text",
				Text: message,
			},
		},
	}

	replyToLine(payload, "replyToCheckTourist")
}

func replyToLine(payload Payload, failfuncname string) {
	jsonBody, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", config.LineReplyEndpoint, bytes.NewBuffer(jsonBody))
	accessToken := os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(failfuncname, " Fail:", err)
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	log.Println("Line respone:", string(respBody))
}
