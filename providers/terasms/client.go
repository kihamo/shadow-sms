package terasms

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/json"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strings"
)

const (
	ParamSignature = "sign"

	BalanceMethod = "outbox/balance/json"
	SendSmsMethod = "outbox/send/json"
)

type client struct {
	auth     int
	apiUrl   *url.URL
	login    string
	password string
	token    string
	sender   string
}

func NewClient(apiUrl string, auth int, login, password, token, sender string) (*client, error) {
	u, err := url.Parse(strings.TrimRight(apiUrl, "/") + "/")
	if err != nil {
		return nil, err
	}

	return &client{
		auth:     auth,
		apiUrl:   u,
		login:    login,
		password: password,
		token:    token,
		sender:   sender,
	}, nil
}

func (c *client) Send(ctx context.Context, phone, message string) (float64, error) {
	u := *(c.apiUrl)
	u.Path += SendSmsMethod

	body, err := c.prepareBalanceBody()
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(body))
	if err != nil {
		return 0, err
	}

	req.WithContext(ctx)

	responseBody, err := c.do(req)
	if err != nil {
		return 0, err
	}

	sendResponse := &sendResponse{}

	if err := json.Unmarshal(responseBody, sendResponse); err != nil {
		return 0, err
	}

	if len(sendResponse.messageInfos) != 1 {
		return 0, fmt.Errorf("No messages found in response")
	}

	return sendResponse.messageInfos[0].price, nil
}

func (c *client) Balance(ctx context.Context) (float64, error) {
	u := *(c.apiUrl)
	u.Path += BalanceMethod

	body, err := c.prepareBalanceBody()
	if err != nil {
		return -1, err
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(body))
	if err != nil {
		return -1, err
	}

	req.WithContext(ctx)

	responseBody, err := c.do(req)
	if err != nil {
		return -1, err
	}

	balanceResponse := &balanceResponse{}

	if err := json.Unmarshal(responseBody, balanceResponse); err != nil {
		return -1, err
	}

	if balanceResponse.status < 0 {
		return -1, fmt.Errorf("Provider returns error with code #%d", balanceResponse.status)
	}

	return balanceResponse.balance, nil
}

func (c *client) do(r *http.Request) ([]byte, error) {

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *client) createSign(v reflect.Value) string {

	params := make([]string, 0, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		if v.Type().Field(i).Name == ParamSignature {
			continue
		}
		params = append(params, v.Type().Field(i).Name+"="+v.Field(i).String())
	}

	sort.Strings(params)

	sig := md5.New()
	io.WriteString(sig, strings.Join(params, ""))
	io.WriteString(sig, c.token)

	return hex.EncodeToString(sig.Sum(nil))
}

func (c *client) prepareSendBody(phone string, message string) ([]byte, error) {

	body := &sendRequest{
		login: c.login,
		target: phone,
		message: message,
		sender: c.sender,
	}

	body.sign = c.createSign(reflect.ValueOf(body))

	bodyJson, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return bodyJson, nil

}

func (c *client) prepareBalanceBody() ([]byte, error) {

	body := &balanceRequest{
		login: c.login,
	}

	body.sign = c.createSign(reflect.ValueOf(body))

	bodyJson, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return bodyJson, nil

}
