package main

import (
	"flag"
	"fmt"
	tele "github.com/shikhovmy/telegram-api-client"
	"sync"
)

var token = flag.String("token", "", "Telegram token")

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	api := tele.New(token)

	go api.UpdatesLoop()
	go pongMessage(api)

	wg.Wait()
}

func pongMessage(api *tele.TelegramApi) {
	defer wg.Done()
	for message := range api.Messages {
		fmt.Print(message.Chat.Id)
		api.SendMessage(tele.TextMessage{
			ChatId: fmt.Sprint(message.Chat.Id),
			Text:   message.Text,
		})
	}
}
