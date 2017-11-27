package terasms

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	AuthByToken = iota + 1
	AuthByLoginAndPassword

	ParamSignature = "sig"
	ParamLogin     = "login"
	ParamPassword  = "password"
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

func (c *client) Send(ctx context.Context, phone, message string) error {
	u := *(c.apiUrl)
	u.Path += "outbox/send"

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)

	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Set(ParamLogin, c.login)
	q.Set("target", phone)
	q.Set("sender", c.sender)
	q.Set("message", message)

	req.URL.RawQuery = q.Encode()
	req.WithContext(ctx)

	resp, err := c.do(req)
	if err != nil {
		return err
	}

	fmt.Println(resp, err)

	return nil
}

func (c *client) Balance(ctx context.Context) (float64, error) {
	u := *(c.apiUrl)
	u.Path += "outbox/balance"

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)

	if err != nil {
		return -1, err
	}

	req.WithContext(ctx)

	resp, err := c.do(req)
	if err != nil {
		return -1, err
	}

	return strconv.ParseFloat(resp, 64)
}

func (c *client) do(r *http.Request) (string, error) {
	c.sign(r)

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body[:]), nil
}

func (c *client) sign(req *http.Request) {
	q := req.URL.Query()

	if c.auth == AuthByLoginAndPassword {
		q.Set(ParamLogin, c.login)
		q.Set(ParamPassword, c.password)
	} else {
		sig := ""

		q.Set(ParamSignature, sig)
	}

	req.URL.RawQuery = q.Encode()
}
