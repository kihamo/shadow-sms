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
	"sort"
	"strings"
)

const (
	BalanceMethod = "outbox/balance/json"
	SendSmsMethod = "outbox/send/json"
)

type client struct {
	apiUrl   *url.URL
	login    string
	password string
	token    string
	sender   string
}

func NewClient(apiUrl string, login, password, token, sender string) (*client, error) {
	u, err := url.Parse(strings.TrimRight(apiUrl, "/") + "/")
	if err != nil {
		return nil, err
	}

	return &client{
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

	body, err := c.prepareBody(map[string]interface{}{
		"login": c.login,
		"target": phone,
		"message": message,
		"sender": c.sender,
	})
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

	sendResponse := &sendResponse{}

	if err := json.Unmarshal(responseBody, sendResponse); err != nil {
		return -1, err
	}

	if len(sendResponse.MessageInfos) != 1 {
		return -1, fmt.Errorf("No messages found in response")
	}

	return sendResponse.MessageInfos[0].Price, nil
}

func (c *client) Balance(ctx context.Context) (float64, error) {
	u := *(c.apiUrl)
	u.Path += BalanceMethod

	body, err := c.prepareBody(map[string]interface{}{
		"login": c.login,
	})
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

	if balanceResponse.Status < 0 {
		return -1, fmt.Errorf("Provider returns error with code #%d", balanceResponse.Status)
	}

	return balanceResponse.Balance, nil
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

func (c *client) prepareBody(body map[string]interface{}) ([]byte, error){

	body["sign"] = c.createSign(body)
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return bodyJson, nil

}

func (c *client) createSign(body map[string]interface{}) string {

	params := make([]string, 0, len(body))
	for k, v := range body {
		params = append(params, k+"="+v.(string))
	}

	sort.Strings(params)

	sig := md5.New()
	io.WriteString(sig, strings.Join(params, ""))
	io.WriteString(sig, c.token)

	return hex.EncodeToString(sig.Sum(nil))

}