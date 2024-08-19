package telegram

import (
	"context"
	"net/http"
	"telegram/common"
)

type GetUserProfilePhotosService struct {
	c      *Client
	chatID *int64
	offset *int64
	limit  *int64
}

func (t *GetUserProfilePhotosService) ChatID(chatID int64) *GetUserProfilePhotosService {
	t.chatID = &chatID
	return t
}

func (t *GetUserProfilePhotosService) Offset(offset int64) *GetUserProfilePhotosService {
	t.offset = &offset
	return t
}

func (t *GetUserProfilePhotosService) Limit(limit int64) *GetUserProfilePhotosService {
	t.limit = &limit
	return t
}

func (t *GetUserProfilePhotosService) Do(ctx context.Context, opts ...RequestOption) (res []*GetUserProfilePhotos, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/getUserProfilePhotos",
	}

	r.setParam("chat_id", *t.chatID)
	if t.offset != nil {
		r.setParam("offset", *t.offset)
	}
	if t.limit != nil {
		r.setParam("limit", *t.limit)
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*GetUserProfilePhotos, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetUserProfilePhotos struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}
