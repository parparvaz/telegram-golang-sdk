package telegram

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type secType int

type params map[string]any

type request struct {
	method     string
	endpoint   string
	query      url.Values
	form       url.Values
	recvWindow int64
	secType    secType
	header     http.Header
	body       io.Reader
	fullURL    string
	json       []byte
}

type RequestOption func(*request)

func (r *request) setFormParams(m params) *request {
	json, err := json.MarshalIndent(m, "", "   ")
	if err != nil {
		return nil
	}
	r.json = json
	for k, v := range m {
		r.setFormParam(k, v)
	}
	return r
}

func (r *request) setFormParam(key string, value interface{}) *request {
	if r.form == nil {
		r.form = url.Values{}
	}

	r.form.Set(key, fmt.Sprintf("%v", value))

	return r
}
