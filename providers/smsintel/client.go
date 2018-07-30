package smsintel

import (
	"context"

	"github.com/kihamo/smsintel"
	"github.com/kihamo/smsintel/procedure"
)

type Client struct {
	sdk *smsintel.SmsIntel
}

func NewClient(apiUrl, login, password string) (*Client, error) {
	sdk := smsintel.NewSmsIntel(login, password)
	sdk.SetOptions(map[string]string{"api_url": apiUrl})

	return &Client{
		sdk: sdk,
	}, nil
}

func (c *Client) Send(ctx context.Context, phone string, message string) (float64, error) {
	input := &procedure.SendSmsInput{
		Txt: message,
		To:  &phone,
	}

	_, err := c.sdk.SendSmsWithContext(ctx, input)

	return 0, err
}

func (c *Client) Balance(ctx context.Context) (float64, error) {
	info, err := c.sdk.InfoWithContext(ctx, nil)

	if err != nil {
		return -1, err
	}

	return info.Account, nil
}
