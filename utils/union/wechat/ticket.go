package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/beewit/beekit/utils/uhttp"
)

// Ticket ticket
type Ticket struct {
	ErrCode   string `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

// NewTicket new ticket
func NewTicket() *Ticket {
	return &Ticket{}
}

// Do do
func (t Ticket) Do() (Ticket, error) {
	result := Ticket{}
	token, err := NewToken().Cache()
	if err != nil {
		return result, err
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%v&type=jsapi",
		token,
	)
	body, err := uhttp.Cmd(uhttp.Request{
		Method: "GET",
		URL:    url,
	})
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
