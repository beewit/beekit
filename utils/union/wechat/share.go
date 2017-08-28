package wechat

import (
	"encoding/hex"
	"net/url"
	"strconv"
	"time"

	"github.com/beewit/beekit/utils/query"
	"github.com/beewit/beekit/utils"
	"github.com/beewit/beekit/utils/encrypt"
)

// Share share
type Share struct {
	AppID       string `json:"appid" url:"-"`
	JSapiTicket string `json:"jsapi_ticket" url:"jsapi_ticket"`
	Noncestr    string `json:"noncestr" url:"noncestr"`
	Timestamp   string `json:"timestamp" url:"timestamp"`
	URL         string `json:"url" url:"url"`
	Signature   string `json:"signature" url:"-"`
}

// NewShare new share
func NewShare() *Share {
	return &Share{}
}

// Do do
func (s Share) Do(page string) (Share, error) {
	result := Share{
		Noncestr:  utils.NewRandom().String(16),
		Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
		URL:       page,
	}
	ticket, err := NewTicket().Do()
	if err != nil {
		return result, err
	}
	result.JSapiTicket = ticket.Ticket
	params, err := query.Values(result)
	if err != nil {
		return result, err
	}
	query, err := url.QueryUnescape(params.Encode())
	if err != nil {
		return result, err
	}
	sign := encrypt.NewRsae().SHA1(query)
	result.Signature = hex.EncodeToString(sign)
	result.AppID = AppID
	return result, nil
}
