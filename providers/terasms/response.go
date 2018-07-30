package terasms

type messageInfo struct {
	id     string  `json:"id"`
	msisdn string  `json:"msisdn"`
	status int     `json:"status"`
	price  float64 `json:"price"`
}

type sendResponse struct {
	status            int32          `json:"status"`
	statusDescription string         `json:"status_description"`
	messageInfos      []*messageInfo `json:"message_infos"`
}

type balanceResponse struct {
	status            int32   `json:"status"`
	statusDescription string  `json:"status_description"`
	balance           float64 `json:"balance"`
}
