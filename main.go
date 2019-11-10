package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var DefaultBotAPI string
var DefaultChatID string
var BotAPI string
var ChatID string
var Hostname string

const TelegramBotSendMessageUrl = "https://api.telegram.org/bot%s/sendMessage"

type Payload struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func init() {
	if EnvBotAPI, ok := os.LookupEnv("BOT_API"); ok {
		BotAPI = EnvBotAPI
	} else {
		BotAPI = DefaultBotAPI
	}
	if EnvChatID, ok := os.LookupEnv("CHAT_ID"); ok {
		ChatID = EnvChatID
	} else {
		ChatID = DefaultChatID
	}

	if BotAPI == "" || ChatID == "" {
		log.Fatalf("BotAPI or ChatID cannot be empty")
	}
	Hostname, _ = os.Hostname()
}

func main() {
	message := "Ping"

	// Override message if stdin has content
	input := os.Stdin
	fi, _ := input.Stat()
	if fi.Size() > 0 {
		stdin, err := ioutil.ReadAll(input)
		if err == nil {
			message = string(stdin)
		}
	}

	// Build the payload
	data := Payload{
		ChatID: ChatID,
		Text:   fmt.Sprintf("Host: %s, Msg: %s", Hostname, message),
	}
	payloadBytes, _ := json.Marshal(data)
	payloadReader := bytes.NewReader(payloadBytes)

	url := fmt.Sprintf(TelegramBotSendMessageUrl, BotAPI)
	req, _ := http.NewRequest("POST", url, payloadReader)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	if resp.StatusCode != 200 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		log.Fatalf("Request failed: %s", string(bodyBytes))
	}

	defer resp.Body.Close()
}
