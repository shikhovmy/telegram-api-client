package TelegramApiClient

import "encoding/json"

type TelegramResponse struct {
	Ok          bool            `json:"ok"`
	Description string          `json:"description"`
	ErrorCode   int             `json:"error_code"`
	Results     json.RawMessage `json:"result"`
}

type Update struct {
	UpdateId          int               `json:"update_id"`
	Message           Message           `json:"message,omitempty"`
	EditedMessage     Message           `json:"edited_message,omitempty"`
	ChannelPost       Message           `json:"channel_post,omitempty"`
	EditedChannelPost Message           `json:"edited_channel_post,omitempty"`
	InlineQuery       InlineQuery       `json:"inline_query,omitempty"`
	ChosenInlineQuery ChosenInlineQuery `json:"chosen_inline_result,omitempty"`
}

type User struct {
	Id        int
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
}

type Chat struct {
	Id        int    `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	UserName  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Message struct {
	Id              int    `json:"message_id"`
	From            User   `json:"from"`
	Date            int    `json:"date"`
	Chat            Chat   `json:"chat"`
	ForwardFrom     User   `json:"forward_from,omitempty"`
	ForwardFromChat string `json:"forward_from_chat,omitempty"`
	ForwardDate     int    `json:"forward_date,omitempty"`
	//ReplyToMessage        Message `json:"reply_to_message,omitempty"`
	Text                  string          `json:"text,omitempty"`
	Entities              []MessageEntity `json:"entities,omitempty"`
	NewChatMember         User            `json:"new_chat_member,omitempty"`
	LeftChatMember        User            `json:"left_chat_member,omitempty"`
	NewChatMTitle         User            `json:"new_chat_title,omitempty"`
	DeleteChatPhoto       bool            `json:"delete_chat_photo,omitempty"`
	GroupChatCreated      bool            `json:"group_chat_created,omitempty"`
	SuperGroupChatCreated bool            `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated    bool            `json:"channel_chat_created,omitempty"`
	MigrateToChatId       int             `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatId     int             `json:"migrate_from_chat_id,omitempty"`
	//PinnedMessage         Message `json:"pinned_message,omitempty"`
}

type MessageEntity struct {
	Type   MessageEntityType `json:"type"`
	Offset int               `json:"offset"`
	Length int               `json:"length"`
	Url    int               `json:"url,omitempty"`
}

type ChosenInlineQuery struct {
	ResultId        string `json:"update_id"`
	From            User
	InlineMessageId int `json:"inline_message_id"`
	Query           string
}

type InlineQueryResult interface{}

type InlineQueryResultArticle struct {
	Type                string              `json:"type"`
	Id                  string              `json:"id"`
	Title               string              `json:"title"`
	InputMessageContent InputMessageContent `json:"input_message_content"`
	Url                 string              `json:"url,omitempty"`
	HideUrl             bool                `json:"hide_url,omitempty"`
	Description         string              `json:"description,omitempty"`
	ThumbUrl            string              `json:"thumb_url,omitempty"`
	ThumbWidth          string              `json:"thumb_width,omitempty"`
	ThumbHeight         string              `json:"thumb_height,omitempty"`
}

type InputMessageContent interface{}

type InputTextMessageContent struct {
	MessageText           string `json:"message_text"`
	ParseMode             string `json:"parse_mode,omitempty"`
	DisableWebPagePreview string `json:"disable_web_page_preview,omitempty"`
}

//Request objects

type InlineQuery struct {
	Id     string
	From   User
	Query  string
	Offset string
}

type InlineQueryAnswer struct {
	InlineQueryId string            `json:"inline_query_id"`
	Results       InlineQueryResult `json:"results"`
	CacheTime     int               `json:"cache_time,omitempty"`
	Personal      bool              `json:"is_personal,omitempty"`
	NextOffset    string            `json:"next_offset,omitempty"`
}

type TextMessage struct {
	ChatId                string      `json:"chat_id"`
	Text                  string      `json:"text"`
	ParseMode             string      `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
	DisableNotification   bool        `json:"disable_notification,omitempty"`
	ReplyToMessageId      int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           interface{} `json:"reply_markup,omitempty"`
}

type ChatAction struct {
	ChatId string         `json:"chat_id"`
	Action ChatActionType `json:"action"`
}

type ChatActionType string
type MessageEntityType string

const (
	TYPING_CHAT_ACTION          = ChatActionType("typing")
	upload_photo_CHAT_ACTION    = ChatActionType("upload_photo")
	record_video_CHAT_ACTION    = ChatActionType("record_video")
	upload_video_CHAT_ACTION    = ChatActionType("upload_video")
	record_audio_CHAT_ACTION    = ChatActionType("record_audio")
	upload_audio_CHAT_ACTION    = ChatActionType("upload_audio")
	upload_document_CHAT_ACTION = ChatActionType("upload_document")
	find_location_CHAT_ACTION   = ChatActionType("find_location")

	BOT_COMMAND_MESSAGE_TYPE = MessageEntityType("bot_command")
)
