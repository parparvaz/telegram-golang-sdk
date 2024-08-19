package telegram

import (
	"context"
	"net/http"
	"telegram/common"

	jsoniter "github.com/json-iterator/go"
)

type SendMessageService struct {
	c                        *Client
	chatID                   *int64
	messageThreadID          *int64
	text                     *string
	parseMode                *string
	entities                 *string
	disableWebPagePreview    *bool
	disableNotification      *bool
	protectContent           *bool
	replyToMessageID         *int64
	allowSendingWithoutReply *bool
	inlineKeyboardMarkup     *string
	replyKeyboardMarkup      *string
	replyKeyboardRemove      *string
	forceReply               *string
}

func (t *SendMessageService) ChatID(chatID int64) *SendMessageService {
	t.chatID = &chatID
	return t
}

func (t *SendMessageService) MessageThreadID(messageThreadID int64) *SendMessageService {
	t.messageThreadID = &messageThreadID
	return t
}

func (t *SendMessageService) Text(text string) *SendMessageService {
	t.text = &text
	return t
}

func (t *SendMessageService) ParseMode(parseMode string) *SendMessageService {
	t.parseMode = &parseMode
	return t
}

func (t *SendMessageService) Entities(entities []MessageEntity) *SendMessageService {
	json, err := jsoniter.Marshal(&entities)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.entities = &jsonString
	return t
}

func (t *SendMessageService) DisableWebPagePreview(disableWebPagePreview bool) *SendMessageService {
	t.disableWebPagePreview = &disableWebPagePreview
	return t
}

func (t *SendMessageService) DisableNotification(disableNotification bool) *SendMessageService {
	t.disableNotification = &disableNotification
	return t
}

func (t *SendMessageService) ProtectContent(protectContent bool) *SendMessageService {
	t.protectContent = &protectContent
	return t
}

func (t *SendMessageService) ReplyToMessageID(replyToMessageID int64) *SendMessageService {
	t.replyToMessageID = &replyToMessageID
	return t
}

func (t *SendMessageService) AllowSendingWithoutReply(allowSendingWithoutReply bool) *SendMessageService {
	t.allowSendingWithoutReply = &allowSendingWithoutReply
	return t
}

func (t *SendMessageService) InlineKeyboardMarkup(inlineKeyboardMarkup InlineKeyboardMarkup) *SendMessageService {
	json, err := jsoniter.Marshal(&inlineKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.inlineKeyboardMarkup = &jsonString
	return t
}

func (t *SendMessageService) ReplyKeyboardMarkup(replyKeyboardMarkup ReplyKeyboardMarkup) *SendMessageService {
	json, err := jsoniter.Marshal(&replyKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardMarkup = &jsonString
	return t
}

func (t *SendMessageService) ReplyKeyboardRemove(replyKeyboardRemove ReplyKeyboardRemove) *SendMessageService {
	json, err := jsoniter.Marshal(&replyKeyboardRemove)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardRemove = &jsonString
	return t
}

func (t *SendMessageService) ForceReply(forceReply ForceReply) *SendMessageService {
	json, err := jsoniter.Marshal(&forceReply)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.forceReply = &jsonString
	return t
}

func (t *SendMessageService) Do(ctx context.Context, opts ...RequestOption) (res []*SendMessage, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sendMessage",
	}

	r.setParam("chat_id", *t.chatID)
	r.setParam("text", *t.text)
	if t.messageThreadID != nil {
		r.setParam("message_thread_id", *t.messageThreadID)
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
	if t.disableNotification != nil {
		r.setParam("disable_notification", *t.disableNotification)
	}
	if t.protectContent != nil {
		r.setParam("protect_content", *t.protectContent)
	}
	if t.replyToMessageID != nil {
		r.setParam("reply_to_message_id", *t.replyToMessageID)
	}
	if t.allowSendingWithoutReply != nil {
		r.setParam("allow_sending_without_reply", *t.allowSendingWithoutReply)
	}
	if t.inlineKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.inlineKeyboardMarkup)
	}
	if t.replyKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.replyKeyboardMarkup)
	}
	if t.replyKeyboardRemove != nil {
		r.setParam("reply_markup", *t.replyKeyboardRemove)
	}
	if t.forceReply != nil {
		r.setParam("reply_markup", *t.forceReply)
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*SendMessage, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SendMessage struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}

type ForwardMessageService struct {
	c                   *Client
	chatID              *int64
	messageThreadID     *int64
	fromChatID          *string
	disableNotification *bool
	protectContent      *bool
	messageID           *int64
}

func (t *ForwardMessageService) ChatID(chatID int64) *ForwardMessageService {
	t.chatID = &chatID
	return t
}

func (t *ForwardMessageService) MessageThreadID(messageThreadID int64) *ForwardMessageService {
	t.messageThreadID = &messageThreadID
	return t
}

func (t *ForwardMessageService) MessageID(messageID int64) *ForwardMessageService {
	t.messageID = &messageID
	return t
}

func (t *ForwardMessageService) FromChatID(fromChatID string) *ForwardMessageService {
	t.fromChatID = &fromChatID
	return t
}

func (t *ForwardMessageService) DisableNotification(disableNotification bool) *ForwardMessageService {
	t.disableNotification = &disableNotification
	return t
}

func (t *ForwardMessageService) ProtectContent(protectContent bool) *ForwardMessageService {
	t.protectContent = &protectContent
	return t
}

func (t *ForwardMessageService) Do(ctx context.Context, opts ...RequestOption) (res []*ForwardMessage, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/forwardMessage",
	}

	r.setParam("chat_id", *t.chatID)
	r.setParam("from_chat_id", *t.fromChatID)
	r.setParam("message_id", *t.messageID)
	if t.messageThreadID != nil {
		r.setParam("message_thread_id", *t.messageThreadID)
	}
	if t.disableNotification != nil {
		r.setParam("disable_notification", *t.disableNotification)
	}
	if t.protectContent != nil {
		r.setParam("protect_content", *t.protectContent)
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*ForwardMessage, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type ForwardMessage struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}

type CopyMessageService struct {
	c                        *Client
	chatID                   *int64
	messageThreadID          *int64
	fromChatID               *string
	messageID                *int64
	caption                  *string
	parseMode                *string
	captionEntities          *string
	disableNotification      *bool
	protectContent           *bool
	replyToMessageID         *int64
	allowSendingWithoutReply *bool
	inlineKeyboardMarkup     *string
	replyKeyboardMarkup      *string
	replyKeyboardRemove      *string
	forceReply               *string
}

func (t *CopyMessageService) ChatID(chatID int64) *CopyMessageService {
	t.chatID = &chatID
	return t
}

func (t *CopyMessageService) MessageThreadID(messageThreadID int64) *CopyMessageService {
	t.messageThreadID = &messageThreadID
	return t
}

func (t *CopyMessageService) ParseMode(parseMode string) *CopyMessageService {
	t.parseMode = &parseMode
	return t
}

func (t *CopyMessageService) CaptionEntities(captionEntities []MessageEntity) *CopyMessageService {
	json, err := jsoniter.Marshal(&captionEntities)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.captionEntities = &jsonString

	return t
}

func (t *CopyMessageService) MessageID(messageID int64) *CopyMessageService {
	t.messageID = &messageID
	return t
}

func (t *CopyMessageService) FromChatID(fromChatID string) *CopyMessageService {
	t.fromChatID = &fromChatID
	return t
}

func (t *CopyMessageService) DisableNotification(disableNotification bool) *CopyMessageService {
	t.disableNotification = &disableNotification
	return t
}

func (t *CopyMessageService) ProtectContent(protectContent bool) *CopyMessageService {
	t.protectContent = &protectContent
	return t
}

func (t *CopyMessageService) ReplyToMessageID(replyToMessageID int64) *CopyMessageService {
	t.replyToMessageID = &replyToMessageID
	return t
}

func (t *CopyMessageService) AllowSendingWithoutReply(allowSendingWithoutReply bool) *CopyMessageService {
	t.allowSendingWithoutReply = &allowSendingWithoutReply
	return t
}

func (t *CopyMessageService) InlineKeyboardMarkup(inlineKeyboardMarkup InlineKeyboardMarkup) *CopyMessageService {
	json, err := jsoniter.Marshal(&inlineKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.inlineKeyboardMarkup = &jsonString
	return t
}

func (t *CopyMessageService) ReplyKeyboardMarkup(replyKeyboardMarkup ReplyKeyboardMarkup) *CopyMessageService {
	json, err := jsoniter.Marshal(&replyKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardMarkup = &jsonString
	return t
}

func (t *CopyMessageService) ReplyKeyboardRemove(replyKeyboardRemove ReplyKeyboardRemove) *CopyMessageService {
	json, err := jsoniter.Marshal(&replyKeyboardRemove)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardRemove = &jsonString
	return t
}

func (t *CopyMessageService) ForceReply(forceReply ForceReply) *CopyMessageService {
	json, err := jsoniter.Marshal(&forceReply)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.forceReply = &jsonString
	return t
}

func (t *CopyMessageService) Do(ctx context.Context, opts ...RequestOption) (res []*CopyMessage, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/copyMessage",
	}

	r.setParam("chat_id", *t.chatID)
	r.setParam("from_chat_id", *t.fromChatID)
	r.setParam("message_id", *t.messageID)
	if t.messageThreadID != nil {
		r.setParam("message_thread_id", *t.messageThreadID)
	}
	if t.parseMode != nil {
		r.setParam("parse_mode", *t.parseMode)
	}
	if t.captionEntities != nil {
		r.setParam("caption_entities", *t.captionEntities)
	}
	if t.disableNotification != nil {
		r.setParam("disable_notification", *t.disableNotification)
	}
	if t.protectContent != nil {
		r.setParam("protect_content", *t.protectContent)
	}
	if t.replyToMessageID != nil {
		r.setParam("reply_to_message_id", *t.replyToMessageID)
	}
	if t.allowSendingWithoutReply != nil {
		r.setParam("allow_sending_without_reply", *t.allowSendingWithoutReply)
	}
	if t.inlineKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.inlineKeyboardMarkup)
	}
	if t.replyKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.replyKeyboardMarkup)
	}
	if t.replyKeyboardRemove != nil {
		r.setParam("reply_markup", *t.replyKeyboardRemove)
	}
	if t.forceReply != nil {
		r.setParam("reply_markup", *t.forceReply)
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*CopyMessage, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type CopyMessage struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}

type SendPhotoService struct {
	c                        *Client
	chatID                   *int64
	messageThreadID          *int64
	photo                    *InputFile
	photoString              *string
	caption                  *string
	parseMode                *string
	captionEntities          *string
	disableNotification      *bool
	protectContent           *bool
	replyToMessageID         *int64
	allowSendingWithoutReply *bool
	inlineKeyboardMarkup     *string
	replyKeyboardMarkup      *string
	replyKeyboardRemove      *string
	forceReply               *string
}

func (t *SendPhotoService) ChatID(chatID int64) *SendPhotoService {
	t.chatID = &chatID
	return t
}

func (t *SendPhotoService) MessageThreadID(messageThreadID int64) *SendPhotoService {
	t.messageThreadID = &messageThreadID
	return t
}

func (t *SendPhotoService) ParseMode(parseMode string) *SendPhotoService {
	t.parseMode = &parseMode
	return t
}

func (t *SendPhotoService) CaptionEntities(captionEntities []MessageEntity) *SendPhotoService {
	json, err := jsoniter.Marshal(&captionEntities)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.captionEntities = &jsonString

	return t
}

func (t *SendPhotoService) Photo(photo InputFile) *SendPhotoService {
	t.photo = &photo
	return t
}

func (t *SendPhotoService) PhotoString(photoString string) *SendPhotoService {
	t.photoString = &photoString
	return t
}

func (t *SendPhotoService) DisableNotification(disableNotification bool) *SendPhotoService {
	t.disableNotification = &disableNotification
	return t
}

func (t *SendPhotoService) ProtectContent(protectContent bool) *SendPhotoService {
	t.protectContent = &protectContent
	return t
}

func (t *SendPhotoService) ReplyToMessageID(replyToMessageID int64) *SendPhotoService {
	t.replyToMessageID = &replyToMessageID
	return t
}

func (t *SendPhotoService) AllowSendingWithoutReply(allowSendingWithoutReply bool) *SendPhotoService {
	t.allowSendingWithoutReply = &allowSendingWithoutReply
	return t
}

func (t *SendPhotoService) InlineKeyboardMarkup(inlineKeyboardMarkup InlineKeyboardMarkup) *SendPhotoService {
	json, err := jsoniter.Marshal(&inlineKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.inlineKeyboardMarkup = &jsonString
	return t
}

func (t *SendPhotoService) ReplyKeyboardMarkup(replyKeyboardMarkup ReplyKeyboardMarkup) *SendPhotoService {
	json, err := jsoniter.Marshal(&replyKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardMarkup = &jsonString
	return t
}

func (t *SendPhotoService) ReplyKeyboardRemove(replyKeyboardRemove ReplyKeyboardRemove) *SendPhotoService {
	json, err := jsoniter.Marshal(&replyKeyboardRemove)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardRemove = &jsonString
	return t
}

func (t *SendPhotoService) ForceReply(forceReply ForceReply) *SendPhotoService {
	json, err := jsoniter.Marshal(&forceReply)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.forceReply = &jsonString
	return t
}

func (t *SendPhotoService) Do(ctx context.Context, opts ...RequestOption) (res []*SendPhoto, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sendPhoto",
	}

	r.setParam("chat_id", *t.chatID)
	if t.photo != nil {
		r.setParam("photo", *t.photo)
	}
	if t.photoString != nil {
		r.setParam("photo", *t.photoString)
	}
	if t.messageThreadID != nil {
		r.setParam("message_thread_id", *t.messageThreadID)
	}
	if t.parseMode != nil {
		r.setParam("parse_mode", *t.parseMode)
	}
	if t.captionEntities != nil {
		r.setParam("caption_entities", *t.captionEntities)
	}
	if t.disableNotification != nil {
		r.setParam("disable_notification", *t.disableNotification)
	}
	if t.protectContent != nil {
		r.setParam("protect_content", *t.protectContent)
	}
	if t.replyToMessageID != nil {
		r.setParam("reply_to_message_id", *t.replyToMessageID)
	}
	if t.allowSendingWithoutReply != nil {
		r.setParam("allow_sending_without_reply", *t.allowSendingWithoutReply)
	}
	if t.inlineKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.inlineKeyboardMarkup)
	}
	if t.replyKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.replyKeyboardMarkup)
	}
	if t.replyKeyboardRemove != nil {
		r.setParam("reply_markup", *t.replyKeyboardRemove)
	}
	if t.forceReply != nil {
		r.setParam("reply_markup", *t.forceReply)
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*SendPhoto, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SendPhoto struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}

type SendAudioService struct {
	c                        *Client
	chatID                   *int64
	messageThreadID          *int64
	audio                    *InputFile
	audioString              *string
	caption                  *string
	parseMode                *string
	captionEntities          *string
	duration                 *int64
	performer                *string
	title                    *string
	thumbnail                *InputFile
	thumbnailString          *string
	disableNotification      *bool
	protectContent           *bool
	replyToMessageID         *int64
	allowSendingWithoutReply *bool
	inlineKeyboardMarkup     *string
	replyKeyboardMarkup      *string
	replyKeyboardRemove      *string
	forceReply               *string
}

func (t *SendAudioService) ChatID(chatID int64) *SendAudioService {
	t.chatID = &chatID
	return t
}

func (t *SendAudioService) MessageThreadID(messageThreadID int64) *SendAudioService {
	t.messageThreadID = &messageThreadID
	return t
}

func (t *SendAudioService) ParseMode(parseMode string) *SendAudioService {
	t.parseMode = &parseMode
	return t
}

func (t *SendAudioService) CaptionEntities(captionEntities []MessageEntity) *SendAudioService {
	json, err := jsoniter.Marshal(&captionEntities)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.captionEntities = &jsonString

	return t
}

func (t *SendAudioService) Audio(audio InputFile) *SendAudioService {
	t.audio = &audio
	return t
}

func (t *SendAudioService) AudioString(audioString string) *SendAudioService {
	t.audioString = &audioString
	return t
}

func (t *SendAudioService) Duration(duration int64) *SendAudioService {
	t.duration = &duration
	return t
}

func (t *SendAudioService) Performer(performer string) *SendAudioService {
	t.performer = &performer
	return t
}

func (t *SendAudioService) Title(title string) *SendAudioService {
	t.title = &title
	return t
}

func (t *SendAudioService) Thumbnail(thumbnail InputFile) *SendAudioService {
	t.thumbnail = &thumbnail
	return t
}

func (t *SendAudioService) ThumbnailString(thumbnailString string) *SendAudioService {
	t.thumbnailString = &thumbnailString
	return t
}

func (t *SendAudioService) DisableNotification(disableNotification bool) *SendAudioService {
	t.disableNotification = &disableNotification
	return t
}

func (t *SendAudioService) ProtectContent(protectContent bool) *SendAudioService {
	t.protectContent = &protectContent
	return t
}

func (t *SendAudioService) ReplyToMessageID(replyToMessageID int64) *SendAudioService {
	t.replyToMessageID = &replyToMessageID
	return t
}

func (t *SendAudioService) AllowSendingWithoutReply(allowSendingWithoutReply bool) *SendAudioService {
	t.allowSendingWithoutReply = &allowSendingWithoutReply
	return t
}

func (t *SendAudioService) InlineKeyboardMarkup(inlineKeyboardMarkup InlineKeyboardMarkup) *SendAudioService {
	json, err := jsoniter.Marshal(&inlineKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.inlineKeyboardMarkup = &jsonString
	return t
}

func (t *SendAudioService) ReplyKeyboardMarkup(replyKeyboardMarkup ReplyKeyboardMarkup) *SendAudioService {
	json, err := jsoniter.Marshal(&replyKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardMarkup = &jsonString
	return t
}

func (t *SendAudioService) ReplyKeyboardRemove(replyKeyboardRemove ReplyKeyboardRemove) *SendAudioService {
	json, err := jsoniter.Marshal(&replyKeyboardRemove)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardRemove = &jsonString
	return t
}

func (t *SendAudioService) ForceReply(forceReply ForceReply) *SendAudioService {
	json, err := jsoniter.Marshal(&forceReply)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.forceReply = &jsonString
	return t
}

func (t *SendAudioService) Do(ctx context.Context, opts ...RequestOption) (res []*SendAudio, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sendAudio",
	}

	r.setParam("chat_id", *t.chatID)
	if t.audio != nil {
		r.setParam("audio", *t.audio)
	}
	if t.audioString != nil {
		r.setParam("audio", *t.audioString)
	}
	if t.messageThreadID != nil {
		r.setParam("message_thread_id", *t.messageThreadID)
	}
	if t.parseMode != nil {
		r.setParam("parse_mode", *t.parseMode)
	}
	if t.captionEntities != nil {
		r.setParam("caption_entities", *t.captionEntities)
	}
	if t.duration != nil {
		r.setParam("duration", *t.duration)
	}
	if t.performer != nil {
		r.setParam("performer", *t.performer)
	}
	if t.title != nil {
		r.setParam("title", *t.title)
	}
	if t.thumbnail != nil {
		r.setParam("thumbnail", *t.thumbnail)
	}
	if t.thumbnailString != nil {
		r.setParam("thumbnail", *t.thumbnailString)
	}
	if t.disableNotification != nil {
		r.setParam("disable_notification", *t.disableNotification)
	}
	if t.protectContent != nil {
		r.setParam("protect_content", *t.protectContent)
	}
	if t.replyToMessageID != nil {
		r.setParam("reply_to_message_id", *t.replyToMessageID)
	}
	if t.allowSendingWithoutReply != nil {
		r.setParam("allow_sending_without_reply", *t.allowSendingWithoutReply)
	}
	if t.inlineKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.inlineKeyboardMarkup)
	}
	if t.replyKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.replyKeyboardMarkup)
	}
	if t.replyKeyboardRemove != nil {
		r.setParam("reply_markup", *t.replyKeyboardRemove)
	}
	if t.forceReply != nil {
		r.setParam("reply_markup", *t.forceReply)
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*SendAudio, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SendAudio struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}

type SendDocumentService struct {
	c                        *Client
	chatID                   *int64
	messageThreadID          *int64
	document                 *InputFile
	documentString           *string
	thumbnail                *InputFile
	thumbnailString          *string
	caption                  *string
	parseMode                *string
	captionEntities          *string
	disableNotification      *bool
	protectContent           *bool
	replyToMessageID         *int64
	allowSendingWithoutReply *bool
	inlineKeyboardMarkup     *string
	replyKeyboardMarkup      *string
	replyKeyboardRemove      *string
	forceReply               *string
}

func (t *SendDocumentService) ChatID(chatID int64) *SendDocumentService {
	t.chatID = &chatID
	return t
}

func (t *SendDocumentService) MessageThreadID(messageThreadID int64) *SendDocumentService {
	t.messageThreadID = &messageThreadID
	return t
}

func (t *SendDocumentService) ParseMode(parseMode string) *SendDocumentService {
	t.parseMode = &parseMode
	return t
}

func (t *SendDocumentService) CaptionEntities(captionEntities []MessageEntity) *SendDocumentService {
	json, err := jsoniter.Marshal(&captionEntities)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.captionEntities = &jsonString

	return t
}

func (t *SendDocumentService) Document(document InputFile) *SendDocumentService {
	t.document = &document
	return t
}

func (t *SendDocumentService) DocumentString(documentString string) *SendDocumentService {
	t.documentString = &documentString
	return t
}

func (t *SendDocumentService) Thumbnail(thumbnail InputFile) *SendDocumentService {
	t.thumbnail = &thumbnail
	return t
}

func (t *SendDocumentService) ThumbnailString(thumbnailString string) *SendDocumentService {
	t.thumbnailString = &thumbnailString
	return t
}

func (t *SendDocumentService) DisableNotification(disableNotification bool) *SendDocumentService {
	t.disableNotification = &disableNotification
	return t
}

func (t *SendDocumentService) ProtectContent(protectContent bool) *SendDocumentService {
	t.protectContent = &protectContent
	return t
}

func (t *SendDocumentService) ReplyToMessageID(replyToMessageID int64) *SendDocumentService {
	t.replyToMessageID = &replyToMessageID
	return t
}

func (t *SendDocumentService) AllowSendingWithoutReply(allowSendingWithoutReply bool) *SendDocumentService {
	t.allowSendingWithoutReply = &allowSendingWithoutReply
	return t
}

func (t *SendDocumentService) InlineKeyboardMarkup(inlineKeyboardMarkup InlineKeyboardMarkup) *SendDocumentService {
	json, err := jsoniter.Marshal(&inlineKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.inlineKeyboardMarkup = &jsonString
	return t
}

func (t *SendDocumentService) ReplyKeyboardMarkup(replyKeyboardMarkup ReplyKeyboardMarkup) *SendDocumentService {
	json, err := jsoniter.Marshal(&replyKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardMarkup = &jsonString
	return t
}

func (t *SendDocumentService) ReplyKeyboardRemove(replyKeyboardRemove ReplyKeyboardRemove) *SendDocumentService {
	json, err := jsoniter.Marshal(&replyKeyboardRemove)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardRemove = &jsonString
	return t
}

func (t *SendDocumentService) ForceReply(forceReply ForceReply) *SendDocumentService {
	json, err := jsoniter.Marshal(&forceReply)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.forceReply = &jsonString
	return t
}

func (t *SendDocumentService) Do(ctx context.Context, opts ...RequestOption) (res []*SendDocument, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sendDocument",
	}

	r.setParam("chat_id", *t.chatID)
	if t.document != nil {
		r.setParam("document", *t.document)
	}
	if t.documentString != nil {
		r.setParam("document", *t.documentString)
	}
	if t.messageThreadID != nil {
		r.setParam("message_thread_id", *t.messageThreadID)
	}
	if t.parseMode != nil {
		r.setParam("parse_mode", *t.parseMode)
	}
	if t.captionEntities != nil {
		r.setParam("caption_entities", *t.captionEntities)
	}
	if t.caption != nil {
		r.setParam("caption", *t.caption)
	}
	if t.disableNotification != nil {
		r.setParam("disable_notification", *t.disableNotification)
	}
	if t.protectContent != nil {
		r.setParam("protect_content", *t.protectContent)
	}
	if t.replyToMessageID != nil {
		r.setParam("reply_to_message_id", *t.replyToMessageID)
	}
	if t.allowSendingWithoutReply != nil {
		r.setParam("allow_sending_without_reply", *t.allowSendingWithoutReply)
	}
	if t.inlineKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.inlineKeyboardMarkup)
	}
	if t.replyKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.replyKeyboardMarkup)
	}
	if t.replyKeyboardRemove != nil {
		r.setParam("reply_markup", *t.replyKeyboardRemove)
	}
	if t.forceReply != nil {
		r.setParam("reply_markup", *t.forceReply)
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*SendDocument, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SendDocument struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}

type SendVideoService struct {
	c                        *Client
	chatID                   *int64
	messageThreadID          *int64
	video                    *InputFile
	videoString              *string
	duration                 *int64
	width                    *int64
	height                   *int64
	thumbnail                *InputFile
	thumbnailString          *string
	caption                  *string
	parseMode                *string
	captionEntities          *string
	hasSpoiler               *bool
	supportsStreaming        *bool
	disableNotification      *bool
	protectContent           *bool
	replyToMessageID         *int64
	allowSendingWithoutReply *bool
	inlineKeyboardMarkup     *string
	replyKeyboardMarkup      *string
	replyKeyboardRemove      *string
	forceReply               *string
}

func (t *SendVideoService) ChatID(chatID int64) *SendVideoService {
	t.chatID = &chatID
	return t
}

func (t *SendVideoService) MessageThreadID(messageThreadID int64) *SendVideoService {
	t.messageThreadID = &messageThreadID
	return t
}

func (t *SendVideoService) ParseMode(parseMode string) *SendVideoService {
	t.parseMode = &parseMode
	return t
}

func (t *SendVideoService) CaptionEntities(captionEntities []MessageEntity) *SendVideoService {
	json, err := jsoniter.Marshal(&captionEntities)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.captionEntities = &jsonString

	return t
}

func (t *SendVideoService) Video(video InputFile) *SendVideoService {
	t.video = &video
	return t
}

func (t *SendVideoService) VideoString(videoString string) *SendVideoService {
	t.videoString = &videoString
	return t
}

func (t *SendVideoService) Duration(duration int64) *SendVideoService {
	t.duration = &duration
	return t
}

func (t *SendVideoService) Width(width int64) *SendVideoService {
	t.width = &width
	return t
}

func (t *SendVideoService) Height(height int64) *SendVideoService {
	t.height = &height
	return t
}

func (t *SendVideoService) Thumbnail(thumbnail InputFile) *SendVideoService {
	t.thumbnail = &thumbnail
	return t
}

func (t *SendVideoService) ThumbnailString(thumbnailString string) *SendVideoService {
	t.thumbnailString = &thumbnailString
	return t
}

func (t *SendVideoService) DisableNotification(disableNotification bool) *SendVideoService {
	t.disableNotification = &disableNotification
	return t
}

func (t *SendVideoService) ProtectContent(protectContent bool) *SendVideoService {
	t.protectContent = &protectContent
	return t
}

func (t *SendVideoService) HasSpoiler(hasSpoiler bool) *SendVideoService {
	t.hasSpoiler = &hasSpoiler
	return t
}

func (t *SendVideoService) SupportsStreaming(supportsStreaming bool) *SendVideoService {
	t.supportsStreaming = &supportsStreaming
	return t
}

func (t *SendVideoService) ReplyToMessageID(replyToMessageID int64) *SendVideoService {
	t.replyToMessageID = &replyToMessageID
	return t
}

func (t *SendVideoService) AllowSendingWithoutReply(allowSendingWithoutReply bool) *SendVideoService {
	t.allowSendingWithoutReply = &allowSendingWithoutReply
	return t
}

func (t *SendVideoService) InlineKeyboardMarkup(inlineKeyboardMarkup InlineKeyboardMarkup) *SendVideoService {
	json, err := jsoniter.Marshal(&inlineKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.inlineKeyboardMarkup = &jsonString
	return t
}

func (t *SendVideoService) ReplyKeyboardMarkup(replyKeyboardMarkup ReplyKeyboardMarkup) *SendVideoService {
	json, err := jsoniter.Marshal(&replyKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardMarkup = &jsonString
	return t
}

func (t *SendVideoService) ReplyKeyboardRemove(replyKeyboardRemove ReplyKeyboardRemove) *SendVideoService {
	json, err := jsoniter.Marshal(&replyKeyboardRemove)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardRemove = &jsonString
	return t
}

func (t *SendVideoService) ForceReply(forceReply ForceReply) *SendVideoService {
	json, err := jsoniter.Marshal(&forceReply)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.forceReply = &jsonString
	return t
}

func (t *SendVideoService) Do(ctx context.Context, opts ...RequestOption) (res []*SendVideo, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sendVideo",
	}

	r.setParam("chat_id", *t.chatID)
	if t.video != nil {
		r.setParam("video", *t.video)
	}
	if t.videoString != nil {
		r.setParam("video", *t.videoString)
	}
	if t.duration != nil {
		r.setParam("duration", *t.duration)
	}
	if t.width != nil {
		r.setParam("width", *t.width)
	}
	if t.height != nil {
		r.setParam("height", *t.height)
	}
	if t.messageThreadID != nil {
		r.setParam("message_thread_id", *t.messageThreadID)
	}
	if t.parseMode != nil {
		r.setParam("parse_mode", *t.parseMode)
	}
	if t.captionEntities != nil {
		r.setParam("caption_entities", *t.captionEntities)
	}
	if t.caption != nil {
		r.setParam("caption", *t.caption)
	}
	if t.disableNotification != nil {
		r.setParam("disable_notification", *t.disableNotification)
	}
	if t.protectContent != nil {
		r.setParam("protect_content", *t.protectContent)
	}
	if t.replyToMessageID != nil {
		r.setParam("reply_to_message_id", *t.replyToMessageID)
	}
	if t.hasSpoiler != nil {
		r.setParam("has_spoiler", *t.hasSpoiler)
	}
	if t.supportsStreaming != nil {
		r.setParam("supports_streaming", *t.supportsStreaming)
	}
	if t.allowSendingWithoutReply != nil {
		r.setParam("allow_sending_without_reply", *t.allowSendingWithoutReply)
	}
	if t.inlineKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.inlineKeyboardMarkup)
	}
	if t.replyKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.replyKeyboardMarkup)
	}
	if t.replyKeyboardRemove != nil {
		r.setParam("reply_markup", *t.replyKeyboardRemove)
	}
	if t.forceReply != nil {
		r.setParam("reply_markup", *t.forceReply)
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*SendVideo, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SendVideo struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}

type SendAnimationService struct {
	c                        *Client
	chatID                   *int64
	messageThreadID          *int64
	animation                *InputFile
	animationString          *string
	duration                 *int64
	width                    *int64
	height                   *int64
	thumbnail                *InputFile
	thumbnailString          *string
	caption                  *string
	parseMode                *string
	captionEntities          *string
	supportsStreaming        *bool
	disableNotification      *bool
	protectContent           *bool
	replyToMessageID         *int64
	allowSendingWithoutReply *bool
	inlineKeyboardMarkup     *string
	replyKeyboardMarkup      *string
	replyKeyboardRemove      *string
	forceReply               *string
}

func (t *SendAnimationService) ChatID(chatID int64) *SendAnimationService {
	t.chatID = &chatID
	return t
}

func (t *SendAnimationService) MessageThreadID(messageThreadID int64) *SendAnimationService {
	t.messageThreadID = &messageThreadID
	return t
}

func (t *SendAnimationService) ParseMode(parseMode string) *SendAnimationService {
	t.parseMode = &parseMode
	return t
}

func (t *SendAnimationService) CaptionEntities(captionEntities []MessageEntity) *SendAnimationService {
	json, err := jsoniter.Marshal(&captionEntities)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.captionEntities = &jsonString

	return t
}

func (t *SendAnimationService) Animation(animation InputFile) *SendAnimationService {
	t.animation = &animation
	return t
}

func (t *SendAnimationService) AnimationString(animationString string) *SendAnimationService {
	t.animationString = &animationString
	return t
}

func (t *SendAnimationService) Duration(duration int64) *SendAnimationService {
	t.duration = &duration
	return t
}

func (t *SendAnimationService) Width(width int64) *SendAnimationService {
	t.width = &width
	return t
}

func (t *SendAnimationService) Height(height int64) *SendAnimationService {
	t.height = &height
	return t
}

func (t *SendAnimationService) Thumbnail(thumbnail InputFile) *SendAnimationService {
	t.thumbnail = &thumbnail
	return t
}

func (t *SendAnimationService) ThumbnailString(thumbnailString string) *SendAnimationService {
	t.thumbnailString = &thumbnailString
	return t
}

func (t *SendAnimationService) DisableNotification(disableNotification bool) *SendAnimationService {
	t.disableNotification = &disableNotification
	return t
}

func (t *SendAnimationService) ProtectContent(protectContent bool) *SendAnimationService {
	t.protectContent = &protectContent
	return t
}

func (t *SendAnimationService) SupportsStreaming(supportsStreaming bool) *SendAnimationService {
	t.supportsStreaming = &supportsStreaming
	return t
}

func (t *SendAnimationService) ReplyToMessageID(replyToMessageID int64) *SendAnimationService {
	t.replyToMessageID = &replyToMessageID
	return t
}

func (t *SendAnimationService) AllowSendingWithoutReply(allowSendingWithoutReply bool) *SendAnimationService {
	t.allowSendingWithoutReply = &allowSendingWithoutReply
	return t
}

func (t *SendAnimationService) InlineKeyboardMarkup(inlineKeyboardMarkup InlineKeyboardMarkup) *SendAnimationService {
	json, err := jsoniter.Marshal(&inlineKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.inlineKeyboardMarkup = &jsonString
	return t
}

func (t *SendAnimationService) ReplyKeyboardMarkup(replyKeyboardMarkup ReplyKeyboardMarkup) *SendAnimationService {
	json, err := jsoniter.Marshal(&replyKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardMarkup = &jsonString
	return t
}

func (t *SendAnimationService) ReplyKeyboardRemove(replyKeyboardRemove ReplyKeyboardRemove) *SendAnimationService {
	json, err := jsoniter.Marshal(&replyKeyboardRemove)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardRemove = &jsonString
	return t
}

func (t *SendAnimationService) ForceReply(forceReply ForceReply) *SendAnimationService {
	json, err := jsoniter.Marshal(&forceReply)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.forceReply = &jsonString
	return t
}

func (t *SendAnimationService) Do(ctx context.Context, opts ...RequestOption) (res []*SendAnimation, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sendAnimation",
	}

	r.setParam("chat_id", *t.chatID)
	if t.animation != nil {
		r.setParam("animation", *t.animation)
	}
	if t.animationString != nil {
		r.setParam("animation", *t.animationString)
	}
	if t.duration != nil {
		r.setParam("duration", *t.duration)
	}
	if t.width != nil {
		r.setParam("width", *t.width)
	}
	if t.height != nil {
		r.setParam("height", *t.height)
	}
	if t.messageThreadID != nil {
		r.setParam("message_thread_id", *t.messageThreadID)
	}
	if t.parseMode != nil {
		r.setParam("parse_mode", *t.parseMode)
	}
	if t.captionEntities != nil {
		r.setParam("caption_entities", *t.captionEntities)
	}
	if t.caption != nil {
		r.setParam("caption", *t.caption)
	}
	if t.disableNotification != nil {
		r.setParam("disable_notification", *t.disableNotification)
	}
	if t.protectContent != nil {
		r.setParam("protect_content", *t.protectContent)
	}
	if t.replyToMessageID != nil {
		r.setParam("reply_to_message_id", *t.replyToMessageID)
	}
	if t.supportsStreaming != nil {
		r.setParam("supports_streaming", *t.supportsStreaming)
	}
	if t.allowSendingWithoutReply != nil {
		r.setParam("allow_sending_without_reply", *t.allowSendingWithoutReply)
	}
	if t.inlineKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.inlineKeyboardMarkup)
	}
	if t.replyKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.replyKeyboardMarkup)
	}
	if t.replyKeyboardRemove != nil {
		r.setParam("reply_markup", *t.replyKeyboardRemove)
	}
	if t.forceReply != nil {
		r.setParam("reply_markup", *t.forceReply)
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*SendAnimation, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SendAnimation struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}

type SendVoiceService struct {
	c                        *Client
	chatID                   *int64
	messageThreadID          *int64
	voice                    *InputFile
	voiceString              *string
	caption                  *string
	parseMode                *string
	captionEntities          *string
	duration                 *int64
	disableNotification      *bool
	protectContent           *bool
	replyToMessageID         *int64
	allowSendingWithoutReply *bool
	inlineKeyboardMarkup     *string
	replyKeyboardMarkup      *string
	replyKeyboardRemove      *string
	forceReply               *string
}

func (t *SendVoiceService) ChatID(chatID int64) *SendVoiceService {
	t.chatID = &chatID
	return t
}

func (t *SendVoiceService) MessageThreadID(messageThreadID int64) *SendVoiceService {
	t.messageThreadID = &messageThreadID
	return t
}

func (t *SendVoiceService) ParseMode(parseMode string) *SendVoiceService {
	t.parseMode = &parseMode
	return t
}

func (t *SendVoiceService) CaptionEntities(captionEntities []MessageEntity) *SendVoiceService {
	json, err := jsoniter.Marshal(&captionEntities)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.captionEntities = &jsonString

	return t
}

func (t *SendVoiceService) Voice(voice InputFile) *SendVoiceService {
	t.voice = &voice
	return t
}

func (t *SendVoiceService) VoiceString(voiceString string) *SendVoiceService {
	t.voiceString = &voiceString
	return t
}

func (t *SendVoiceService) Duration(duration int64) *SendVoiceService {
	t.duration = &duration
	return t
}

func (t *SendVoiceService) DisableNotification(disableNotification bool) *SendVoiceService {
	t.disableNotification = &disableNotification
	return t
}

func (t *SendVoiceService) ProtectContent(protectContent bool) *SendVoiceService {
	t.protectContent = &protectContent
	return t
}

func (t *SendVoiceService) ReplyToMessageID(replyToMessageID int64) *SendVoiceService {
	t.replyToMessageID = &replyToMessageID
	return t
}

func (t *SendVoiceService) AllowSendingWithoutReply(allowSendingWithoutReply bool) *SendVoiceService {
	t.allowSendingWithoutReply = &allowSendingWithoutReply
	return t
}

func (t *SendVoiceService) InlineKeyboardMarkup(inlineKeyboardMarkup InlineKeyboardMarkup) *SendVoiceService {
	json, err := jsoniter.Marshal(&inlineKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.inlineKeyboardMarkup = &jsonString
	return t
}

func (t *SendVoiceService) ReplyKeyboardMarkup(replyKeyboardMarkup ReplyKeyboardMarkup) *SendVoiceService {
	json, err := jsoniter.Marshal(&replyKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardMarkup = &jsonString
	return t
}

func (t *SendVoiceService) ReplyKeyboardRemove(replyKeyboardRemove ReplyKeyboardRemove) *SendVoiceService {
	json, err := jsoniter.Marshal(&replyKeyboardRemove)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardRemove = &jsonString
	return t
}

func (t *SendVoiceService) ForceReply(forceReply ForceReply) *SendVoiceService {
	json, err := jsoniter.Marshal(&forceReply)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.forceReply = &jsonString
	return t
}

func (t *SendVoiceService) Do(ctx context.Context, opts ...RequestOption) (res []*SendVoice, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sendVoice",
	}

	r.setParam("chat_id", *t.chatID)
	if t.voice != nil {
		r.setParam("voice", *t.voice)
	}
	if t.voiceString != nil {
		r.setParam("voice", *t.voiceString)
	}
	if t.duration != nil {
		r.setParam("duration", *t.duration)
	}
	if t.messageThreadID != nil {
		r.setParam("message_thread_id", *t.messageThreadID)
	}
	if t.parseMode != nil {
		r.setParam("parse_mode", *t.parseMode)
	}
	if t.captionEntities != nil {
		r.setParam("caption_entities", *t.captionEntities)
	}
	if t.caption != nil {
		r.setParam("caption", *t.caption)
	}
	if t.disableNotification != nil {
		r.setParam("disable_notification", *t.disableNotification)
	}
	if t.protectContent != nil {
		r.setParam("protect_content", *t.protectContent)
	}
	if t.replyToMessageID != nil {
		r.setParam("reply_to_message_id", *t.replyToMessageID)
	}
	if t.allowSendingWithoutReply != nil {
		r.setParam("allow_sending_without_reply", *t.allowSendingWithoutReply)
	}
	if t.inlineKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.inlineKeyboardMarkup)
	}
	if t.replyKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.replyKeyboardMarkup)
	}
	if t.replyKeyboardRemove != nil {
		r.setParam("reply_markup", *t.replyKeyboardRemove)
	}
	if t.forceReply != nil {
		r.setParam("reply_markup", *t.forceReply)
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*SendVoice, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SendVoice struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}

type SendVideoNoteService struct {
	c                        *Client
	chatID                   *int64
	messageThreadID          *int64
	videoNote                *InputFile
	videoNoteString          *string
	duration                 *int64
	length                   *int64
	thumbnail                *InputFile
	thumbnailString          *string
	disableNotification      *bool
	protectContent           *bool
	replyToMessageID         *int64
	allowSendingWithoutReply *bool
	inlineKeyboardMarkup     *string
	replyKeyboardMarkup      *string
	replyKeyboardRemove      *string
	forceReply               *string
}

func (t *SendVideoNoteService) ChatID(chatID int64) *SendVideoNoteService {
	t.chatID = &chatID
	return t
}

func (t *SendVideoNoteService) MessageThreadID(messageThreadID int64) *SendVideoNoteService {
	t.messageThreadID = &messageThreadID
	return t
}

func (t *SendVideoNoteService) Animation(videoNote InputFile) *SendVideoNoteService {
	t.videoNote = &videoNote
	return t
}

func (t *SendVideoNoteService) AnimationString(videoNoteString string) *SendVideoNoteService {
	t.videoNoteString = &videoNoteString
	return t
}

func (t *SendVideoNoteService) Duration(duration int64) *SendVideoNoteService {
	t.duration = &duration
	return t
}

func (t *SendVideoNoteService) Length(length int64) *SendVideoNoteService {
	t.length = &length
	return t
}

func (t *SendVideoNoteService) Thumbnail(thumbnail InputFile) *SendVideoNoteService {
	t.thumbnail = &thumbnail
	return t
}

func (t *SendVideoNoteService) ThumbnailString(thumbnailString string) *SendVideoNoteService {
	t.thumbnailString = &thumbnailString
	return t
}

func (t *SendVideoNoteService) DisableNotification(disableNotification bool) *SendVideoNoteService {
	t.disableNotification = &disableNotification
	return t
}

func (t *SendVideoNoteService) ProtectContent(protectContent bool) *SendVideoNoteService {
	t.protectContent = &protectContent
	return t
}

func (t *SendVideoNoteService) ReplyToMessageID(replyToMessageID int64) *SendVideoNoteService {
	t.replyToMessageID = &replyToMessageID
	return t
}

func (t *SendVideoNoteService) AllowSendingWithoutReply(allowSendingWithoutReply bool) *SendVideoNoteService {
	t.allowSendingWithoutReply = &allowSendingWithoutReply
	return t
}

func (t *SendVideoNoteService) InlineKeyboardMarkup(inlineKeyboardMarkup InlineKeyboardMarkup) *SendVideoNoteService {
	json, err := jsoniter.Marshal(&inlineKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.inlineKeyboardMarkup = &jsonString
	return t
}

func (t *SendVideoNoteService) ReplyKeyboardMarkup(replyKeyboardMarkup ReplyKeyboardMarkup) *SendVideoNoteService {
	json, err := jsoniter.Marshal(&replyKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardMarkup = &jsonString
	return t
}

func (t *SendVideoNoteService) ReplyKeyboardRemove(replyKeyboardRemove ReplyKeyboardRemove) *SendVideoNoteService {
	json, err := jsoniter.Marshal(&replyKeyboardRemove)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardRemove = &jsonString
	return t
}

func (t *SendVideoNoteService) ForceReply(forceReply ForceReply) *SendVideoNoteService {
	json, err := jsoniter.Marshal(&forceReply)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.forceReply = &jsonString
	return t
}

func (t *SendVideoNoteService) Do(ctx context.Context, opts ...RequestOption) (res []*SendVideoNote, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sendVideoNote",
	}

	r.setParam("chat_id", *t.chatID)
	if t.videoNote != nil {
		r.setParam("video_note", *t.videoNote)
	}
	if t.videoNoteString != nil {
		r.setParam("video_note", *t.videoNoteString)
	}
	if t.duration != nil {
		r.setParam("duration", *t.duration)
	}
	if t.length != nil {
		r.setParam("length", *t.length)
	}
	if t.messageThreadID != nil {
		r.setParam("message_thread_id", *t.messageThreadID)
	}
	if t.disableNotification != nil {
		r.setParam("disable_notification", *t.disableNotification)
	}
	if t.protectContent != nil {
		r.setParam("protect_content", *t.protectContent)
	}
	if t.replyToMessageID != nil {
		r.setParam("reply_to_message_id", *t.replyToMessageID)
	}
	if t.allowSendingWithoutReply != nil {
		r.setParam("allow_sending_without_reply", *t.allowSendingWithoutReply)
	}
	if t.inlineKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.inlineKeyboardMarkup)
	}
	if t.replyKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.replyKeyboardMarkup)
	}
	if t.replyKeyboardRemove != nil {
		r.setParam("reply_markup", *t.replyKeyboardRemove)
	}
	if t.forceReply != nil {
		r.setParam("reply_markup", *t.forceReply)
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*SendVideoNote, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SendVideoNote struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}

type SendMediaGroupService struct {
	c                        *Client
	chatID                   *int64
	messageThreadID          *int64
	inputMediaAudio          *string
	inputMediaDocument       *string
	inputMediaPhoto          *string
	inputMediaVideo          *string
	disableNotification      *bool
	protectContent           *bool
	replyToMessageID         *int64
	allowSendingWithoutReply *bool
}

func (t *SendMediaGroupService) ChatID(chatID int64) *SendMediaGroupService {
	t.chatID = &chatID
	return t
}

func (t *SendMediaGroupService) MessageThreadID(messageThreadID int64) *SendMediaGroupService {
	t.messageThreadID = &messageThreadID
	return t
}

func (t *SendMediaGroupService) DisableNotification(disableNotification bool) *SendMediaGroupService {
	t.disableNotification = &disableNotification
	return t
}

func (t *SendMediaGroupService) ProtectContent(protectContent bool) *SendMediaGroupService {
	t.protectContent = &protectContent
	return t
}

func (t *SendMediaGroupService) ReplyToMessageID(replyToMessageID int64) *SendMediaGroupService {
	t.replyToMessageID = &replyToMessageID
	return t
}

func (t *SendMediaGroupService) AllowSendingWithoutReply(allowSendingWithoutReply bool) *SendMediaGroupService {
	t.allowSendingWithoutReply = &allowSendingWithoutReply
	return t
}

func (t *SendMediaGroupService) InputMediaAudio(inputMediaAudio InputMediaAudio) *SendMediaGroupService {
	json, err := jsoniter.Marshal(&inputMediaAudio)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.inputMediaAudio = &jsonString
	return t
}

func (t *SendMediaGroupService) InputMediaDocument(inputMediaDocument InputMediaDocument) *SendMediaGroupService {
	json, err := jsoniter.Marshal(&inputMediaDocument)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.inputMediaDocument = &jsonString
	return t
}

func (t *SendMediaGroupService) InputMediaPhoto(inputMediaPhoto InputMediaPhoto) *SendMediaGroupService {
	json, err := jsoniter.Marshal(&inputMediaPhoto)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.inputMediaPhoto = &jsonString
	return t
}

func (t *SendMediaGroupService) MediaInputMediaVideo(inputMediaVideo InputMediaVideo) *SendMediaGroupService {
	json, err := jsoniter.Marshal(&inputMediaVideo)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.inputMediaVideo = &jsonString
	return t
}

func (t *SendMediaGroupService) Do(ctx context.Context, opts ...RequestOption) (res []*SendMediaGroup, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sendMediaGroup",
	}

	r.setParam("chat_id", *t.chatID)
	if t.messageThreadID != nil {
		r.setParam("message_thread_id", *t.messageThreadID)
	}
	if t.disableNotification != nil {
		r.setParam("disable_notification", *t.disableNotification)
	}
	if t.protectContent != nil {
		r.setParam("protect_content", *t.protectContent)
	}
	if t.replyToMessageID != nil {
		r.setParam("reply_to_message_id", *t.replyToMessageID)
	}
	if t.allowSendingWithoutReply != nil {
		r.setParam("allow_sending_without_reply", *t.allowSendingWithoutReply)
	}
	if t.inputMediaAudio != nil {
		r.setParam("media", *t.inputMediaAudio)
	}
	if t.inputMediaDocument != nil {
		r.setParam("media", *t.inputMediaDocument)
	}
	if t.inputMediaPhoto != nil {
		r.setParam("media", *t.inputMediaPhoto)
	}
	if t.inputMediaVideo != nil {
		r.setParam("media", *t.inputMediaVideo)
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*SendMediaGroup, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SendMediaGroup struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}

type SendLocationService struct {
	c                        *Client
	chatID                   *int64
	messageThreadID          *int64
	latitude                 *float64
	longitude                *float64
	horizontalAccuracy       *float64
	livePeriod               *int64
	heading                  *int64
	proximityAlertRadius     *int64
	disableNotification      *bool
	protectContent           *bool
	replyToMessageID         *int64
	allowSendingWithoutReply *bool
	inlineKeyboardMarkup     *string
	replyKeyboardMarkup      *string
	replyKeyboardRemove      *string
	forceReply               *string
}

func (t *SendLocationService) ChatID(chatID int64) *SendLocationService {
	t.chatID = &chatID
	return t
}

func (t *SendLocationService) MessageThreadID(messageThreadID int64) *SendLocationService {
	t.messageThreadID = &messageThreadID
	return t
}

func (t *SendLocationService) Latitude(latitude float64) *SendLocationService {
	t.latitude = &latitude
	return t
}

func (t *SendLocationService) Longitude(longitude float64) *SendLocationService {
	t.longitude = &longitude
	return t
}

func (t *SendLocationService) HorizontalAccuracy(horizontalAccuracy float64) *SendLocationService {
	t.horizontalAccuracy = &horizontalAccuracy
	return t
}

func (t *SendLocationService) LivePeriod(livePeriod int64) *SendLocationService {
	t.livePeriod = &livePeriod
	return t
}

func (t *SendLocationService) Heading(heading int64) *SendLocationService {
	t.heading = &heading
	return t
}

func (t *SendLocationService) ProximityAlertRadius(proximityAlertRadius int64) *SendLocationService {
	t.proximityAlertRadius = &proximityAlertRadius
	return t
}

func (t *SendLocationService) DisableNotification(disableNotification bool) *SendLocationService {
	t.disableNotification = &disableNotification
	return t
}

func (t *SendLocationService) ProtectContent(protectContent bool) *SendLocationService {
	t.protectContent = &protectContent
	return t
}

func (t *SendLocationService) ReplyToMessageID(replyToMessageID int64) *SendLocationService {
	t.replyToMessageID = &replyToMessageID
	return t
}

func (t *SendLocationService) AllowSendingWithoutReply(allowSendingWithoutReply bool) *SendLocationService {
	t.allowSendingWithoutReply = &allowSendingWithoutReply
	return t
}

func (t *SendLocationService) InlineKeyboardMarkup(inlineKeyboardMarkup InlineKeyboardMarkup) *SendLocationService {
	json, err := jsoniter.Marshal(&inlineKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.inlineKeyboardMarkup = &jsonString
	return t
}

func (t *SendLocationService) ReplyKeyboardMarkup(replyKeyboardMarkup ReplyKeyboardMarkup) *SendLocationService {
	json, err := jsoniter.Marshal(&replyKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardMarkup = &jsonString
	return t
}

func (t *SendLocationService) ReplyKeyboardRemove(replyKeyboardRemove ReplyKeyboardRemove) *SendLocationService {
	json, err := jsoniter.Marshal(&replyKeyboardRemove)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardRemove = &jsonString
	return t
}

func (t *SendLocationService) ForceReply(forceReply ForceReply) *SendLocationService {
	json, err := jsoniter.Marshal(&forceReply)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.forceReply = &jsonString
	return t
}

func (t *SendLocationService) Do(ctx context.Context, opts ...RequestOption) (res []*SendLocation, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sendLocation",
	}

	r.setParam("chat_id", *t.chatID)
	r.setParam("latitude", *t.latitude)
	r.setParam("longitude", *t.longitude)
	if t.horizontalAccuracy != nil {
		r.setParam("horizontal_accuracy", *t.horizontalAccuracy)
	}
	if t.livePeriod != nil {
		r.setParam("live_period", *t.livePeriod)
	}
	if t.heading != nil {
		r.setParam("heading", *t.heading)
	}
	if t.proximityAlertRadius != nil {
		r.setParam("proximity_alert_radius", *t.proximityAlertRadius)
	}
	if t.messageThreadID != nil {
		r.setParam("message_thread_id", *t.messageThreadID)
	}
	if t.disableNotification != nil {
		r.setParam("disable_notification", *t.disableNotification)
	}
	if t.protectContent != nil {
		r.setParam("protect_content", *t.protectContent)
	}
	if t.replyToMessageID != nil {
		r.setParam("reply_to_message_id", *t.replyToMessageID)
	}
	if t.allowSendingWithoutReply != nil {
		r.setParam("allow_sending_without_reply", *t.allowSendingWithoutReply)
	}
	if t.inlineKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.inlineKeyboardMarkup)
	}
	if t.replyKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.replyKeyboardMarkup)
	}
	if t.replyKeyboardRemove != nil {
		r.setParam("reply_markup", *t.replyKeyboardRemove)
	}
	if t.forceReply != nil {
		r.setParam("reply_markup", *t.forceReply)
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*SendLocation, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SendLocation struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}

type SendVenueService struct {
	c                        *Client
	chatID                   *int64
	messageThreadID          *int64
	latitude                 *float64
	longitude                *float64
	title                    *string
	address                  *string
	foursquareID             *string
	foursquareType           *string
	googlePlaceID            *string
	googlePlaceType          *string
	disableNotification      *bool
	protectContent           *bool
	replyToMessageID         *int64
	allowSendingWithoutReply *bool
	inlineKeyboardMarkup     *string
	replyKeyboardMarkup      *string
	replyKeyboardRemove      *string
	forceReply               *string
}

func (t *SendVenueService) ChatID(chatID int64) *SendVenueService {
	t.chatID = &chatID
	return t
}

func (t *SendVenueService) MessageThreadID(messageThreadID int64) *SendVenueService {
	t.messageThreadID = &messageThreadID
	return t
}

func (t *SendVenueService) Latitude(latitude float64) *SendVenueService {
	t.latitude = &latitude
	return t
}

func (t *SendVenueService) Longitude(longitude float64) *SendVenueService {
	t.longitude = &longitude
	return t
}

func (t *SendVenueService) Title(title string) *SendVenueService {
	t.title = &title
	return t
}

func (t *SendVenueService) Address(address string) *SendVenueService {
	t.address = &address
	return t
}

func (t *SendVenueService) FoursquareID(foursquareID string) *SendVenueService {
	t.foursquareID = &foursquareID
	return t
}

func (t *SendVenueService) FoursquareType(foursquareType string) *SendVenueService {
	t.foursquareType = &foursquareType
	return t
}

func (t *SendVenueService) GooglePlaceID(googlePlaceID string) *SendVenueService {
	t.googlePlaceID = &googlePlaceID
	return t
}

func (t *SendVenueService) GooglePlaceType(googlePlaceType string) *SendVenueService {
	t.googlePlaceType = &googlePlaceType
	return t
}

func (t *SendVenueService) DisableNotification(disableNotification bool) *SendVenueService {
	t.disableNotification = &disableNotification
	return t
}

func (t *SendVenueService) ProtectContent(protectContent bool) *SendVenueService {
	t.protectContent = &protectContent
	return t
}

func (t *SendVenueService) ReplyToMessageID(replyToMessageID int64) *SendVenueService {
	t.replyToMessageID = &replyToMessageID
	return t
}

func (t *SendVenueService) AllowSendingWithoutReply(allowSendingWithoutReply bool) *SendVenueService {
	t.allowSendingWithoutReply = &allowSendingWithoutReply
	return t
}

func (t *SendVenueService) InlineKeyboardMarkup(inlineKeyboardMarkup InlineKeyboardMarkup) *SendVenueService {
	json, err := jsoniter.Marshal(&inlineKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.inlineKeyboardMarkup = &jsonString
	return t
}

func (t *SendVenueService) ReplyKeyboardMarkup(replyKeyboardMarkup ReplyKeyboardMarkup) *SendVenueService {
	json, err := jsoniter.Marshal(&replyKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardMarkup = &jsonString
	return t
}

func (t *SendVenueService) ReplyKeyboardRemove(replyKeyboardRemove ReplyKeyboardRemove) *SendVenueService {
	json, err := jsoniter.Marshal(&replyKeyboardRemove)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardRemove = &jsonString
	return t
}

func (t *SendVenueService) ForceReply(forceReply ForceReply) *SendVenueService {
	json, err := jsoniter.Marshal(&forceReply)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.forceReply = &jsonString
	return t
}

func (t *SendVenueService) Do(ctx context.Context, opts ...RequestOption) (res []*SendVenue, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sendVenue",
	}

	r.setParam("chat_id", *t.chatID)
	r.setParam("latitude", *t.latitude)
	r.setParam("longitude", *t.longitude)
	r.setParam("title", *t.title)
	r.setParam("address", *t.address)
	if t.foursquareID != nil {
		r.setParam("foursquare_id", *t.foursquareID)
	}
	if t.foursquareType != nil {
		r.setParam("foursquare_type", *t.foursquareType)
	}
	if t.googlePlaceID != nil {
		r.setParam("google_place_id", *t.googlePlaceID)
	}
	if t.googlePlaceType != nil {
		r.setParam("google_place_type", *t.googlePlaceType)
	}
	if t.messageThreadID != nil {
		r.setParam("message_thread_id", *t.messageThreadID)
	}
	if t.disableNotification != nil {
		r.setParam("disable_notification", *t.disableNotification)
	}
	if t.protectContent != nil {
		r.setParam("protect_content", *t.protectContent)
	}
	if t.replyToMessageID != nil {
		r.setParam("reply_to_message_id", *t.replyToMessageID)
	}
	if t.allowSendingWithoutReply != nil {
		r.setParam("allow_sending_without_reply", *t.allowSendingWithoutReply)
	}
	if t.inlineKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.inlineKeyboardMarkup)
	}
	if t.replyKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.replyKeyboardMarkup)
	}
	if t.replyKeyboardRemove != nil {
		r.setParam("reply_markup", *t.replyKeyboardRemove)
	}
	if t.forceReply != nil {
		r.setParam("reply_markup", *t.forceReply)
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*SendVenue, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SendVenue struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}

type SendContactService struct {
	c                        *Client
	chatID                   *int64
	messageThreadID          *int64
	phoneNumber              *string
	firstName                *string
	lastName                 *string
	vcard                    *string
	disableNotification      *bool
	protectContent           *bool
	replyToMessageID         *int64
	allowSendingWithoutReply *bool
	inlineKeyboardMarkup     *string
	replyKeyboardMarkup      *string
	replyKeyboardRemove      *string
	forceReply               *string
}

func (t *SendContactService) ChatID(chatID int64) *SendContactService {
	t.chatID = &chatID
	return t
}

func (t *SendContactService) MessageThreadID(messageThreadID int64) *SendContactService {
	t.messageThreadID = &messageThreadID
	return t
}

func (t *SendContactService) PhoneNumber(phoneNumber string) *SendContactService {
	t.phoneNumber = &phoneNumber
	return t
}

func (t *SendContactService) FirstName(firstName string) *SendContactService {
	t.firstName = &firstName
	return t
}

func (t *SendContactService) LastName(lastName string) *SendContactService {
	t.lastName = &lastName
	return t
}

func (t *SendContactService) Vcard(vcard string) *SendContactService {
	t.vcard = &vcard
	return t
}

func (t *SendContactService) DisableNotification(disableNotification bool) *SendContactService {
	t.disableNotification = &disableNotification
	return t
}

func (t *SendContactService) ProtectContent(protectContent bool) *SendContactService {
	t.protectContent = &protectContent
	return t
}

func (t *SendContactService) ReplyToMessageID(replyToMessageID int64) *SendContactService {
	t.replyToMessageID = &replyToMessageID
	return t
}

func (t *SendContactService) AllowSendingWithoutReply(allowSendingWithoutReply bool) *SendContactService {
	t.allowSendingWithoutReply = &allowSendingWithoutReply
	return t
}

func (t *SendContactService) InlineKeyboardMarkup(inlineKeyboardMarkup InlineKeyboardMarkup) *SendContactService {
	json, err := jsoniter.Marshal(&inlineKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.inlineKeyboardMarkup = &jsonString
	return t
}

func (t *SendContactService) ReplyKeyboardMarkup(replyKeyboardMarkup ReplyKeyboardMarkup) *SendContactService {
	json, err := jsoniter.Marshal(&replyKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardMarkup = &jsonString
	return t
}

func (t *SendContactService) ReplyKeyboardRemove(replyKeyboardRemove ReplyKeyboardRemove) *SendContactService {
	json, err := jsoniter.Marshal(&replyKeyboardRemove)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardRemove = &jsonString
	return t
}

func (t *SendContactService) ForceReply(forceReply ForceReply) *SendContactService {
	json, err := jsoniter.Marshal(&forceReply)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.forceReply = &jsonString
	return t
}

func (t *SendContactService) Do(ctx context.Context, opts ...RequestOption) (res []*SendContact, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sendContact",
	}

	r.setParam("chat_id", *t.chatID)
	r.setParam("phone_number", *t.phoneNumber)
	r.setParam("first_name", *t.firstName)
	if t.lastName != nil {
		r.setParam("last_name", *t.lastName)
	}
	if t.vcard != nil {
		r.setParam("vcard", *t.vcard)
	}
	if t.messageThreadID != nil {
		r.setParam("message_thread_id", *t.messageThreadID)
	}
	if t.disableNotification != nil {
		r.setParam("disable_notification", *t.disableNotification)
	}
	if t.protectContent != nil {
		r.setParam("protect_content", *t.protectContent)
	}
	if t.replyToMessageID != nil {
		r.setParam("reply_to_message_id", *t.replyToMessageID)
	}
	if t.allowSendingWithoutReply != nil {
		r.setParam("allow_sending_without_reply", *t.allowSendingWithoutReply)
	}
	if t.inlineKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.inlineKeyboardMarkup)
	}
	if t.replyKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.replyKeyboardMarkup)
	}
	if t.replyKeyboardRemove != nil {
		r.setParam("reply_markup", *t.replyKeyboardRemove)
	}
	if t.forceReply != nil {
		r.setParam("reply_markup", *t.forceReply)
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*SendContact, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SendContact struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}

type SendPollService struct {
	c                        *Client
	chatID                   *int64
	messageThreadID          *int64
	question                 *string
	options                  *string
	isAnonymous              *bool
	pollType                 *string
	allowsMultipleAnswers    *bool
	correctOptionID          *int64
	explanation              *string
	explanationParseMode     *string
	explanationEntities      *string
	openPeriod               *int64
	closeDate                *int64
	isClosed                 *bool
	disableNotification      *bool
	protectContent           *bool
	replyToMessageID         *int64
	allowSendingWithoutReply *bool
	inlineKeyboardMarkup     *string
	replyKeyboardMarkup      *string
	replyKeyboardRemove      *string
	forceReply               *string
}

func (t *SendPollService) ChatID(chatID int64) *SendPollService {
	t.chatID = &chatID
	return t
}

func (t *SendPollService) MessageThreadID(messageThreadID int64) *SendPollService {
	t.messageThreadID = &messageThreadID
	return t
}

func (t *SendPollService) Question(question string) *SendPollService {
	t.question = &question
	return t
}

func (t *SendPollService) Options(options []string) *SendPollService {
	json, err := jsoniter.Marshal(&options)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.options = &jsonString
	return t
}

func (t *SendPollService) IsAnonymous(isAnonymous bool) *SendPollService {
	t.isAnonymous = &isAnonymous
	return t
}

func (t *SendPollService) PollType(pollType PollType) *SendPollService {
	json, err := jsoniter.Marshal(&pollType)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.pollType = &jsonString
	return t
}

func (t *SendPollService) AllowsMultipleAnswers(allowsMultipleAnswers bool) *SendPollService {
	t.allowsMultipleAnswers = &allowsMultipleAnswers
	return t
}

func (t *SendPollService) CorrectOptionID(correctOptionID int64) *SendPollService {
	t.correctOptionID = &correctOptionID
	return t
}

func (t *SendPollService) Explanation(explanation string) *SendPollService {
	t.explanation = &explanation
	return t
}

func (t *SendPollService) ExplanationParseMode(explanationParseMode string) *SendPollService {
	t.explanationParseMode = &explanationParseMode
	return t
}

func (t *SendPollService) ExplanationEntities(explanationEntities []MessageEntity) *SendPollService {
	json, err := jsoniter.Marshal(&explanationEntities)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.explanationEntities = &jsonString
	return t
}

func (t *SendPollService) OpenPeriod(openPeriod int64) *SendPollService {
	t.openPeriod = &openPeriod
	return t
}

func (t *SendPollService) CloseDate(closeDate int64) *SendPollService {
	t.closeDate = &closeDate
	return t
}

func (t *SendPollService) IsClosed(isClosed bool) *SendPollService {
	t.isClosed = &isClosed
	return t
}

func (t *SendPollService) DisableNotification(disableNotification bool) *SendPollService {
	t.disableNotification = &disableNotification
	return t
}

func (t *SendPollService) ProtectContent(protectContent bool) *SendPollService {
	t.protectContent = &protectContent
	return t
}

func (t *SendPollService) ReplyToMessageID(replyToMessageID int64) *SendPollService {
	t.replyToMessageID = &replyToMessageID
	return t
}

func (t *SendPollService) AllowSendingWithoutReply(allowSendingWithoutReply bool) *SendPollService {
	t.allowSendingWithoutReply = &allowSendingWithoutReply
	return t
}

func (t *SendPollService) InlineKeyboardMarkup(inlineKeyboardMarkup InlineKeyboardMarkup) *SendPollService {
	json, err := jsoniter.Marshal(&inlineKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.inlineKeyboardMarkup = &jsonString
	return t
}

func (t *SendPollService) ReplyKeyboardMarkup(replyKeyboardMarkup ReplyKeyboardMarkup) *SendPollService {
	json, err := jsoniter.Marshal(&replyKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardMarkup = &jsonString
	return t
}

func (t *SendPollService) ReplyKeyboardRemove(replyKeyboardRemove ReplyKeyboardRemove) *SendPollService {
	json, err := jsoniter.Marshal(&replyKeyboardRemove)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardRemove = &jsonString
	return t
}

func (t *SendPollService) ForceReply(forceReply ForceReply) *SendPollService {
	json, err := jsoniter.Marshal(&forceReply)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.forceReply = &jsonString
	return t
}

func (t *SendPollService) Do(ctx context.Context, opts ...RequestOption) (res []*SendPoll, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sendPoll",
	}

	r.setParam("chat_id", *t.chatID)
	r.setParam("question", *t.question)
	r.setParam("options", *t.options)
	if t.isAnonymous != nil {
		r.setParam("isAnonymous", *t.isAnonymous)
	}
	if t.pollType != nil {
		r.setParam("pollType", *t.pollType)
	}
	if t.allowsMultipleAnswers != nil {
		r.setParam("allowsMultipleAnswers", *t.allowsMultipleAnswers)
	}
	if t.correctOptionID != nil {
		r.setParam("correctOptionID", *t.correctOptionID)
	}
	if t.explanation != nil {
		r.setParam("explanation", *t.explanation)
	}
	if t.explanationParseMode != nil {
		r.setParam("explanationParseMode", *t.explanationParseMode)
	}
	if t.explanationEntities != nil {
		r.setParam("explanationEntities", *t.explanationEntities)
	}
	if t.openPeriod != nil {
		r.setParam("openPeriod", *t.openPeriod)
	}
	if t.closeDate != nil {
		r.setParam("closeDate", *t.closeDate)
	}
	if t.isClosed != nil {
		r.setParam("isClosed", *t.isClosed)
	}
	if t.messageThreadID != nil {
		r.setParam("message_thread_id", *t.messageThreadID)
	}
	if t.disableNotification != nil {
		r.setParam("disable_notification", *t.disableNotification)
	}
	if t.protectContent != nil {
		r.setParam("protect_content", *t.protectContent)
	}
	if t.replyToMessageID != nil {
		r.setParam("reply_to_message_id", *t.replyToMessageID)
	}
	if t.allowSendingWithoutReply != nil {
		r.setParam("allow_sending_without_reply", *t.allowSendingWithoutReply)
	}
	if t.inlineKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.inlineKeyboardMarkup)
	}
	if t.replyKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.replyKeyboardMarkup)
	}
	if t.replyKeyboardRemove != nil {
		r.setParam("reply_markup", *t.replyKeyboardRemove)
	}
	if t.forceReply != nil {
		r.setParam("reply_markup", *t.forceReply)
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*SendPoll, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SendPoll struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}

type SendDiceService struct {
	c                        *Client
	chatID                   *int64
	messageThreadID          *int64
	emoji                    *string
	disableNotification      *bool
	protectContent           *bool
	replyToMessageID         *int64
	allowSendingWithoutReply *bool
	inlineKeyboardMarkup     *string
	replyKeyboardMarkup      *string
	replyKeyboardRemove      *string
	forceReply               *string
}

func (t *SendDiceService) ChatID(chatID int64) *SendDiceService {
	t.chatID = &chatID
	return t
}

func (t *SendDiceService) MessageThreadID(messageThreadID int64) *SendDiceService {
	t.messageThreadID = &messageThreadID
	return t
}

func (t *SendDiceService) Emoji(emoji string) *SendDiceService {
	t.emoji = &emoji
	return t
}

func (t *SendDiceService) DisableNotification(disableNotification bool) *SendDiceService {
	t.disableNotification = &disableNotification
	return t
}

func (t *SendDiceService) ProtectContent(protectContent bool) *SendDiceService {
	t.protectContent = &protectContent
	return t
}

func (t *SendDiceService) ReplyToMessageID(replyToMessageID int64) *SendDiceService {
	t.replyToMessageID = &replyToMessageID
	return t
}

func (t *SendDiceService) AllowSendingWithoutReply(allowSendingWithoutReply bool) *SendDiceService {
	t.allowSendingWithoutReply = &allowSendingWithoutReply
	return t
}

func (t *SendDiceService) InlineKeyboardMarkup(inlineKeyboardMarkup InlineKeyboardMarkup) *SendDiceService {
	json, err := jsoniter.Marshal(&inlineKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.inlineKeyboardMarkup = &jsonString
	return t
}

func (t *SendDiceService) ReplyKeyboardMarkup(replyKeyboardMarkup ReplyKeyboardMarkup) *SendDiceService {
	json, err := jsoniter.Marshal(&replyKeyboardMarkup)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardMarkup = &jsonString
	return t
}

func (t *SendDiceService) ReplyKeyboardRemove(replyKeyboardRemove ReplyKeyboardRemove) *SendDiceService {
	json, err := jsoniter.Marshal(&replyKeyboardRemove)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.replyKeyboardRemove = &jsonString
	return t
}

func (t *SendDiceService) ForceReply(forceReply ForceReply) *SendDiceService {
	json, err := jsoniter.Marshal(&forceReply)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.forceReply = &jsonString
	return t
}

func (t *SendDiceService) Do(ctx context.Context, opts ...RequestOption) (res []*SendDice, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sendDice",
	}

	r.setParam("chat_id", *t.chatID)
	if t.messageThreadID != nil {
		r.setParam("message_thread_id", *t.messageThreadID)
	}
	if t.emoji != nil {
		r.setParam("emoji", *t.emoji)
	}
	if t.disableNotification != nil {
		r.setParam("disable_notification", *t.disableNotification)
	}
	if t.protectContent != nil {
		r.setParam("protect_content", *t.protectContent)
	}
	if t.replyToMessageID != nil {
		r.setParam("reply_to_message_id", *t.replyToMessageID)
	}
	if t.allowSendingWithoutReply != nil {
		r.setParam("allow_sending_without_reply", *t.allowSendingWithoutReply)
	}
	if t.inlineKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.inlineKeyboardMarkup)
	}
	if t.replyKeyboardMarkup != nil {
		r.setParam("reply_markup", *t.replyKeyboardMarkup)
	}
	if t.replyKeyboardRemove != nil {
		r.setParam("reply_markup", *t.replyKeyboardRemove)
	}
	if t.forceReply != nil {
		r.setParam("reply_markup", *t.forceReply)
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*SendDice, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SendDice struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}

type SendChatActionService struct {
	c               *Client
	chatID          *int64
	messageThreadID *int64
	action          *string
}

func (t *SendChatActionService) ChatID(chatID int64) *SendChatActionService {
	t.chatID = &chatID
	return t
}

func (t *SendChatActionService) MessageThreadID(messageThreadID int64) *SendChatActionService {
	t.messageThreadID = &messageThreadID
	return t
}

func (t *SendChatActionService) Action(action string) *SendChatActionService {
	t.action = &action
	return t
}

func (t *SendChatActionService) Do(ctx context.Context, opts ...RequestOption) (res []*SendChatAction, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sendChatAction",
	}

	r.setParam("chat_id", *t.chatID)
	r.setParam("action", *t.action)
	if t.messageThreadID != nil {
		r.setParam("message_thread_id", *t.messageThreadID)
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*SendChatAction, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SendChatAction struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}
