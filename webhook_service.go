package telegram

import (
	"context"
	"net/http"
	"telegram/common"

	jsoniter "github.com/json-iterator/go"
)

type SetWebhookService struct {
	c                  *Client
	url                *string
	certificate        *string
	ipAddress          *string
	maxConnections     *int64
	allowedUpdates     *string
	dropPendingUpdates *bool
	secretToken        *string
}

func (t *SetWebhookService) URL(url string) *SetWebhookService {
	t.url = &url
	return t
}

func (t *SetWebhookService) Certificate(certificate InputFile) *SetWebhookService {
	json, err := jsoniter.Marshal(&certificate)
	if err != nil {
		return nil
	}

	jsonString := string(json)

	t.certificate = &jsonString
	return t
}

func (t *SetWebhookService) IpAddress(ipAddress string) *SetWebhookService {
	t.ipAddress = &ipAddress
	return t
}

func (t *SetWebhookService) MaxConnections(maxConnections int64) *SetWebhookService {
	t.maxConnections = &maxConnections
	return t
}

func (t *SetWebhookService) AllowedUpdates(allowedUpdates string) *SetWebhookService {
	t.allowedUpdates = &allowedUpdates
	return t
}

func (t *SetWebhookService) DropPendingUpdates(dropPendingUpdates bool) *SetWebhookService {
	t.dropPendingUpdates = &dropPendingUpdates
	return t
}

func (t *SetWebhookService) SecretToken(secretToken string) *SetWebhookService {
	t.secretToken = &secretToken
	return t
}

func (t *SetWebhookService) Do(ctx context.Context, opts ...RequestOption) (res []*SetWebhook, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/setWebhook",
	}

	r.setParam("url", *t.url)
	if t.certificate != nil {
		r.setParam("certificate", *t.certificate)
	}
	if t.ipAddress != nil {
		r.setParam("ip_address", *t.ipAddress)
	}
	if t.maxConnections != nil {
		r.setParam("max_connections", *t.maxConnections)
	}
	if t.allowedUpdates != nil {
		r.setParam("allowed_updates", *t.allowedUpdates)
	}
	if t.dropPendingUpdates != nil {
		r.setParam("drop_pending_updates", *t.dropPendingUpdates)
	}
	if t.secretToken != nil {
		r.setParam("secret_token", *t.secretToken)
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*SetWebhook, 0)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type SetWebhook struct {
	Ok          bool   `json:"ok"`
	Result      bool   `json:"result"`
	Description string `json:"description"`
}

type DeleteWebhookService struct {
	c                  *Client
	dropPendingUpdates *bool
}

func (t *DeleteWebhookService) DropPendingUpdates(dropPendingUpdates bool) *DeleteWebhookService {
	t.dropPendingUpdates = &dropPendingUpdates
	return t
}

func (t *DeleteWebhookService) Do(ctx context.Context, opts ...RequestOption) (res []*DeleteWebhook, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/deleteWebhook",
	}

	if t.dropPendingUpdates != nil {
		r.setParam("drop_pending_updates", *t.dropPendingUpdates)
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*DeleteWebhook, 0)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type DeleteWebhook struct {
	Ok          bool   `json:"ok"`
	Result      bool   `json:"result"`
	Description string `json:"description"`
}

type GetWebhookInfoService struct {
	c *Client
}

func (t *GetWebhookInfoService) Do(ctx context.Context, opts ...RequestOption) (res []*GetWebhookInfo, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/getWebhookInfo",
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*GetWebhookInfo, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetWebhookInfo struct {
	Ok     bool        `json:"ok"`
	Result WebhookInfo `json:"result"`
}
