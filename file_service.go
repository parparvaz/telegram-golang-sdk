package telegram

import (
	"context"
	"net/http"
	"telegram/common"
)

type GetFileService struct {
	c      *Client
	fileID *string
}

func (t *GetFileService) FleID(fileID string) *GetFileService {
	t.fileID = &fileID
	return t
}

func (t *GetFileService) Do(ctx context.Context, opts ...RequestOption) (res []*GetFile, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/getFile",
	}

	r.setParam("file_id", *t.fileID)

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*GetFile, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetFile struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}
