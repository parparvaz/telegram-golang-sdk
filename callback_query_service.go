package telegram

import (
	"context"
	"net/http"
	"telegram/common"
)

type AnswerCallbackQueryService struct {
	c               *Client
	callbackQueryID *string
	text            *string
	showAlert       *bool
	url             *string
	cacheTime       *int64
}

func (t *AnswerCallbackQueryService) CallbackQueryID(callbackQueryID string) *AnswerCallbackQueryService {
	t.callbackQueryID = &callbackQueryID
	return t
}

func (t *AnswerCallbackQueryService) Text(text string) *AnswerCallbackQueryService {
	t.text = &text
	return t
}

func (t *AnswerCallbackQueryService) ShowAlert(showAlert bool) *AnswerCallbackQueryService {
	t.showAlert = &showAlert
	return t
}

func (t *AnswerCallbackQueryService) Url(url string) *AnswerCallbackQueryService {
	t.url = &url
	return t
}

func (t *AnswerCallbackQueryService) CacheTime(cacheTime int64) *AnswerCallbackQueryService {
	t.cacheTime = &cacheTime
	return t
}

func (t *AnswerCallbackQueryService) Do(ctx context.Context, opts ...RequestOption) (res []*AnswerCallbackQuery, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/answerCallbackQuery",
	}

	r.setParam("callback_query_id", *t.callbackQueryID)
	if t.text != nil {
		r.setParam("text", *t.text)
	}
	if t.showAlert != nil {
		r.setParam("show_alert", *t.showAlert)
	}
	if t.url != nil {
		r.setParam("url", *t.url)
	}
	if t.cacheTime != nil {
		r.setParam("cache_time", *t.cacheTime)
	}
	data, err := t.c.callAPI(ctx, r, opts...)

	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*AnswerCallbackQuery, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type AnswerCallbackQuery struct {
	Ok     bool `json:"ok"`
	Result bool `json:"result"`
}
