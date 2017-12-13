package weibo

import (
	"encoding/json"
	"fmt"
	"github.com/beewit/beekit/utils/uhttp"
)

type (
	// Weibo weibo
	Weibo struct {
		ID               int64       `json:"id,omitempty"`                 // 用户UID
		IDStr            string      `json:"idstr,omitempty"`              // 字符串型的用户UID
		ScreenName       string      `json:"screen_name,omitempty"`        //用户昵称
		Name             string      `json:"name,omitempty"`               // 友好显示名称
		Province         string         `json:"province,omitempty"`           // 用户所在省级ID
		City             string         `json:"city,omitempty"`               // 用户所在城市ID
		Location         string      `json:"location,omitempty"`           // 用户所在地
		Description      string      `json:"description,omitempty"`        // 用户个人描述
		URL              string      `json:"url,omitempty"`                // 用户博客地址
		ProfileImageURL  string      `json:"profile_image_url,omitempty"`  // 用户头像地址（中图），50×50像素
		ProfileURL       string      `json:"profile_url,omitempty"`        // 用户的微博统一URL地址
		Domain           string      `json:"domain,omitempty"`             // 用户的个性化域名
		Weihao           string      `json:"weihao,omitempty"`             // 用户的微号
		Gender           string      `json:"gender,omitempty"`             // 性别，m：男、f：女、n：未知
		FollowersCount   int         `json:"followers_count,omitempty"`    // 粉丝数
		FriendsCount     int         `json:"friends_count,omitempty"`      // 关注数
		StatusesCount    int         `json:"statuses_count,omitempty"`     // 微博数
		FavouritesCount  int         `json:"favourites_count,omitempty"`   // 收藏数
		CreatedAt        string      `json:"created_at,omitempty"`         // 用户创建（注册）时间
		Following        bool        `json:"following,omitempty"`          // 暂未支持
		AllowAllActMsg   bool        `json:"allow_all_act_msg,omitempty"`  // 是否允许所有人给我发私信，true：是，false：否
		GeoEnabled       bool        `json:"geo_enabled,omitempty"`        // 是否允许标识用户的地理位置，true：是，false：否
		Verified         bool        `json:"verified,omitempty"`           // 是否是微博认证用户，即加V用户，true：是，false：否
		Remark           string      `json:"remark,omitempty"`             // 用户备注信息，只有在查询用户关系时才返回此字段
		Status           interface{} `json:"status,omitempty"`             // 用户的最近一条微博信息字段 详细
		AllowAllComment  bool        `json:"allow_all_comment,omitempty"`  // 是否允许所有人对我的微博进行评论，true：是，false：否
		AvatarLarge      string      `json:"avatar_large,omitempty"`       // 用户头像地址（大图），180×180像素
		AvatarHD         string      `json:"avatar_hd,omitempty"`          // 用户头像地址（高清），高清头像原图
		VerifiedReason   string      `json:"verified_reason,omitempty"`    // 认证原因
		FollowMe         bool        `json:"follow_me,omitempty"`          // 该用户是否关注当前登录用户，true：是，false：否
		OnlineStatus     int         `json:"online_status,omitempty"`      // 用户的在线状态，0：不在线、1：在线
		BiFollowersCount int         `json:"bi_followers_count,omitempty"` // 用户的互粉数
		Lang             string      `json:"lang,omitempty"`               // 用户当前的语言版本，zh-cn：简体中文，zh-tw：繁体中文，en：英语
	}

	AccessToken struct {
		AccessToken string      `json:"access_token,omitempty"`
		RemindIn    string      `json:"remind_in,omitempty"`
		ExpiresIn   int         `json:"expires_in,omitempty"`
		Uid         string      `json:"uid,omitempty"`
	}
)

// NewWeibo new weibo
func NewWeibo() *Weibo {
	return &Weibo{}
}

// User user
func (w Weibo) User(accessToken, uID string) (Weibo, error) {
	result := Weibo{}
	url := fmt.Sprintf("https://api.weibo.com/2/users/show.json?access_token=%v&uid=%v",
		accessToken,
		uID,
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

func (w Weibo) GetAccessToken(appKey, appSecret, redirectUri, code string) (AccessToken, error) {
	result := AccessToken{}
	url := fmt.Sprintf("https://api.weibo.com/oauth2/access_token?client_id=%v&client_secret=%v&grant_type=authorization_code" +
		"&redirect_uri=%v&code=%v",
		appKey,
		appSecret,
		redirectUri,
		code,
	)
	body, err := uhttp.Cmd(uhttp.Request{
		Method: "POST",
		URL:    url,
	})
	if err != nil {
		return result, err
	}
	//失败 {"error":"invalid_grant","error_code":21325,"request":"/oauth2/access_token","error_uri":"/oauth2/access_token","error_description":"invalid authorization code:85677e9ece16654ca1432e7269a3b86f"}
	println(string(body[:]))
	err = json.Unmarshal(body, &result)
	return result, err
}
