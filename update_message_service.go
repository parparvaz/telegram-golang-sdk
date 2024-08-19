package telegram

import (
	"context"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

type EditMessageReplyMarkupService struct {
	c               *Client
	chatID          *int64
	messageID       *int64
	inlineMessageID *string
	replyMarkup     *string
}

func (t *EditMessageReplyMarkupService) ChatID(chatID int64) *EditMessageReplyMarkupService {
	t.chatID = &chatID
	return t
}

func (t *EditMessageReplyMarkupService) MessageID(messageID int64) *EditMessageReplyMarkupService {
	t.messageID = &messageID
	return t
}

func (t *EditMessageReplyMarkupService) InlineMessageID(inlineMessageID string) *EditMessageReplyMarkupService {
	t.inlineMessageID = &inlineMessageID
	return t
}

func (t *EditMessageReplyMarkupService) ReplyMarkup(replyMarkup InlineKeyboardMarkup) *EditMessageReplyMarkupService {
	json, err := jsoniter.Marshal(&replyMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyMarkup = &jsonString
	return t
}

func (t *EditMessageReplyMarkupService) Do(ctx context.Context, opts ...RequestOption) (res []*EditMessageReplyMarkup, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/editMessageReplyMarkup",
	}

	if t.chatID != nil {
		r.setParam("chat_id", *t.chatID)
	}
	if t.messageID != nil {
		r.setParam("message_id", *t.messageID)
	}
	if t.inlineMessageID != nil {
		r.setParam("inline_message_id", *t.inlineMessageID)
	}
	if t.replyMarkup != nil {
		r.setParam("reply_markup", *t.replyMarkup)
	}
	data, err := t.c.callAPI(ctx, r, opts...)

	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*EditMessageReplyMarkup, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type EditMessageReplyMarkup struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}

type EditMessageTextService struct {
	c                     *Client
	chatID                *int64
	messageID             *int64
	inlineMessageID       *string
	text                  *string
	parseMode             *string
	entities              *string
	disableWebPagePreview *bool
	replyMarkup           *string
}

func (t *EditMessageTextService) Text(text string) *EditMessageTextService {
	t.text = &text
	return t
}

func (t *EditMessageTextService) ChatID(chatID int64) *EditMessageTextService {
	t.chatID = &chatID
	return t
}

func (t *EditMessageTextService) MessageID(messageID int64) *EditMessageTextService {
	t.messageID = &messageID
	return t
}

func (t *EditMessageTextService) InlineMessageID(inlineMessageID string) *EditMessageTextService {
	t.inlineMessageID = &inlineMessageID
	return t
}

func (t *EditMessageTextService) ParseMode(parseMode string) *EditMessageTextService {
	t.parseMode = &parseMode
	return t
}

func (t *EditMessageTextService) Entities(entities []MessageEntity) *EditMessageTextService {
	json, err := jsoniter.Marshal(&entities)
	if err != nil {
		return nil
	}

	jsonString := string(json)
	t.entities = &jsonString

	return t
}

func (t *EditMessageTextService) DisableWebPagePreview(disableWebPagePreview bool) *EditMessageTextService {
	t.disableWebPagePreview = &disableWebPagePreview
	return t
}

func (t *EditMessageTextService) ReplyMarkup(replyMarkup InlineKeyboardMarkup) *EditMessageTextService {
	json, err := jsoniter.Marshal(&replyMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)
	t.replyMarkup = &jsonString

	return t
}

func (t *EditMessageTextService) Do(ctx context.Context, opts ...RequestOption) (res []*EditMessageText, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/editMessageText",
	}

	r.setParam("text", *t.text)
	if t.chatID != nil {
		r.setParam("chat_id", *t.chatID)
	}
	if t.messageID != nil {
		r.setParam("message_id", *t.messageID)
	}
	if t.inlineMessageID != nil {
		r.setParam("inline_message_id", *t.inlineMessageID)
	}
	if t.parseMode != nil {
		r.setParam("parse_mode", *t.parseMode)
	}
	if t.entities != nil {
		r.setParam("entities", *t.entities)
	}
	if t.disableWebPagePreview != nil {
		r.setParam("disable_web_page_preview", *t.disableWebPagePreview)
	}
	if t.replyMarkup != nil {
		r.setParam("reply_markup", *t.replyMarkup)
	}

	data, err := t.c.callAPI(ctx, r, opts...)

	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*EditMessageText, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type EditMessageText struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}

type DeleteMessageService struct {
	c         *Client
	chatID    *int64
	messageID *int64
}

func (t *DeleteMessageService) ChatID(chatID int64) *DeleteMessageService {
	t.chatID = &chatID
	return t
}

func (t *DeleteMessageService) MessageID(messageID int64) *DeleteMessageService {
	t.messageID = &messageID
	return t
}

func (t *DeleteMessageService) Do(ctx context.Context, opts ...RequestOption) (res []*DeleteMessage, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/deleteMessage",
	}

	r.setParam("chat_id", *t.chatID)
	r.setParam("message_id", *t.messageID)

	data, err := t.c.callAPI(ctx, r, opts...)

	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*DeleteMessage, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type DeleteMessage struct {
	Ok     bool `json:"ok"`
	Result bool `json:"result"`
}
