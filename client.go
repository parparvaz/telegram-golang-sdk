package telegram

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"telegram/common"

	"github.com/bitly/go-simplejson"
	jsoniter "github.com/json-iterator/go"
)

type Client struct {
	Token      string
	SecretKey  string
	BaseURL    string
	UserAgent  string
	HTTPClient *http.Client
	Debug      bool
	Logger     *log.Logger
	TimeOffset int64
	do         doFunc
}

func NewClient(token, secretKey string) *Client {
	return &Client{
		Token:      token,
		SecretKey:  secretKey,
		UserAgent:  "Telegram/Golang",
		BaseURL:    baseAPIMainURL + token,
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, "Telegram-golang ", log.LstdFlags),
		Debug:      true,
	}
}

const (
	baseAPIMainURL = "https://api.telegram.org/bot"
)

type doFunc func(req *http.Request) (*http.Response, error)

func (c *Client) GetMe() *GetMeService {
	return &GetMeService{c: c}
}

func (c *Client) callAPI(ctx context.Context, r *request, opts ...RequestOption) (data []byte, err error) {
	err = c.parseRequest(r, opts...)
	if err != nil {
		return []byte{}, err
	}
	req, err := http.NewRequest(r.method, r.fullURL, r.body)

	if err != nil {
		return []byte{}, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header
	c.debug("request: %#v", req)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		cerr := res.Body.Close()

		if err == nil && cerr != nil {
			err = cerr
		}
	}()

	c.debug("response: %#v", res)
	c.debug("response body: %s", string(data))
	c.debug("response status code: %d", res.StatusCode)

	if res.StatusCode >= http.StatusBadRequest {
		apiErr := new(common.APIError)
		e := json.Unmarshal(data, apiErr)
		if e != nil {
			c.debug("failed to unmarshal json: %s", e)
		}
		return nil, apiErr
	}
	return data, nil
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
	}
}

func (r *request) validate() (err error) {
	if r.query == nil {
		r.query = url.Values{}
	}
	if r.form == nil {
		r.form = url.Values{}
	}
	return nil
}

func (r *request) setParam(key string, value interface{}) *request {
	if r.query == nil {
		r.query = url.Values{}
	}

	if reflect.TypeOf(value).Kind() == reflect.Slice {
		v, err := json.Marshal(value)
		if err == nil {
			value = string(v)
		}
	}

	r.query.Set(key, fmt.Sprintf("%v", value))
	return r
}

func (c *Client) parseRequest(r *request, opts ...RequestOption) (err error) {
	for _, opt := range opts {
		opt(r)
	}
	err = r.validate()
	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, r.endpoint)
	queryString := r.query.Encode()
	body := &bytes.Buffer{}
	bodyString := r.form.Encode()
	header := http.Header{}
	if r.header != nil {
		header = r.header.Clone()
	}

	if bodyString != "" {
		header.Set("Content-Type", "application/json")
		body = bytes.NewBufferString(string(r.json))
	}

	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}

	c.debug("full url: %s, body: %s", fullURL, bodyString)

	r.fullURL = fullURL
	r.header = header
	r.body = body
	return nil
}

func newJSON(data []byte) (j *simplejson.Json, err error) {
	j, err = simplejson.NewJson(data)
	if err != nil {
		return nil, err
	}
	return j, nil
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func (c *Client) NewSendMessageService() *SendMessageService {
	return &SendMessageService{c: c}
}

func (c *Client) NewGetMeService() *GetMeService {
	return &GetMeService{c: c}
}

func (c *Client) NewLogOutService() *LogOutService {
	return &LogOutService{c: c}
}

func (c *Client) NewCloseService() *CloseService {
	return &CloseService{c: c}
}

func (c *Client) NewAnswerCallbackQueryService() *AnswerCallbackQueryService {
	return &AnswerCallbackQueryService{c: c}
}

func (c *Client) NewEditMessageReplyMarkupService() *EditMessageReplyMarkupService {
	return &EditMessageReplyMarkupService{c: c}
}

func (c *Client) NewEditMessageTextService() *EditMessageTextService {
	return &EditMessageTextService{c: c}
}

func (c *Client) NewDeleteMessageService() *DeleteMessageService {
	return &DeleteMessageService{c: c}
}

func (c *Client) NewSetWebhookService() *SetWebhookService {
	return &SetWebhookService{c: c}
}

func (c *Client) NewDeleteWebhookService() *DeleteWebhookService {
	return &DeleteWebhookService{c: c}
}

func (c *Client) NewGetWebhookInfoService() *GetWebhookInfoService {
	return &GetWebhookInfoService{c: c}
}

const (
	PollTypeRegular PollType = "regular"
	PollTypeQuiz    PollType = "quiz"
)

type PollType string
