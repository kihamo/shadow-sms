package terasms

type messageInfo struct {
	Id     string  `json:"id"`
	Msisdn string  `json:"msisdn"`
	Status int     `json:"status"`
	Price  float64 `json:"price"`
}

type sendResponse struct {
	Status            int32          `json:"status"`
	StatusDescription string         `json:"status_description"`
	MessageInfos      []*messageInfo `json:"message_infos"`
}

type balanceResponse struct {
	Status            int32   `json:"status"`
	StatusDescription string  `json:"status_description"`
	Balance           float64 `json:"balance"`
}
