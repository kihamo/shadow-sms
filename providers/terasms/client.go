package terasms

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

const (
	AuthByToken = iota + 1
	AuthByLoginAndPassword

	ParamSignature = "sign"
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
	q.Set("target", phone)
	q.Set("sender", c.sender)
	q.Set("message", message)

	req.URL.RawQuery = q.Encode()
	req.WithContext(ctx)

	_, code, err := c.do(req)
	if err != nil {
		return err
	}

	if code < 0 {
		return fmt.Errorf("Provider returns error with code #%d", code)
	}

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

	resp, code, err := c.do(req)
	if err != nil {
		return -1, err
	}

	if code < 0 {
		return -1, fmt.Errorf("Provider returns error with code #%d", code)
	}

	return strconv.ParseFloat(resp, 64)
}

func (c *client) do(r *http.Request) (string, int64, error) {
	c.sign(r)

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}

	response := string(body[:])
	code, _ := strconv.ParseInt(response, 10, 0)

	return response, code, nil
}

func (c *client) sign(req *http.Request) {
	q := req.URL.Query()
	q.Set(ParamLogin, c.login)

	if c.auth == AuthByLoginAndPassword {
		q.Set(ParamPassword, c.password)
	} else {
		params := make([]string, 0, len(q))
		for k, v := range q {
			if k == ParamSignature {
				continue
			}

			params = append(params, k+"="+v[0])
		}

		sort.Strings(params)

		sig := md5.New()
		io.WriteString(sig, strings.Join(params, ""))
		io.WriteString(sig, c.token)

		q.Set(ParamSignature, hex.EncodeToString(sig.Sum(nil)))
	}

	req.URL.RawQuery = q.Encode()
}
