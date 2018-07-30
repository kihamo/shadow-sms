package terasms

type sendRequest struct {
	login   string `json:"login"`
	sign    string `json:"sign"`
	target  string `json:"target"`
	message string `json:"message"`
	sender  string `json:"sender"`
}

type balanceRequest struct {
	login string `json:"login"`
	sign  string `json:"sign"`
}
