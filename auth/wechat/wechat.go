package wechat

import "github.com/tiantour/fetch"

type (
	// Token token
	Token struct {
		AccessToken  string `json:"access_token"`  // 网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
		ExpiresIn    int    `json:"expires_in"`    // access_token接口调用凭证超时时间，单位（秒）
		RefreshToken string `json:"refresh_token"` // 用户刷新access_token
		OpenID       string `json:"openid"`        // 用户唯一标识
		Scope        string `json:"scope"`         // 用户授权的作用域，使用逗号（,）分隔
		ErrCode      int    `json:"errcode"`       // 错误代码
		ErrMsg       string `json:"errmsg"`        // 错误消息
	}
	// User user
	User struct {
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
	// Prompt prompt
	Prompt struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}
	// Message message
	Message struct{}
)

// request
func request(requestURL string) ([]byte, error) {
	return fetch.Cmd("get", requestURL)
}
