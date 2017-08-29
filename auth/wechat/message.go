package wechat

import (
	"encoding/json"
	"fmt"

	"github.com/beewit/beekit/redis"
	"github.com/beewit/beekit/utils/uhttp"
)

// Token message
func (m *Message) token(appID, appSecret string) (string, error) {
	key := "wechat_access_token"
	token, err := redis.Cache.GetString(key) //  cache.String.GET(key).Str()
	if err != nil {
		result := Token{}
		url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%v&secret=%v",
			appID,
			appSecret,
		)
		body, err := request(url)
		if err == nil && json.Unmarshal(body, &result) == nil {
			_, _ = redis.Cache.SetAndExpire(key, result.AccessToken, 7200)
			return result.AccessToken, nil
		}
		return "", err
	}
	return token, nil

}

// Push message
func (m *Message) Push(appID, appSecret string, data interface{}) ([]byte, error) {
	accessToken, err := m.token(appID, appSecret)
	if err != nil {
		return nil, err
	}
	requestURL := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%v", accessToken)
	requestData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return uhttp.Cmd("post", requestURL, requestData)
}
