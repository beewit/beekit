package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/beewit/beekit/utils/uhttp"
	"strings"
	"github.com/beewit/beekit/utils/convert"
	"github.com/pkg/errors"
	"github.com/beewit/beekit/log"
)

var (
	// AppID appid
	AppID string

	// AppSecret app secret
	AppSecret string
)

type (
	// Wechat wechat
	Wechat struct {
		OpenID     string   `json:"openid"`     // 用户的唯一标识
		NickName   string   `json:"nickname"`   // 用户昵称
		Sex        int      `json:"sex"`        // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
		Province   string   `json:"province"`   // 用户个人资料填写的省份
		City       string   `json:"city"`       // 普通用户个人资料填写的城市
		Country    string   `json:"country"`    // 国家，如中国为CN
		HeadImgURL string   `json:"headimgurl"` // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
		Privilege  []string `json:"privilege"`  // 用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）
		UnionID    string   `json:"unionid"`    // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。详见：获取用户个人信息（UnionID机制）
		Language   string   `json:"language"`   // 语言
	}
)

// NewWechat new wechat
func NewWechat() *Wechat {
	return &Wechat{}
}

// User user
func (w Wechat) User(accessToken, openID string) (Wechat, error) {
	result := Wechat{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%v&openid=%v",
		accessToken,
		openID,
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

func (w Wechat) GetAccessToken(appId, secret, code string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		appId,
		secret,
		code,
	)
	body, err := uhttp.Cmd(uhttp.Request{
		Method: "GET",
		URL:    url,
	})
	if err != nil {
		return nil, err
	}
	res := string(body[:])
	log.Logger.Info(res)
	flog := strings.Contains(res, "errcode")
	if flog {
		return nil, errors.New(res)
	}
	m, err2 := convert.Byte2Map(body)
	if err2 != nil {
		return nil, nil
	}
	return m, nil
}
