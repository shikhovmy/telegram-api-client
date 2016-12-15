package TelegramApiClient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type TelegramApi struct {
	client              HttpClientWrapper
	Messages            chan Message
	InlineUpdates       chan InlineQuery
	ChosenInlineUpdates chan ChosenInlineQuery
}

func New(token string) *TelegramApi {
	return &TelegramApi{
		client: &BasicHttpClientWrapper{
			baseUrl:      fmt.Sprintf("https://api.telegram.org/bot%s", token),
			postResponse: ParseResponse,
		},
		Messages:            make(chan Message, 50),
		InlineUpdates:       make(chan InlineQuery, 50),
		ChosenInlineUpdates: make(chan ChosenInlineQuery, 50),
	}
}

func (api *TelegramApi) UpdatesLoop() {
	defer close(api.Messages)

	fmt.Println("Telegram auto-update activited")

	offset := -1
	for {
		select {
		case <-time.After(500 * time.Millisecond):
			var updatesList []Update
			var err error
			if offset == -1 {
				updatesList, err = api.GetUpdates()
			} else {
				updatesList, err = api.GetSpecifiedUpdates(offset, 50, 0)
			}
			if err != nil {
				fmt.Println(err)
				continue
			}
			for _, update := range updatesList {
				offset = update.UpdateId + 1
				switch {
				case update.Message.Id != 0:
					api.Messages <- update.Message
				case update.ChannelPost.Id != 0:
					api.Messages <- update.ChannelPost
				case update.InlineQuery != (InlineQuery{}):
					api.InlineUpdates <- update.InlineQuery
				case update.ChosenInlineQuery != (ChosenInlineQuery{}):
					api.ChosenInlineUpdates <- update.ChosenInlineQuery
				}
			}
		}
	}
}

func (api *TelegramApi) GetMe() (user User, err error) {
	err = api.client.Get(getMe, nil, &user)
	return
}

func (api *TelegramApi) GetSpecifiedUpdates(offset int, limit int, timeout int) (update []Update, err error) {
	err = api.client.Get(getUpdates, map[string]string{
		"offset":  fmt.Sprint(offset),
		"limit":   fmt.Sprint(offset),
		"timeout": fmt.Sprint(timeout),
	}, &update)
	return
}

func (api *TelegramApi) GetUpdates() (update []Update, err error) {
	err = api.client.Get(getUpdates, nil, &update)
	return
}

func (api *TelegramApi) SendMessage(textMessage TextMessage) (message Message, err error) {
	err = api.client.Post(sendMessage, textMessage, &message)
	return
}

func (api *TelegramApi) SendChatAction(action ChatAction) (err error) {
	var response interface{}
	err = api.client.Post(sendChatAction, action, &response)
	return
}

func (api *TelegramApi) AnswerInlineQuery(answer InlineQueryAnswer) (err error) {
	var response interface{}
	err = api.client.Post(answerInlineQuery, answer, &response)
	return
}

func ParseResponse(originalResponse *http.Response, objectToReturn interface{}) (err error) {

	rawBody, _ := ioutil.ReadAll(originalResponse.Body)

	if originalResponse.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprint(originalResponse.Status, string(rawBody)))
		return
	}
	var teleResponse TelegramResponse
	if err = json.Unmarshal(rawBody, &teleResponse); err != nil {
		return
	}
	if teleResponse.Ok {
		json.Unmarshal(teleResponse.Results, objectToReturn)
		return

	}
	fmt.Println(teleResponse.Description)
	return errors.New(teleResponse.Description)
}

type HttpClientWrapper interface {
	Get(endpoint Endpoint, queryParams map[string]string, objectToReturn interface{}) (err error)
	Post(endpoint Endpoint, data interface{}, objectToReturn interface{}) (err error)
}

type Endpoint string

const (
	getMe             = Endpoint("/getMe")
	getUpdates        = Endpoint("/getUpdates")
	sendMessage       = Endpoint("/sendMessage")
	sendChatAction    = Endpoint("/sendChatAction")
	answerInlineQuery = Endpoint("/answerInlineQuery")
)

type InternalUpdate struct {
	Message Message
	error   bool
}

func (upd *InternalUpdate) containsError() bool {
	return upd.error
}
