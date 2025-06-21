package linebot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"guidelinebot/config"
	"guidelinebot/models"
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
	Type       string      `json:"type"`
	Text       string      `json:"text"`
	QuickReply *QuickReply `json:"quickReply,omitempty"`
}
type Payload struct {
	ReplyToken string    `json:"replyToken"`
	Messages   []Message `json:"messages"`
}

func (h *LineBotHandler) LineWebhookHandler(c *gin.Context) {
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
		userId := event.Source.UserId
		log.Println(userId)
		replyToken := event.ReplyToken
		if event.Type == "message" && event.Message.Type == "text" {
			text := event.Message.Text
			switch text {
			case "查詢行程":
				h.replyReginOptions(replyToken)
			default:
				area, err := models.CheckJapanAreaExists(h.DB, text)
				if err != nil {
					log.Println("DB error:", err)
					replyToCheckTourist(replyToken, "系統錯誤, 請稍後再試")
					return
				}

				if area != nil {
					h.replyTheTouristSpot(replyToken, *area)
				} else {
					replyToCheckTourist(replyToken, "請輸入「查詢行程」 或是 日本地域")
				}
			}

		}
	}

	c.Status((http.StatusOK))
}

func (h *LineBotHandler) replyReginOptions(replyToken string) {
	regions, err := models.GetAllJapanAreaName(h.DB)
	if err != nil {
		log.Println("DB error:", err)
		replyToCheckTourist(replyToken, "系統錯誤, 請稍後再試")
		return
	}
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

func (h *LineBotHandler) replyTheTouristSpot(replyToken string, area models.JapanArea) {
	spots, err := models.GetAreaSpotListByAreaId(h.DB, int64(area.ID))
	if err != nil {
		log.Println("DB error:", err)
		replyToCheckTourist(replyToken, "系統錯誤, 請稍後再試")
		return
	}
	message := fmt.Sprintf("這是 %s 的行程資訊！\n", area.Name)
	for i, spot := range spots {
		message += fmt.Sprintf("%d.%s\n", i+1, spot.Name)
	}

	payload := Payload{
		ReplyToken: replyToken,
		Messages: []Message{
			{
				Type: "text",
				Text: message,
			},
		},
	}

	replyToLine(payload, "replyTheTouristSpot")
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
