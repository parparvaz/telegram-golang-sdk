package telegram

import (
	"context"
	"net/http"
	"telegram/common"
)

type GetMeService struct {
	c *Client
}

func (t *GetMeService) Do(ctx context.Context, opts ...RequestOption) (res []*GetMe, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/getMe",
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*GetMe, 0)
	err = json.Unmarshal(data, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

type GetMe struct {
	Ok     bool `json:"ok"`
	Result User `json:"result"`
}

type LogOutService struct {
	c *Client
}

func (t *LogOutService) Do(ctx context.Context, opts ...RequestOption) (res []*LogOut, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/logOut",
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*LogOut, 0)
	err = json.Unmarshal(data, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

type LogOut struct {
}

type CloseService struct {
	c *Client
}

func (t *CloseService) Do(ctx context.Context, opts ...RequestOption) (res []*Close, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/close",
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*Close, 0)
	err = json.Unmarshal(data, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

type Close struct {
}

type SetMyNameService struct {
	c            *Client
	name         *string
	languageCode *string
}

func (t *SetMyNameService) Name(name string) *SetMyNameService {
	t.name = &name
	return t
}

func (t *SetMyNameService) LanguageCode(languageCode string) *SetMyNameService {
	t.languageCode = &languageCode
	return t
}

func (t *SetMyNameService) Do(ctx context.Context, opts ...RequestOption) (res []*SetMyName, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/setMyName",
	}

	if t.name != nil {
		r.setParam("name", *t.name)
	}
	if t.languageCode != nil {
		r.setParam("language_code", *t.languageCode)
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*SetMyName, 0)
	err = json.Unmarshal(data, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

type SetMyName struct {
	Ok     bool `json:"ok"`
	Result bool `json:"result"`
}

type GetMyNameService struct {
	c            *Client
	languageCode *string
}

func (t *GetMyNameService) LanguageCode(languageCode string) *GetMyNameService {
	t.languageCode = &languageCode
	return t
}

func (t *GetMyNameService) Do(ctx context.Context, opts ...RequestOption) (res []*GetMyName, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/getMyName",
	}

	if t.languageCode != nil {
		r.setParam("language_code", *t.languageCode)
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*GetMyName, 0)
	err = json.Unmarshal(data, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

type GetMyName struct {
	Ok     bool `json:"ok"`
	Result bool `json:"result"`
}

type SetMyDescriptionService struct {
	c            *Client
	description  *string
	languageCode *string
}

func (t *SetMyDescriptionService) Description(description string) *SetMyDescriptionService {
	t.description = &description
	return t
}

func (t *SetMyDescriptionService) LanguageCode(languageCode string) *SetMyDescriptionService {
	t.languageCode = &languageCode
	return t
}

func (t *SetMyDescriptionService) Do(ctx context.Context, opts ...RequestOption) (res []*SetMyDescription, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/setMyDescription",
	}

	if t.description != nil {
		r.setParam("description", *t.description)
	}
	if t.languageCode != nil {
		r.setParam("language_code", *t.languageCode)
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*SetMyDescription, 0)
	err = json.Unmarshal(data, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

type SetMyDescription struct {
	Ok     bool `json:"ok"`
	Result bool `json:"result"`
}

type GetMyDescriptionService struct {
	c            *Client
	languageCode *string
}

func (t *GetMyDescriptionService) LanguageCode(languageCode string) *GetMyDescriptionService {
	t.languageCode = &languageCode
	return t
}

func (t *GetMyDescriptionService) Do(ctx context.Context, opts ...RequestOption) (res []*GetMyDescription, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/getMyDescription",
	}

	if t.languageCode != nil {
		r.setParam("language_code", *t.languageCode)
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*GetMyDescription, 0)
	err = json.Unmarshal(data, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

type GetMyDescription struct {
	Ok     bool `json:"ok"`
	Result bool `json:"result"`
}

type SetMyShortDescriptionService struct {
	c                *Client
	shortDescription *string
	languageCode     *string
}

func (t *SetMyShortDescriptionService) ShortDescription(shortDescription string) *SetMyShortDescriptionService {
	t.shortDescription = &shortDescription
	return t
}

func (t *SetMyShortDescriptionService) LanguageCode(languageCode string) *SetMyShortDescriptionService {
	t.languageCode = &languageCode
	return t
}

func (t *SetMyShortDescriptionService) Do(ctx context.Context, opts ...RequestOption) (res []*SetMyShortDescription, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/setMyShortDescription",
	}

	if t.shortDescription != nil {
		r.setParam("short_description", *t.shortDescription)
	}
	if t.languageCode != nil {
		r.setParam("language_code", *t.languageCode)
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*SetMyShortDescription, 0)
	err = json.Unmarshal(data, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

type SetMyShortDescription struct {
	Ok     bool `json:"ok"`
	Result bool `json:"result"`
}

type GetMyShortDescriptionService struct {
	c            *Client
	languageCode *string
}

func (t *GetMyShortDescriptionService) LanguageCode(languageCode string) *GetMyShortDescriptionService {
	t.languageCode = &languageCode
	return t
}

func (t *GetMyShortDescriptionService) Do(ctx context.Context, opts ...RequestOption) (res []*GetMyShortDescription, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/getMyShortDescription",
	}

	if t.languageCode != nil {
		r.setParam("language_code", *t.languageCode)
	}

	data, err := t.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	data = common.ToJSONList(data)
	if err != nil {
		return nil, err
	}

	res = make([]*GetMyShortDescription, 0)
	err = json.Unmarshal(data, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

type GetMyShortDescription struct {
	Ok     bool `json:"ok"`
	Result bool `json:"result"`
}
