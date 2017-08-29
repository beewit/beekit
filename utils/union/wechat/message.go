package wechat

import (
	"fmt"

	"github.com/beewit/beekit/utils/uhttp"
)

// Message message
type Message struct{}

// NewMessage new message
func NewMessage() *Message {
	return &Message{}
}

// Do do
func (m Message) Do(body []byte) ([]byte, error) {
	token, err := NewToken().Cache()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%v",
		token,
	)
	return uhttp.Cmd(uhttp.Request{
		Method: "POST",
		URL:    url,
		Body:   body,
	})
}
